package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	lastRequestTime time.Time
	requestLock     sync.Mutex
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("VIRUSTOTAL_API_KEY")
	if apiKey == "" {
		log.Fatal("VIRUSTOTAL_API_KEY not set in .env")
	}

	// Initialize Gin router
	r := gin.Default()

	// Serve static files (frontend)
	r.Static("/", "./frontend")

	// File upload endpoint
	r.POST("/upload", func(c *gin.Context) {
		// Get uploaded file
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
			return
		}

		// Save file temporarily
		filePath := filepath.Join("frontend/uploads", file.Filename)
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
			return
		}
		defer os.Remove(filePath) // Clean up after processing

		// Step 1: Upload file to VirusTotal
		analysisID, err := uploadFileToVT(filePath, apiKey)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "VirusTotal upload failed: " + err.Error()})
			return
		}

		// Step 2: Wait for analysis to complete and get results
		results, err := getAnalysisResults(analysisID, apiKey)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get analysis results: " + err.Error()})
			return
		}

		// Return scan results
		c.JSON(http.StatusOK, gin.H{
			"filename": file.Filename,
			"scan_id":  analysisID,
			"results":  results,
		})
	})

	// Start server
	fmt.Println("Server running on http://localhost:8080")
	r.Run("0.0.0.0:8080")
}

func uploadFileToVT(filePath string, apiKey string) (string, error) {
	requestLock.Lock()
	if time.Since(lastRequestTime) < 15*time.Second { // Free tier - only 4 reqs/min
		time.Sleep(15*time.Second - time.Since(lastRequestTime))
	}
	lastRequestTime = time.Now()
	requestLock.Unlock()

	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		return "", err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return "", err
	}
	writer.Close()

	req, err := http.NewRequest("POST", "https://www.virustotal.com/api/v3/files", body)
	if err != nil {
		return "", err
	}

	req.Header.Set("x-apikey", apiKey)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("upload failed: %s", string(respBody))
	}

	var result struct {
		Data struct {
			ID string `json:"id"`
		} `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	return result.Data.ID, nil
}

func getAnalysisResults(analysisID string, apiKey string) (map[string]interface{}, error) {
	// Poll for results (max 2 minutes)
	for i := 0; i < 12; i++ {
		req, err := http.NewRequest("GET",
			fmt.Sprintf("https://www.virustotal.com/api/v3/analyses/%s", analysisID),
			nil)
		if err != nil {
			return nil, err
		}

		req.Header.Set("x-apikey", apiKey)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		var analysis struct {
			Data struct {
				Attributes struct {
					Status  string                 `json:"status"`
					Stats   map[string]int         `json:"stats"`
					Results map[string]interface{} `json:"results"`
				} `json:"attributes"`
			} `json:"data"`
		}

		if err := json.NewDecoder(resp.Body).Decode(&analysis); err != nil {
			return nil, err
		}

		if analysis.Data.Attributes.Status == "completed" {
			return map[string]interface{}{
				"status":  analysis.Data.Attributes.Status,
				"stats":   analysis.Data.Attributes.Stats,
				"results": analysis.Data.Attributes.Results,
			}, nil
		}

		time.Sleep(10 * time.Second)
	}

	return nil, fmt.Errorf("analysis timed out")
}
