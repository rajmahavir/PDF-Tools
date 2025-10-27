package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

const maxUploadSize = 50 * 1024 * 1024 // 50 MB

func main() {
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/merge", handleMergePage)
	http.HandleFunc("/remove", handleRemovePage)
	http.HandleFunc("/credits", handleCreditsPage)
	http.HandleFunc("/merge-pdfs", handleMerge)
	http.HandleFunc("/remove-pages", handleRemovePages)
	http.HandleFunc("/pdfinfo", handlePDFInfo)

	port := "8080"
	
	fmt.Printf("Server starting on:\n")
	fmt.Printf("  Local:   http://localhost:%s\n", port)
	fmt.Printf("  Network: http://[YOUR-IP-ADDRESS]:%s\n", port)
	fmt.Printf("\nReplace [YOUR-IP-ADDRESS] with your computer's IP address\n")
	fmt.Printf("To find it on Windows, run: ipconfig\n\n")
	
	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, nil))
}

func getFooterHTML() string {
	return `
	<footer style="margin-top: 40px; padding-top: 20px; border-top: 1px solid #eee; text-align: center; color: #666; font-size: 13px;">
		<div style="margin-bottom: 10px;">
			<span style="display: inline-block; margin: 0 10px;">🤖 Built with <a href="https://claude.ai" target="_blank" style="color: #667eea; text-decoration: none;">Claude.ai</a></span>
			<span style="display: inline-block; margin: 0 10px;">⚡ Powered by <a href="https://github.com/pdfcpu/pdfcpu" target="_blank" style="color: #667eea; text-decoration: none;">pdfcpu</a></span>
		</div>
		<div>
			<a href="/credits" style="color: #667eea; text-decoration: none; margin: 0 10px;">Credits</a>
			<span style="color: #ddd;">|</span>
			<a href="https://github.com/rajmahavir/PDF-Tools" target="_blank" style="color: #667eea; text-decoration: none; margin: 0 10px;">Source Code</a>
			<span style="color: #ddd;">|</span>
			<span style="margin: 0 10px;">MIT License</span>
		</div>
		<div style="margin-top: 10px; font-size: 12px; color: #999;">
			Free & Open Source PDF Tools
		</div>
	</footer>
	`
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	html := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>PDF Tools - Free Online PDF Editor</title>
    <style>
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }
        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, sans-serif;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            min-height: 100vh;
            display: flex;
            align-items: center;
            justify-content: center;
            padding: 20px;
        }
        .container {
            background: white;
            border-radius: 20px;
            box-shadow: 0 20px 60px rgba(0,0,0,0.3);
            padding: 40px;
            max-width: 700px;
            width: 100%;
        }
        h1 {
            color: #333;
            margin-bottom: 10px;
            font-size: 32px;
            text-align: center;
        }
        .subtitle {
            color: #666;
            margin-bottom: 40px;
            font-size: 14px;
            text-align: center;
        }
        .options-grid {
            display: grid;
            grid-template-columns: 1fr 1fr;
            gap: 20px;
            margin-top: 20px;
        }
        .option-card {
            background: linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%);
            border-radius: 15px;
            padding: 30px 20px;
            text-align: center;
            cursor: pointer;
            transition: all 0.3s;
            border: 2px solid transparent;
        }
        .option-card:hover {
            transform: translateY(-5px);
            box-shadow: 0 10px 30px rgba(0,0,0,0.2);
            border-color: #667eea;
        }
        .option-card.remove {
            background: linear-gradient(135deg, #ffecd2 0%, #fcb69f 100%);
        }
        .option-card.merge {
            background: linear-gradient(135deg, #a8edea 0%, #fed6e3 100%);
        }
        .option-icon {
            font-size: 48px;
            margin-bottom: 15px;
        }
        .option-title {
            font-size: 20px;
            font-weight: 600;
            color: #333;
            margin-bottom: 10px;
        }
        .option-desc {
            font-size: 13px;
            color: #666;
        }
        @media (max-width: 600px) {
            .options-grid {
                grid-template-columns: 1fr;
            }
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>🔧 PDF Tools</h1>
        <p class="subtitle">Free & Open Source PDF Editor</p>
        
        <div class="options-grid">
            <div class="option-card remove" onclick="window.location.href='/remove'">
                <div class="option-icon">✂️</div>
                <div class="option-title">Remove Pages</div>
                <div class="option-desc">Delete specific pages from your PDF or split it into parts</div>
            </div>
            
            <div class="option-card merge" onclick="window.location.href='/merge'">
                <div class="option-icon">📄</div>
                <div class="option-title">Merge PDFs</div>
                <div class="option-desc">Insert all pages from one PDF into another at any position</div>
            </div>
        </div>
        ` + getFooterHTML() + `
    </div>
</body>
</html>`
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html))
}

func handleCreditsPage(w http.ResponseWriter, r *http.Request) {
	html := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Credits - PDF Tools</title>
    <style>
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }
        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, sans-serif;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            min-height: 100vh;
            padding: 40px 20px;
        }
        .container {
            background: white;
            border-radius: 20px;
            box-shadow: 0 20px 60px rgba(0,0,0,0.3);
            padding: 40px;
            max-width: 800px;
            margin: 0 auto;
        }
        .back-btn {
            display: inline-block;
            color: #667eea;
            text-decoration: none;
            margin-bottom: 20px;
            font-size: 14px;
            font-weight: 500;
        }
        .back-btn:hover {
            text-decoration: underline;
        }
        h1 {
            color: #333;
            margin-bottom: 10px;
            font-size: 32px;
        }
        .subtitle {
            color: #666;
            margin-bottom: 30px;
            font-size: 16px;
        }
        .section {
            margin-bottom: 30px;
            padding: 20px;
            background: #f8f9fa;
            border-radius: 10px;
            border-left: 4px solid #667eea;
        }
        .section h2 {
            color: #333;
            font-size: 20px;
            margin-bottom: 15px;
        }
        .section h3 {
            color: #555;
            font-size: 16px;
            margin-top: 15px;
            margin-bottom: 10px;
        }
        .section p {
            color: #666;
            line-height: 1.6;
            margin-bottom: 10px;
        }
        .credit-item {
            background: white;
            padding: 15px;
            border-radius: 8px;
            margin-bottom: 15px;
        }
        .credit-item strong {
            color: #667eea;
            display: block;
            margin-bottom: 5px;
        }
        .credit-item a {
            color: #667eea;
            text-decoration: none;
        }
        .credit-item a:hover {
            text-decoration: underline;
        }
        .badge {
            display: inline-block;
            padding: 5px 12px;
            background: #e8f0fe;
            color: #1967d2;
            border-radius: 20px;
            font-size: 12px;
            margin: 5px 5px 5px 0;
            font-weight: 500;
        }
        .badge.ai {
            background: #fce8f3;
            color: #c2185b;
        }
        .badge.license {
            background: #e8f5e9;
            color: #2e7d32;
        }
        ul {
            margin-left: 20px;
            color: #666;
            line-height: 1.8;
        }
        .hero {
            text-align: center;
            padding: 20px;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            color: white;
            border-radius: 10px;
            margin-bottom: 30px;
        }
        .hero h1 {
            color: white;
            font-size: 36px;
        }
        .hero p {
            color: rgba(255,255,255,0.9);
            font-size: 18px;
        }
    </style>
</head>
<body>
    <div class="container">
        <a href="/" class="back-btn">← Back to Home</a>
        
        <div class="hero">
            <h1>🔧 PDF Tools</h1>
            <p>Free & Open Source</p>
        </div>

        <div class="section">
            <h2>💡 About This Project</h2>
            <p>
                PDF Tools is a free, open-source web application for merging and editing PDF files. 
                Built with modern technology and AI assistance, this tool provides powerful PDF 
                manipulation capabilities accessible to everyone, completely free of charge.
            </p>
            <div style="margin-top: 15px;">
                <span class="badge license">MIT License</span>
                <span class="badge">Open Source</span>
                <span class="badge">No Registration Required</span>
                <span class="badge">Privacy Focused</span>
            </div>
        </div>

        <div class="section">
            <h2>🤖 AI-Assisted Development</h2>
            
            <div class="credit-item">
                <strong>Claude.ai by Anthropic</strong>
                <span class="badge ai">AI Assistant</span>
                <p>
                    This application was developed with significant assistance from Claude.ai, 
                    an AI assistant created by Anthropic. The collaboration involved architecture 
                    design, code generation, user interface development, and implementation of 
                    PDF processing features.
                </p>
                <p>
                    <strong>Model:</strong> Claude Sonnet 4.5<br>
                    <strong>Website:</strong> <a href="https://claude.ai" target="_blank">https://claude.ai</a><br>
                    <strong>Company:</strong> Anthropic PBC
                </p>
            </div>

            <div class="credit-item">
                <strong>Human Developer</strong>
                <p>
                    Project direction, testing, customization, and deployment managed by the human developer.
                    The final implementation represents a collaborative effort between human creativity 
                    and AI capabilities.
                </p>
            </div>
        </div>

        <div class="section">
            <h2>🔧 Technology Stack</h2>
            
            <div class="credit-item">
                <strong>pdfcpu</strong>
                <span class="badge license">Apache License 2.0</span>
                <p>
                    A powerful PDF processing library written in Go. This is the core technology 
                    that powers all PDF manipulation features in this application.
                </p>
                <p>
                    <strong>Author:</strong> Horst Rutter<br>
                    <strong>Repository:</strong> <a href="https://github.com/pdfcpu/pdfcpu" target="_blank">github.com/pdfcpu/pdfcpu</a><br>
                    <strong>License:</strong> Apache License 2.0
                </p>
            </div>

            <div class="credit-item">
                <strong>Go Programming Language</strong>
                <span class="badge license">BSD 3-Clause</span>
                <p>
                    The backend server and PDF operations are built using Go (Golang), 
                    a fast, reliable, and efficient programming language.
                </p>
                <p>
                    <strong>Website:</strong> <a href="https://golang.org" target="_blank">https://golang.org</a><br>
                    <strong>License:</strong> BSD 3-Clause License
                </p>
            </div>
        </div>

        <div class="section">
            <h2>📄 License Information</h2>
            <h3>This Application (PDF Tools)</h3>
            <p>
                <strong>License:</strong> MIT License<br>
                <strong>Copyright:</strong> © 2025
            </p>
            <p>
                Permission is hereby granted, free of charge, to any person obtaining a copy 
                of this software and associated documentation files, to deal in the Software 
                without restriction, including without limitation the rights to use, copy, 
                modify, merge, publish, distribute, sublicense, and/or sell copies of the Software.
            </p>
            
            <h3>Third-Party Licenses</h3>
            <ul>
                <li><strong>pdfcpu:</strong> Apache License 2.0 - Requires attribution and license inclusion</li>
                <li><strong>Go:</strong> BSD 3-Clause - Permissive open source license</li>
            </ul>
        </div>

        <div class="section">
            <h2>🔒 Privacy & Security</h2>
            <p><strong>Your privacy is our priority:</strong></p>
            <ul>
                <li>All PDF processing happens on the server temporarily</li>
                <li>Files are automatically deleted immediately after processing</li>
                <li>We do not store, access, or analyze your PDF content</li>
                <li>No user data collection or tracking</li>
                <li>No cookies or analytics</li>
                <li>No registration or login required</li>
            </ul>
        </div>

        <div class="section">
            <h2>⚖️ Disclaimer</h2>
            <p>
                This tool is provided "as is" without warranty of any kind, express or implied. 
                The developers and contributors are not responsible for any data loss, corruption, 
                or issues arising from the use of this service. Use at your own risk.
            </p>
            <p>
                While we take precautions to ensure file security and privacy, users should not 
                upload sensitive or confidential documents to any online service without proper 
                risk assessment.
            </p>
        </div>

        <div class="section">
            <h2>📦 Source Code</h2>
            <p>
                This project is open source and available on GitHub. You can view the code, 
                report issues, or contribute improvements.
            </p>
            <p>
                <strong>Repository:</strong> <a href="https://github.com/rajmahavir/PDF-Tools" target="_blank">github.com/rajmahavir/PDF-Tools</a>
            </p>
            <p style="margin-top: 15px;">
                <strong>How to Contribute:</strong>
            </p>
            <ul>
                <li>Report bugs or suggest features via GitHub Issues</li>
                <li>Submit pull requests for improvements</li>
                <li>Share the project with others who might find it useful</li>
                <li>Star the repository to show your support</li>
            </ul>
        </div>

        <div class="section">
            <h2>🙏 Acknowledgments</h2>
            <p>Special thanks to:</p>
            <ul>
                <li><strong>Anthropic</strong> - For creating Claude.ai and enabling AI-assisted development</li>
                <li><strong>Horst Rutter</strong> - For developing and maintaining pdfcpu</li>
                <li><strong>Go Team</strong> - For the excellent Go programming language</li>
                <li><strong>Open Source Community</strong> - For making projects like this possible</li>
            </ul>
        </div>

        <div style="text-align: center; margin-top: 40px; padding-top: 20px; border-top: 2px solid #eee;">
            <p style="color: #999; font-size: 14px;">
                Made with ❤️ using Claude.ai<br>
                © 2025 PDF Tools • Free Forever
            </p>
        </div>
    </div>
</body>
</html>`
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html))
}

func handleRemovePage(w http.ResponseWriter, r *http.Request) {
	html := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Remove Pages - PDF Tools</title>
    <style>
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }
        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, sans-serif;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            min-height: 100vh;
            display: flex;
            align-items: center;
            justify-content: center;
            padding: 20px;
        }
        .container {
            background: white;
            border-radius: 20px;
            box-shadow: 0 20px 60px rgba(0,0,0,0.3);
            padding: 40px;
            max-width: 600px;
            width: 100%;
        }
        .back-btn {
            display: inline-block;
            color: #667eea;
            text-decoration: none;
            margin-bottom: 20px;
            font-size: 14px;
            font-weight: 500;
        }
        .back-btn:hover {
            text-decoration: underline;
        }
        h1 {
            color: #333;
            margin-bottom: 10px;
            font-size: 28px;
        }
        .subtitle {
            color: #666;
            margin-bottom: 30px;
            font-size: 14px;
        }
        .upload-section {
            margin-bottom: 25px;
        }
        label {
            display: block;
            margin-bottom: 8px;
            color: #333;
            font-weight: 500;
            font-size: 14px;
        }
        .file-input-wrapper {
            position: relative;
            overflow: hidden;
            display: inline-block;
            width: 100%;
        }
        .file-input-wrapper input[type=file] {
            position: absolute;
            left: -9999px;
        }
        .file-input-label {
            display: flex;
            align-items: center;
            justify-content: center;
            padding: 15px;
            background: #f8f9fa;
            border: 2px dashed #ddd;
            border-radius: 10px;
            cursor: pointer;
            transition: all 0.3s;
        }
        .file-input-label:hover {
            background: #e9ecef;
            border-color: #667eea;
        }
        .file-name {
            margin-top: 8px;
            font-size: 13px;
            color: #666;
            font-style: italic;
        }
        .pdf-info {
            margin-top: 10px;
            padding: 10px;
            background: #f0f7ff;
            border: 1px solid #b3d9ff;
            border-radius: 8px;
            font-size: 13px;
            display: none;
        }
        .pdf-info.visible {
            display: block;
        }
        .pdf-info strong {
            color: #0066cc;
        }
        .pdf-preview {
            margin-top: 10px;
            border: 2px solid #ddd;
            border-radius: 8px;
            overflow: hidden;
            display: none;
            max-height: 300px;
        }
        .pdf-preview.visible {
            display: block;
        }
        .pdf-preview iframe {
            width: 100%;
            height: 300px;
            border: none;
        }
        .pages-to-remove {
            margin-top: 20px;
            display: none;
        }
        .pages-to-remove.visible {
            display: block;
        }
        .page-selector {
            display: flex;
            flex-wrap: wrap;
            gap: 8px;
            margin-top: 10px;
            max-height: 200px;
            overflow-y: auto;
            padding: 10px;
            background: #f8f9fa;
            border-radius: 8px;
        }
        .page-checkbox {
            display: flex;
            align-items: center;
            gap: 5px;
            padding: 8px 12px;
            background: white;
            border: 2px solid #ddd;
            border-radius: 6px;
            cursor: pointer;
            transition: all 0.2s;
            user-select: none;
        }
        .page-checkbox:hover {
            border-color: #667eea;
        }
        .page-checkbox input[type="checkbox"] {
            cursor: pointer;
        }
        .page-checkbox.selected {
            background: #667eea;
            border-color: #667eea;
            color: white;
        }
        .hint {
            font-size: 12px;
            color: #999;
            margin-top: 5px;
        }
        .action-buttons {
            display: flex;
            gap: 10px;
            margin-top: 20px;
        }
        button {
            flex: 1;
            padding: 15px;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            color: white;
            border: none;
            border-radius: 10px;
            font-size: 16px;
            font-weight: 600;
            cursor: pointer;
            transition: transform 0.2s, box-shadow 0.2s;
        }
        button:hover {
            transform: translateY(-2px);
            box-shadow: 0 10px 20px rgba(102, 126, 234, 0.4);
        }
        button:active {
            transform: translateY(0);
        }
        button:disabled {
            opacity: 0.6;
            cursor: not-allowed;
            transform: none;
        }
        .select-all-btn {
            background: linear-gradient(135deg, #11998e 0%, #38ef7d 100%);
            margin-bottom: 10px;
        }
        .error {
            background: #fee;
            border: 1px solid #fcc;
            color: #c33;
            padding: 12px;
            border-radius: 10px;
            margin-top: 15px;
            font-size: 14px;
        }
        .success {
            background: #efe;
            border: 1px solid #cfc;
            color: #3c3;
            padding: 12px;
            border-radius: 10px;
            margin-top: 15px;
            font-size: 14px;
        }
        #result {
            margin-top: 20px;
        }
        #pdfViewer {
            width: 100%;
            height: 600px;
            border: 2px solid #ddd;
            border-radius: 10px;
            margin-top: 15px;
        }
        .download-btn {
            background: linear-gradient(135deg, #11998e 0%, #38ef7d 100%);
            margin-top: 10px;
        }
        .loading {
            display: none;
            text-align: center;
            margin-top: 20px;
        }
        .spinner {
            border: 3px solid #f3f3f3;
            border-top: 3px solid #667eea;
            border-radius: 50%;
            width: 40px;
            height: 40px;
            animation: spin 1s linear infinite;
            margin: 0 auto;
        }
        @keyframes spin {
            0% { transform: rotate(0deg); }
            100% { transform: rotate(360deg); }
        }
    </style>
</head>
<body>
    <div class="container">
        <a href="/" class="back-btn">← Back to Home</a>
        <h1>✂️ Remove Pages</h1>
        <p class="subtitle">Select pages to remove from your PDF</p>
        
        <form id="uploadForm" enctype="multipart/form-data">
            <div class="upload-section">
                <label for="pdf">Upload PDF *</label>
                <div class="file-input-wrapper">
                    <input type="file" id="pdf" name="pdf" accept=".pdf" required>
                    <label for="pdf" class="file-input-label">
                        Choose PDF file
                    </label>
                </div>
                <div id="fileName" class="file-name"></div>
                <div id="pdfInfo" class="pdf-info"></div>
                <div id="pdfPreview" class="pdf-preview"></div>
            </div>

            <div id="pageSelectorSection" class="pages-to-remove">
                <label>Select pages to remove:</label>
                <div class="hint">Click on pages you want to delete from the PDF</div>
                <button type="button" class="select-all-btn" onclick="toggleSelectAll()">Select All</button>
                <div id="pageSelector" class="page-selector"></div>
                <div class="action-buttons">
                    <button type="submit">Remove Selected Pages</button>
                </div>
            </div>
        </form>

        <div class="loading" id="loading">
            <div class="spinner"></div>
            <p style="margin-top: 10px; color: #666;">Processing PDF...</p>
        </div>

        <div id="result"></div>
        ` + getFooterHTML() + `
    </div>

    <script>
        var pdfInput = document.getElementById('pdf');
        var fileName = document.getElementById('fileName');
        var pdfInfoDiv = document.getElementById('pdfInfo');
        var pdfPreviewDiv = document.getElementById('pdfPreview');
        var pageSelectorSection = document.getElementById('pageSelectorSection');
        var pageSelector = document.getElementById('pageSelector');
        var currentPDFFile = null;
        var totalPages = 0;
        var selectedPages = [];

        pdfInput.addEventListener('change', function(e) {
            if (e.target.files.length > 0) {
                currentPDFFile = e.target.files[0];
                fileName.textContent = 'Selected: ' + currentPDFFile.name;
                loadPDFInfo(currentPDFFile);
            }
        });

        function loadPDFInfo(file) {
            var formData = new FormData();
            formData.append('pdf', file);

            pdfInfoDiv.innerHTML = 'Loading info...';
            pdfInfoDiv.classList.add('visible');
            pdfPreviewDiv.classList.remove('visible');
            pageSelectorSection.classList.remove('visible');

            fetch('/pdfinfo', {
                method: 'POST',
                body: formData
            }).then(function(response) {
                return response.json();
            }).then(function(data) {
                if (data.error) {
                    pdfInfoDiv.innerHTML = 'Error: ' + data.error;
                } else {
                    totalPages = data.pageCount;
                    pdfInfoDiv.innerHTML = '<strong>Pages:</strong> ' + data.pageCount + 
                                          ' | <strong>Size:</strong> ' + formatFileSize(file.size);
                    
                    var url = URL.createObjectURL(file);
                    pdfPreviewDiv.innerHTML = '<iframe src="' + url + '#view=FitH"></iframe>';
                    pdfPreviewDiv.classList.add('visible');
                    
                    createPageSelector(data.pageCount);
                    pageSelectorSection.classList.add('visible');
                }
            }).catch(function(error) {
                pdfInfoDiv.innerHTML = 'Error loading PDF info';
            });
        }

        function createPageSelector(pageCount) {
            pageSelector.innerHTML = '';
            selectedPages = [];
            
            for (var i = 1; i <= pageCount; i++) {
                var pageDiv = document.createElement('div');
                pageDiv.className = 'page-checkbox';
                pageDiv.setAttribute('data-page', i);
                
                var checkbox = document.createElement('input');
                checkbox.type = 'checkbox';
                checkbox.id = 'page-' + i;
                checkbox.value = i;
                
                var label = document.createElement('label');
                label.setAttribute('for', 'page-' + i);
                label.textContent = 'Page ' + i;
                label.style.cursor = 'pointer';
                
                pageDiv.appendChild(checkbox);
                pageDiv.appendChild(label);
                
                pageDiv.addEventListener('click', function(e) {
                    var checkbox = this.querySelector('input[type="checkbox"]');
                    if (e.target !== checkbox) {
                        checkbox.checked = !checkbox.checked;
                    }
                    
                    if (checkbox.checked) {
                        this.classList.add('selected');
                        selectedPages.push(parseInt(checkbox.value));
                    } else {
                        this.classList.remove('selected');
                        var index = selectedPages.indexOf(parseInt(checkbox.value));
                        if (index > -1) {
                            selectedPages.splice(index, 1);
                        }
                    }
                });
                
                pageSelector.appendChild(pageDiv);
            }
        }

        function toggleSelectAll() {
            var checkboxes = pageSelector.querySelectorAll('input[type="checkbox"]');
            var allSelected = selectedPages.length === totalPages;
            
            selectedPages = [];
            checkboxes.forEach(function(checkbox) {
                checkbox.checked = !allSelected;
                var pageDiv = checkbox.parentElement;
                if (!allSelected) {
                    pageDiv.classList.add('selected');
                    selectedPages.push(parseInt(checkbox.value));
                } else {
                    pageDiv.classList.remove('selected');
                }
            });
        }

        function formatFileSize(bytes) {
            if (bytes < 1024) return bytes + ' B';
            if (bytes < 1024 * 1024) return (bytes / 1024).toFixed(1) + ' KB';
            return (bytes / (1024 * 1024)).toFixed(1) + ' MB';
        }

        document.getElementById('uploadForm').addEventListener('submit', function(e) {
            e.preventDefault();
            
            if (selectedPages.length === 0) {
                alert('Please select at least one page to remove');
                return;
            }
            
            if (selectedPages.length === totalPages) {
                alert('You cannot remove all pages. At least one page must remain.');
                return;
            }
            
            var formData = new FormData();
            formData.append('pdf', currentPDFFile);
            formData.append('pagesToRemove', selectedPages.sort(function(a, b) { return a - b; }).join(','));
            
            var resultDiv = document.getElementById('result');
            var loadingDiv = document.getElementById('loading');
            var submitBtn = e.target.querySelector('button[type="submit"]');
            
            resultDiv.innerHTML = '';
            loadingDiv.style.display = 'block';
            submitBtn.disabled = true;

            fetch('/remove-pages', {
                method: 'POST',
                body: formData
            }).then(function(response) {
                loadingDiv.style.display = 'none';
                submitBtn.disabled = false;

                if (response.ok) {
                    return response.blob().then(function(blob) {
                        var url = URL.createObjectURL(blob);
                        
                        resultDiv.innerHTML = 
                            '<div class="success">Pages removed successfully! Remaining pages: ' + (totalPages - selectedPages.length) + '</div>' +
                            '<iframe id="pdfViewer" src="' + url + '"></iframe>' +
                            '<button class="download-btn" onclick="downloadPDF(\'' + url + '\')">Download Modified PDF</button>';
                    });
                } else {
                    return response.text().then(function(error) {
                        resultDiv.innerHTML = '<div class="error">Error: ' + error + '</div>';
                    });
                }
            }).catch(function(error) {
                loadingDiv.style.display = 'none';
                submitBtn.disabled = false;
                resultDiv.innerHTML = '<div class="error">Error: ' + error.message + '</div>';
            });
        });

        function downloadPDF(url) {
            var a = document.createElement('a');
            a.href = url;
            a.download = 'modified_' + Date.now() + '.pdf';
            document.body.appendChild(a);
            a.click();
            document.body.removeChild(a);
        }
    </script>
</body>
</html>`
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html))
}

func handleMergePage(w http.ResponseWriter, r *http.Request) {
	html := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Merge PDFs - PDF Tools</title>
    <style>
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }
        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, sans-serif;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            min-height: 100vh;
            display: flex;
            align-items: center;
            justify-content: center;
            padding: 20px;
        }
        .container {
            background: white;
            border-radius: 20px;
            box-shadow: 0 20px 60px rgba(0,0,0,0.3);
            padding: 40px;
            max-width: 600px;
            width: 100%;
        }
        .back-btn {
            display: inline-block;
            color: #667eea;
            text-decoration: none;
            margin-bottom: 20px;
            font-size: 14px;
            font-weight: 500;
        }
        .back-btn:hover {
            text-decoration: underline;
        }
        h1 {
            color: #333;
            margin-bottom: 10px;
            font-size: 28px;
        }
        .subtitle {
            color: #666;
            margin-bottom: 30px;
            font-size: 14px;
        }
        .upload-section {
            margin-bottom: 25px;
        }
        label {
            display: block;
            margin-bottom: 8px;
            color: #333;
            font-weight: 500;
            font-size: 14px;
        }
        .file-input-wrapper {
            position: relative;
            overflow: hidden;
            display: inline-block;
            width: 100%;
        }
        .file-input-wrapper input[type=file] {
            position: absolute;
            left: -9999px;
        }
        .file-input-label {
            display: flex;
            align-items: center;
            justify-content: center;
            padding: 15px;
            background: #f8f9fa;
            border: 2px dashed #ddd;
            border-radius: 10px;
            cursor: pointer;
            transition: all 0.3s;
        }
        .file-input-label:hover {
            background: #e9ecef;
            border-color: #667eea;
        }
        .file-name {
            margin-top: 8px;
            font-size: 13px;
            color: #666;
            font-style: italic;
        }
        .pdf-info {
            margin-top: 10px;
            padding: 10px;
            background: #f0f7ff;
            border: 1px solid #b3d9ff;
            border-radius: 8px;
            font-size: 13px;
            display: none;
        }
        .pdf-info.visible {
            display: block;
        }
        .pdf-info strong {
            color: #0066cc;
        }
        .pdf-preview {
            margin-top: 10px;
            border: 2px solid #ddd;
            border-radius: 8px;
            overflow: hidden;
            display: none;
            max-height: 200px;
        }
        .pdf-preview.visible {
            display: block;
        }
        .pdf-preview iframe {
            width: 100%;
            height: 200px;
            border: none;
        }
        input[type="number"] {
            width: 100%;
            padding: 12px;
            border: 2px solid #ddd;
            border-radius: 10px;
            font-size: 16px;
            transition: border-color 0.3s;
        }
        input[type="number"]:focus {
            outline: none;
            border-color: #667eea;
        }
        .hint {
            font-size: 12px;
            color: #999;
            margin-top: 5px;
        }
        button {
            width: 100%;
            padding: 15px;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            color: white;
            border: none;
            border-radius: 10px;
            font-size: 16px;
            font-weight: 600;
            cursor: pointer;
            transition: transform 0.2s, box-shadow 0.2s;
            margin-top: 10px;
        }
        button:hover {
            transform: translateY(-2px);
            box-shadow: 0 10px 20px rgba(102, 126, 234, 0.4);
        }
        button:active {
            transform: translateY(0);
        }
        button:disabled {
            opacity: 0.6;
            cursor: not-allowed;
            transform: none;
        }
        .error {
            background: #fee;
            border: 1px solid #fcc;
            color: #c33;
            padding: 12px;
            border-radius: 10px;
            margin-top: 15px;
            font-size: 14px;
        }
        .success {
            background: #efe;
            border: 1px solid #cfc;
            color: #3c3;
            padding: 12px;
            border-radius: 10px;
            margin-top: 15px;
            font-size: 14px;
        }
        #result {
            margin-top: 20px;
        }
        #pdfViewer {
            width: 100%;
            height: 600px;
            border: 2px solid #ddd;
            border-radius: 10px;
            margin-top: 15px;
        }
        .download-btn {
            background: linear-gradient(135deg, #11998e 0%, #38ef7d 100%);
            margin-top: 10px;
        }
        .loading {
            display: none;
            text-align: center;
            margin-top: 20px;
        }
        .spinner {
            border: 3px solid #f3f3f3;
            border-top: 3px solid #667eea;
            border-radius: 50%;
            width: 40px;
            height: 40px;
            animation: spin 1s linear infinite;
            margin: 0 auto;
        }
        @keyframes spin {
            0% { transform: rotate(0deg); }
            100% { transform: rotate(360deg); }
        }
    </style>
</head>
<body>
    <div class="container">
        <a href="/" class="back-btn">← Back to Home</a>
        <h1>📄 Merge PDFs</h1>
        <p class="subtitle">Insert all pages from one PDF into another at a specific position</p>
        
        <form id="uploadForm" enctype="multipart/form-data">
            <div class="upload-section">
                <label for="pdf1">First PDF (Base document) *</label>
                <div class="file-input-wrapper">
                    <input type="file" id="pdf1" name="pdf1" accept=".pdf" required>
                    <label for="pdf1" class="file-input-label">
                        Choose PDF file
                    </label>
                </div>
                <div id="file1Name" class="file-name"></div>
                <div id="pdf1Info" class="pdf-info"></div>
                <div id="pdf1Preview" class="pdf-preview"></div>
            </div>

            <div class="upload-section">
                <label for="pdf2">Second PDF (To be inserted) *</label>
                <div class="file-input-wrapper">
                    <input type="file" id="pdf2" name="pdf2" accept=".pdf" required>
                    <label for="pdf2" class="file-input-label">
                        Choose PDF file
                    </label>
                </div>
                <div id="file2Name" class="file-name"></div>
                <div id="pdf2Info" class="pdf-info"></div>
                <div id="pdf2Preview" class="pdf-preview"></div>
            </div>

            <div class="upload-section">
                <label for="pageNumber">Insert at page number *</label>
                <input type="number" id="pageNumber" name="pageNumber" min="1" required placeholder="e.g., 3">
                <div class="hint">Pages from PDF 2 will be inserted after this page in PDF 1</div>
            </div>

            <button type="submit">Merge PDFs</button>
        </form>

        <div class="loading" id="loading">
            <div class="spinner"></div>
            <p style="margin-top: 10px; color: #666;">Processing PDFs...</p>
        </div>

        <div id="result"></div>
        ` + getFooterHTML() + `
    </div>

    <script>
        var pdf1Input = document.getElementById('pdf1');
        var pdf2Input = document.getElementById('pdf2');
        var file1Name = document.getElementById('file1Name');
        var file2Name = document.getElementById('file2Name');
        var pdf1InfoDiv = document.getElementById('pdf1Info');
        var pdf2InfoDiv = document.getElementById('pdf2Info');
        var pdf1PreviewDiv = document.getElementById('pdf1Preview');
        var pdf2PreviewDiv = document.getElementById('pdf2Preview');

        pdf1Input.addEventListener('change', function(e) {
            if (e.target.files.length > 0) {
                var file = e.target.files[0];
                file1Name.textContent = 'Selected: ' + file.name;
                loadPDFInfo(file, pdf1InfoDiv, pdf1PreviewDiv, 1);
            }
        });

        pdf2Input.addEventListener('change', function(e) {
            if (e.target.files.length > 0) {
                var file = e.target.files[0];
                file2Name.textContent = 'Selected: ' + file.name;
                loadPDFInfo(file, pdf2InfoDiv, pdf2PreviewDiv, 2);
            }
        });

        function loadPDFInfo(file, infoDiv, previewDiv, pdfNumber) {
            var formData = new FormData();
            formData.append('pdf', file);

            infoDiv.innerHTML = 'Loading info...';
            infoDiv.classList.add('visible');
            previewDiv.classList.remove('visible');

            fetch('/pdfinfo', {
                method: 'POST',
                body: formData
            }).then(function(response) {
                return response.json();
            }).then(function(data) {
                if (data.error) {
                    infoDiv.innerHTML = 'Error: ' + data.error;
                } else {
                    infoDiv.innerHTML = '<strong>Pages:</strong> ' + data.pageCount + 
                                      ' | <strong>Size:</strong> ' + formatFileSize(file.size);
                    
                    var url = URL.createObjectURL(file);
                    previewDiv.innerHTML = '<iframe src="' + url + '#view=FitH"></iframe>';
                    previewDiv.classList.add('visible');
                }
            }).catch(function(error) {
                infoDiv.innerHTML = 'Error loading PDF info';
            });
        }

        function formatFileSize(bytes) {
            if (bytes < 1024) return bytes + ' B';
            if (bytes < 1024 * 1024) return (bytes / 1024).toFixed(1) + ' KB';
            return (bytes / (1024 * 1024)).toFixed(1) + ' MB';
        }

        document.getElementById('uploadForm').addEventListener('submit', function(e) {
            e.preventDefault();
            
            var formData = new FormData(e.target);
            var resultDiv = document.getElementById('result');
            var loadingDiv = document.getElementById('loading');
            var submitBtn = e.target.querySelector('button[type="submit"]');
            
            resultDiv.innerHTML = '';
            loadingDiv.style.display = 'block';
            submitBtn.disabled = true;

            fetch('/merge-pdfs', {
                method: 'POST',
                body: formData
            }).then(function(response) {
                loadingDiv.style.display = 'none';
                submitBtn.disabled = false;

                if (response.ok) {
                    return response.blob().then(function(blob) {
                        var url = URL.createObjectURL(blob);
                        
                        resultDiv.innerHTML = 
                            '<div class="success">PDFs merged successfully!</div>' +
                            '<iframe id="pdfViewer" src="' + url + '"></iframe>' +
                            '<button class="download-btn" onclick="downloadPDF(\'' + url + '\')">Download Merged PDF</button>';
                    });
                } else {
                    return response.text().then(function(error) {
                        resultDiv.innerHTML = '<div class="error">Error: ' + error + '</div>';
                    });
                }
            }).catch(function(error) {
                loadingDiv.style.display = 'none';
                submitBtn.disabled = false;
                resultDiv.innerHTML = '<div class="error">Error: ' + error.message + '</div>';
            });
        });

        function downloadPDF(url) {
            var a = document.createElement('a');
            a.href = url;
            a.download = 'merged_' + Date.now() + '.pdf';
            document.body.appendChild(a);
            a.click();
            document.body.removeChild(a);
        }
    </script>
</body>
</html>`
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html))
}

func handlePDFInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"error": "File too large"}`))
		return
	}

	pdfFile, _, err := r.FormFile("pdf")
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"error": "Error retrieving PDF"}`))
		return
	}
	defer pdfFile.Close()

	tempFile, err := os.CreateTemp("", "pdfinfo-*.pdf")
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"error": "Error creating temp file"}`))
		return
	}
	tempPath := tempFile.Name()
	defer os.Remove(tempPath)

	if _, err := io.Copy(tempFile, pdfFile); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"error": "Error saving PDF"}`))
		return
	}
	tempFile.Close()

	pageCount, err := getPageCount(tempPath)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"error": "Error reading PDF"}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := fmt.Sprintf(`{"pageCount": %d}`, pageCount)
	w.Write([]byte(response))
}

func handleRemovePages(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		http.Error(w, "File too large. Maximum size is 50MB", http.StatusBadRequest)
		return
	}

	pdfFile, pdfHeader, err := r.FormFile("pdf")
	if err != nil {
		http.Error(w, "Error retrieving PDF", http.StatusBadRequest)
		return
	}
	defer pdfFile.Close()

	pagesToRemoveStr := r.FormValue("pagesToRemove")
	if pagesToRemoveStr == "" {
		http.Error(w, "No pages specified for removal", http.StatusBadRequest)
		return
	}

	if filepath.Ext(pdfHeader.Filename) != ".pdf" {
		http.Error(w, "Only PDF files are allowed", http.StatusBadRequest)
		return
	}

	tempDir, err := os.MkdirTemp("", "pdfremove-*")
	if err != nil {
		http.Error(w, "Error creating temporary directory", http.StatusInternalServerError)
		return
	}
	defer os.RemoveAll(tempDir)

	inputPath := filepath.Join(tempDir, "input.pdf")
	outputPath := filepath.Join(tempDir, "output.pdf")

	if err := saveFile(pdfFile, inputPath); err != nil {
		http.Error(w, "Error saving PDF", http.StatusInternalServerError)
		return
	}

	pageCount, err := getPageCount(inputPath)
	if err != nil {
		http.Error(w, "Error reading PDF: "+err.Error(), http.StatusBadRequest)
		return
	}

	pagesToRemove := parsePageNumbers(pagesToRemoveStr)
	if len(pagesToRemove) == 0 {
		http.Error(w, "Invalid page numbers", http.StatusBadRequest)
		return
	}

	if len(pagesToRemove) >= pageCount {
		http.Error(w, "Cannot remove all pages. At least one page must remain.", http.StatusBadRequest)
		return
	}

	pagesToKeep := getPagesToKeep(pageCount, pagesToRemove)
	if len(pagesToKeep) == 0 {
		http.Error(w, "No pages would remain after removal", http.StatusBadRequest)
		return
	}

	pageRanges := createPageRanges(pagesToKeep)
	if err := api.TrimFile(inputPath, outputPath, pageRanges, nil); err != nil {
		http.Error(w, "Error removing pages: "+err.Error(), http.StatusInternalServerError)
		return
	}

	modifiedPDF, err := os.ReadFile(outputPath)
	if err != nil {
		http.Error(w, "Error reading modified PDF", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", fmt.Sprintf("inline; filename=modified_%d.pdf", time.Now().Unix()))
	w.Write(modifiedPDF)
}

func handleMerge(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		http.Error(w, "File too large. Maximum size is 50MB", http.StatusBadRequest)
		return
	}

	pdf1File, pdf1Header, err := r.FormFile("pdf1")
	if err != nil {
		http.Error(w, "Error retrieving first PDF", http.StatusBadRequest)
		return
	}
	defer pdf1File.Close()

	pdf2File, pdf2Header, err := r.FormFile("pdf2")
	if err != nil {
		http.Error(w, "Error retrieving second PDF", http.StatusBadRequest)
		return
	}
	defer pdf2File.Close()

	pageNumStr := r.FormValue("pageNumber")
	pageNum, err := strconv.Atoi(pageNumStr)
	if err != nil || pageNum < 1 {
		http.Error(w, "Invalid page number", http.StatusBadRequest)
		return
	}

	if filepath.Ext(pdf1Header.Filename) != ".pdf" || filepath.Ext(pdf2Header.Filename) != ".pdf" {
		http.Error(w, "Only PDF files are allowed", http.StatusBadRequest)
		return
	}

	tempDir, err := os.MkdirTemp("", "pdfmerge-*")
	if err != nil {
		http.Error(w, "Error creating temporary directory", http.StatusInternalServerError)
		return
	}
	defer os.RemoveAll(tempDir)

	pdf1Path := filepath.Join(tempDir, "pdf1.pdf")
	pdf2Path := filepath.Join(tempDir, "pdf2.pdf")
	outputPath := filepath.Join(tempDir, "merged.pdf")

	if err := saveFile(pdf1File, pdf1Path); err != nil {
		http.Error(w, "Error saving first PDF", http.StatusInternalServerError)
		return
	}

	if err := saveFile(pdf2File, pdf2Path); err != nil {
		http.Error(w, "Error saving second PDF", http.StatusInternalServerError)
		return
	}

	pdf1PageCount, err := getPageCount(pdf1Path)
	if err != nil {
		http.Error(w, "Error reading first PDF: "+err.Error(), http.StatusBadRequest)
		return
	}

	pdf2PageCount, err := getPageCount(pdf2Path)
	if err != nil {
		http.Error(w, "Error reading second PDF: "+err.Error(), http.StatusBadRequest)
		return
	}

	if pdf1PageCount < 2 {
		http.Error(w, fmt.Sprintf("First PDF must have multiple pages (found %d page)", pdf1PageCount), http.StatusBadRequest)
		return
	}

	if pdf2PageCount < 2 {
		http.Error(w, fmt.Sprintf("Second PDF must have multiple pages (found %d page)", pdf2PageCount), http.StatusBadRequest)
		return
	}

	if pageNum > pdf1PageCount {
		http.Error(w, fmt.Sprintf("Page number %d exceeds first PDF page count (%d pages)", pageNum, pdf1PageCount), http.StatusBadRequest)
		return
	}

	if err := mergePDFs(pdf1Path, pdf2Path, outputPath, pageNum); err != nil {
		http.Error(w, "Error merging PDFs: "+err.Error(), http.StatusInternalServerError)
		return
	}

	mergedPDF, err := os.ReadFile(outputPath)
	if err != nil {
		http.Error(w, "Error reading merged PDF", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", fmt.Sprintf("inline; filename=merged_%d.pdf", time.Now().Unix()))
	w.Write(mergedPDF)
}

func saveFile(src io.Reader, dst string) error {
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}

func getPageCount(pdfPath string) (int, error) {
	ctx, err := api.ReadContextFile(pdfPath)
	if err != nil {
		return 0, err
	}
	return ctx.PageCount, nil
}

func mergePDFs(pdf1Path, pdf2Path, outputPath string, insertAfterPage int) error {
	ctx1, err := api.ReadContextFile(pdf1Path)
	if err != nil {
		return fmt.Errorf("failed to read first PDF: %w", err)
	}

	ctx2, err := api.ReadContextFile(pdf2Path)
	if err != nil {
		return fmt.Errorf("failed to read second PDF: %w", err)
	}

	tempDir := filepath.Dir(pdf1Path)
	part1Path := filepath.Join(tempDir, "part1.pdf")
	part2Path := filepath.Join(tempDir, "part2.pdf")
	
	if insertAfterPage > 0 {
		pageRange := fmt.Sprintf("1-%d", insertAfterPage)
		if err := api.TrimFile(pdf1Path, part1Path, []string{pageRange}, nil); err != nil {
			return fmt.Errorf("failed to create first part: %w", err)
		}
	}
	
	if insertAfterPage < ctx1.PageCount {
		pageRange := fmt.Sprintf("%d-%d", insertAfterPage+1, ctx1.PageCount)
		if err := api.TrimFile(pdf1Path, part2Path, []string{pageRange}, nil); err != nil {
			return fmt.Errorf("failed to create second part: %w", err)
		}
	}
	
	var filesToMerge []string
	if insertAfterPage > 0 {
		filesToMerge = append(filesToMerge, part1Path)
	}
	filesToMerge = append(filesToMerge, pdf2Path)
	if insertAfterPage < ctx1.PageCount {
		filesToMerge = append(filesToMerge, part2Path)
	}
	
	if err := api.MergeCreateFile(filesToMerge, outputPath, false, nil); err != nil {
		return fmt.Errorf("failed to merge PDFs: %w", err)
	}

	if err := api.ValidateFile(outputPath, model.NewDefaultConfiguration()); err != nil {
		return fmt.Errorf("merged PDF validation failed: %w", err)
	}

	_ = ctx1
	_ = ctx2

	return nil
}

func parsePageNumbers(pageStr string) []int {
	var pages []int
	parts := strings.Split(pageStr, ",")
	
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		
		pageNum, err := strconv.Atoi(part)
		if err == nil && pageNum > 0 {
			pages = append(pages, pageNum)
		}
	}
	
	return pages
}

func getPagesToKeep(totalPages int, pagesToRemove []int) []int {
	removeMap := make(map[int]bool)
	for _, page := range pagesToRemove {
		removeMap[page] = true
	}
	
	var pagesToKeep []int
	for i := 1; i <= totalPages; i++ {
		if !removeMap[i] {
			pagesToKeep = append(pagesToKeep, i)
		}
	}
	
	return pagesToKeep
}

func createPageRanges(pages []int) []string {
	if len(pages) == 0 {
		return nil
	}
	
	var ranges []string
	start := pages[0]
	end := pages[0]
	
	for i := 1; i < len(pages); i++ {
		if pages[i] == end+1 {
			end = pages[i]
		} else {
			if start == end {
				ranges = append(ranges, fmt.Sprintf("%d", start))
			} else {
				ranges = append(ranges, fmt.Sprintf("%d-%d", start, end))
			}
			start = pages[i]
			end = pages[i]
		}
	}
	
	if start == end {
		ranges = append(ranges, fmt.Sprintf("%d", start))
	} else {
		ranges = append(ranges, fmt.Sprintf("%d-%d", start, end))
	}
	
	return ranges
}