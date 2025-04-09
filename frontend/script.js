document.addEventListener('DOMContentLoaded', function() {
    const fileInput = document.getElementById('fileInput');
    const dropZone = document.getElementById('dropZone');
    const fileInfo = document.getElementById('fileInfo');
    const scanBtn = document.getElementById('scanBtn');
    const loadingIndicator = document.getElementById('loadingIndicator');
    const resultsContainer = document.getElementById('results');
    const resultsDiv = document.getElementById('resultsContainer');
    
    let selectedFile = null;

    // Drag and drop functionality
    dropZone.addEventListener('dragover', (e) => {
        e.preventDefault();
        dropZone.classList.add('drag-over');
    });

    ['dragleave', 'dragend'].forEach(type => {
        dropZone.addEventListener(type, () => {
            dropZone.classList.remove('drag-over');
        });
    });

    dropZone.addEventListener('drop', (e) => {
        e.preventDefault();
        dropZone.classList.remove('drag-over');
        
        if (e.dataTransfer.files.length) {
            fileInput.files = e.dataTransfer.files;
            handleFileSelection(e.dataTransfer.files[0]);
        }
    });

    // File input change handler
    fileInput.addEventListener('change', () => {
        if (fileInput.files.length) {
            handleFileSelection(fileInput.files[0]);
        }
    });

    // Scan button click handler
    scanBtn.addEventListener('click', scanFile);

    function handleFileSelection(file) {
        // Validate file size (32MB limit for VirusTotal free tier)
        if (file.size > 32 * 1024 * 1024) {
            showError('File size exceeds 32MB limit for free tier');
            return;
        }

        selectedFile = file;
        fileInfo.innerHTML = `
            <strong>Selected file:</strong> ${file.name}<br>
            <strong>Size:</strong> ${formatFileSize(file.size)}
        `;
        fileInfo.style.display = 'block';
        scanBtn.disabled = false;
    }

    async function scanFile() {
        if (!selectedFile) return;

        // Show loading state
        loadingIndicator.style.display = 'flex';
        resultsContainer.style.display = 'none';
        scanBtn.disabled = true;

        const formData = new FormData();
        formData.append('file', selectedFile);

        try {
            const response = await fetch('/api/upload', {
                method: 'POST',
                body: formData
            });

            if (!response.ok) {
                throw new Error(await response.text());
            }

            const data = await response.json();
            displayResults(data);
        } catch (error) {
            showError(error.message);
        } finally {
            loadingIndicator.style.display = 'none';
        }
    }

    function displayResults(data) {
        resultsContainer.innerHTML = `
            <div class="result-header">
                <h3>Scan Results for <span class="filename">${data.filename}</span></h3>
                <small>Scan ID: ${data.scan_id}</small>
            </div>
            
            <div class="stats">
                <div class="stat-card malicious">
                    <h4>${data.results.stats.malicious || 0}</h4>
                    <p>Malicious</p>
                </div>
                <div class="stat-card suspicious">
                    <h4>${data.results.stats.suspicious || 0}</h4>
                    <p>Suspicious</p>
                </div>
                <div class="stat-card undetected">
                    <h4>${data.results.stats.undetected || 0}</h4>
                    <p>Undetected</p>
                </div>
            </div>
            
            <h4>Detection Details:</h4>
            <div class="detections">
                ${Object.entries(data.results.results || {}).map(([engine, result]) => `
                    <div class="detection">
                        <span class="detection-name">${engine}</span>
                        <span class="detection-result ${getDetectionClass(result.category)}">
                            ${result.category || 'undetected'}
                        </span>
                    </div>
                `).join('')}
            </div>
        `;
        
        resultsContainer.style.display = 'block';
    }

    function getDetectionClass(category) {
        if (!category) return 'detection-undetected';
        return category === 'malicious' ? 'detection-malicious' :
               category === 'suspicious' ? 'detection-suspicious' :
               'detection-undetected';
    }

    function showError(message) {
        resultsContainer.innerHTML = `
            <div class="error">
                <h3>Error</h3>
                <p>${message}</p>
            </div>
        `;
        resultsContainer.style.display = 'block';
    }

    function formatFileSize(bytes) {
        if (bytes === 0) return '0 Bytes';
        const k = 1024;
        const sizes = ['Bytes', 'KB', 'MB', 'GB'];
        const i = Math.floor(Math.log(bytes) / Math.log(k));
        return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i];
    }
});