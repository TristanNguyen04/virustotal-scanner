# VirusTotal File Scanner

This is a simple web application that allows users to upload files, scan them using VirusTotal’s API, and view the results directly in the browser. The backend is built using Go, and the frontend is a simple HTML/CSS interface. The application leverages the [VirusTotal API v3](https://www.virustotal.com/gui/home/upload) to check for malware using VirusTotal's detection engines.

## Features:
- **File upload**: Drag-and-drop or browse files to upload.
- **VirusTotal integration**: Files are scanned using the VirusTotal API.
- **Scan results**: View detailed results of the scan, including detection statistics.

## Prerequisites

Before you can run this application, you need to do the following:

1. **Go (Golang)**: Ensure you have Go installed. You can download it from the official [Go website](https://golang.org/dl/).

2. **VirusTotal API Key**: You need an API key from VirusTotal to use their scan services.

### Steps to get a VirusTotal API Key:
1. Visit [VirusTotal's sign-in page](https://www.virustotal.com/gui/sign-in).
2. Sign up for a new account (or log in if you already have one).
3. Once signed in, go to your [API key page](https://www.virustotal.com/gui/account).
4. Copy your API key.

## Setup and Installation

### 1. Clone the repository

First, clone the repository to your local machine.

```bash
git clone <your-repository-url>
cd virustotal-scanner
```

### 2. Clone the repository
Navigate to the `backend` directory and run the following command to install the necessary Go dependencies:
```bash
cd backend
go mod tidy
```

### 3. Create `.env` file
Create a `.env` file in the `backend` folder and add your VirusTotal API key like this:
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
Open you browser and navigate to `http://localhost:8080`. The frontend interface will allow you to upload files and view scan results.

### 6. File uploads
You can drag and drop a file into the "Drop Zone" or browse files to upload. The server will upload the file to VirusTotal for scanning, and you will receive a detailed scan result once the analysis is complete.

## Project Structure
```
virustotal-scanner/
├── backend/                   # Backend (Go) logic
│   ├── main.go                # The main Go file for the server
│   ├── .env                   # Environment variables (for storing your API key)
├── frontend/                  # Frontend (HTML, CSS, JS) files
│   ├── index.html             # Main HTML interface
│   ├── styles.css             # Styling for the UI
│   ├── script.js              # JS for handling file uploads and API requests
│   └── uploads/               # Folder to temporarily store uploaded files
├── test_files/                # Folder containing test files you can upload to scan
└── README.md                  # This README file
```








