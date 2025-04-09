# VirusTotal File Scanner

This is a simple web application that allows users to upload files, scan them using VirusTotal’s API, and view the results directly in the browser. The backend is built using Go, and the frontend is a simple HTML/CSS interface. The application leverages the [VirusTotal API v3](https://www.virustotal.com/gui/home/upload) to check for malware using VirusTotal's detection engines.

## ✨ Features:
- **File upload**: Drag-and-drop or browse files to upload.
- **VirusTotal integration**: Files are scanned using the VirusTotal API.
- **Scan results**: View detailed results of the scan, including detection statistics.

## 🛠 Prerequisites

Before you can run this application, you need to do the following:

1. **Go (Golang)**: Ensure you have Go installed. You can download it from the official [Go website](https://golang.org/dl/).

2. **VirusTotal API Key**: You need an API key from VirusTotal to use their scan services.

### Steps to get a VirusTotal API Key:
1. Visit [VirusTotal's sign-in page](https://www.virustotal.com/gui/sign-in).
2. Sign up for a new account (or log in if you already have one).
3. Once signed in, go to your [API key page](https://www.virustotal.com/gui/account).
4. Copy your API key.

## ⚙️ Setup and Installation

### 1. Clone the repository

First, clone the repository to your local machine.

```bash
git clone https://github.com/TristanNguyen04/virustotal-scanner.git
git checkout local
cd virustotal-scanner
```

### 2. Install Go dependencies
Run the following command to install the necessary Go dependencies:
```bash
go mod tidy
```

### 3. Create `.env` file
Create a `.env` file in the folder and add your VirusTotal API key like this:
```text
VIRUSTOTAL_API_KEY="your_VirusTotal_API"
```

### 4. Run the backend
To run the backend server, use the following command:
```bash
go run main.go
```

This will start the server, and you should see the following message in the terminal:
```bash
Server running on http://localhost:8080
```

### 5. Access the web interface
Open you browser and navigate to [http://localhost:8080](http://localhost:8080). The frontend interface will allow you to upload files and view scan results.

### 6. File uploads
You can drag and drop a file into the "Drop Zone" or browse files to upload. The server will upload the file to VirusTotal for scanning, and you will receive a detailed scan result once the analysis is complete.

## 🖼 Project Structure
```
virustotal-scanner/
├── main.go                # The main Go file for the server
├── .env                   # Environment variables (for storing your API key)
├── frontend/                  # Frontend (HTML, CSS, JS) files
│   ├── index.html             # Main HTML interface
│   ├── styles.css             # Styling for the UI
│   ├── script.js              # JS for handling file uploads and API requests
│   └── uploads/               # Folder to temporarily store uploaded files
├── test_files/                # Folder containing test files you can upload to scan
└── README.md                  # This README file
```
## 🚀 Deployment on EC2
To make the server accessible online, I deployed it to an AWS EC2 Ubuntu instance:
**Stpes**:
1. SSH into the EC2 instance:
```bash
ssh -i virustotal-scanner-2.pem ubuntu@18.138.249.141
```
2. Install Go:
```bash
sudo apt update && sudo apt install golang -y
```
3. Clone the repo and follow the setup steps above.
4. To run the server even after closing terminal:
```bash
nohup go run main.go > server.log 2>&1 &
```
To see live logs:
```bash
tail -f server.log
```
To stop the server:
First find the process:
```bash
ps aux | grep main
```
Then kill it with:
```bash
kill <PID>
```
5. Allow traffic on port 8080 in my EC2 security group settings.
6. To access the hosted web application, visit [http://18.138.249.141:8080](http://18.138.249.141:8080)

## 🧪 Usage Instructions:
- Drag and drop or browse to upload a file
- The backend uploads it to VirusTotal and polls until results are ready
- The frontend shows a loader during scanning
- Once ready, it displays detection rates and scan details.

## ⚠️ Troubleshooting
- **API Key Issues**:  If you get a "missing API key" error, make sure the `.env` file contains the correct VirusTotal API key.
- **Go Errors**: If you get errors related to Go, ensure you have Go properly installed and that your environment is correctly set up for Go development. Use `go version` to check your Go installation.

## 💡 Challenges and Solutions
### 1. Rate Limiting
**Problem**: VirusTotal's free API allows only 4 requests per minute.
**Solution**: Added `time.Sleep()` delays and a polling mechanism to avoid overloading the API.

### 2. Long Scan Times
**Problem**: Large files or less common formats took longer to scan.
**Solution**: Implemented frontend feedback like a loading spinner and periodic polling for results.

### 3. `.env` File Handling in Go
**Problem**: Go doesn't natively support `.env` files.
**Solution**: Used a Go package to load `.env` configs at runtime (`github.com/joho/godotenv`).

### 4. Deployment Persistence
**Problem**: Server would stop after logging out from SSH.
**Solution**: Used `nohup` to keep the Go server alive after disconnecting from terminal.

## 🌟 Enhancements
- ✅ **Async Polling**: Improved user experience by updating scan results live
- ✅ **Drag-and-Drop Upload UI**
- 🚧 *Future Idea*: Implement user authentication and scan history
- 🚧 *Future Idea*: Add virus name highlights and color-coded severity scores

## ✅ Evaluation Criteria Mapping
- **Functionality**: Fully working file scan, real-time results, VirusTotal integration
- **Code Quality**: Modular backend, separate frontend, `.env` config, reusable functions
- **Problem-Solving**:	Solved issues with `.env`, rate limits, server persistence
- **Creativity**:	Added drag-and-drop, live polling, EC2 deployment
- **Presentation**:	Polished UI, detailed logs, well-documented README

## 📝 Additional Notes
- **Free Tier Limits**:  VirusTotal's free tier has limitations. You can only make a limited number of requests per minute (4 requests per minute). The backend implements a rate-limiting feature to handle this limitation by waiting 15 seconds between requests.
- **File Size**: VirusTotal has a file size limit of 32 MB for the free tier. Make sure the files you are uploading are under this size limit.
- **Frontend Updates**: The frontend will show you a loading spinner while the scan is in progress. Once the analysis is complete, you’ll see detailed results, including detection statistics and detection details.

## 📎 Acknowledgement
This project was completed as part of the **CloudsineAI WebTest Take-Home Assignment**.
Thanks to the VirusTotal team for their comprehensive and free malware scanning API.
