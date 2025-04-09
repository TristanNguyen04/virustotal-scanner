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

