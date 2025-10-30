# 🔧 PDF Tools

Free & Open Source web application for merging and editing PDF files. Available in two versions: **Server-based (Go)** and **Client-side (HTML)**.

![License](https://img.shields.io/badge/license-MIT-blue.svg)
![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go)
![Built with Claude.ai](https://img.shields.io/badge/Built%20with-Claude.ai-8A2BE2)

## ✨ Features

- **Remove Pages**: Delete specific pages from PDFs with visual page selector
- **Merge PDFs**: Combine multiple PDFs into one document
- **Optimize PDF**: Reduce PDF file size by optimizing and compressing
- **PDF Preview**: View uploaded PDFs and see page counts before processing
- **Privacy Focused**: Files are processed temporarily (server) or locally (client)
- **No Registration**: Use freely without any account or tracking
- **Mobile Friendly**: Works on desktop and mobile devices

## 📦 Two Versions Available

### 🖥️ Server Version (Go + pdfcpu)
**File:** `main.go`
- Full-featured PDF processing with advanced capabilities
- Requires Go runtime and server deployment
- Better optimization and compression
- Ideal for production deployments on Railway, Render, etc.

### 🌐 Client-Side Version (Pure HTML/JavaScript)
**File:** `pdf-tools-standalone.html`
- Single HTML file - no server needed!
- Runs entirely in your browser using pdf-lib
- 100% privacy - files never leave your device
- Works offline after initial load
- Perfect for GitHub Pages or local use
- Just open the HTML file in any modern browser

**Choose based on your needs:**
- Need advanced features and don't mind server setup? → Use **Server Version**
- Want maximum privacy and zero setup? → Use **Client-Side Version**

## 🚀 Quick Start

### Option 1: Client-Side Version (Easiest!)

**No installation needed!**

1. Download or open `pdf-tools-standalone.html`
2. Double-click the file or open it in any modern web browser
3. Start processing PDFs immediately!

**Or use it online:**
- GitHub Pages: `https://rajmahavir.github.io/PDF-Tools/pdf-tools-standalone.html`
- Or host it anywhere - it's just one HTML file!

### Option 2: Server Version (Advanced)

**Prerequisites:**
- Go 1.21 or higher
- Git (optional)

**Installation:**

1. Clone or download this repository
```bash
git clone https://github.com/rajmahavir/PDF-Tools.git
cd PDF-Tools
```

2. Install dependencies
```bash
go mod download
```

3. Run the application
```bash
go run main.go
```

4. Open your browser
```
http://localhost:8080
```

**Access from Mobile Devices:**

1. Find your computer's IP address:
   - Windows: Run `ipconfig` in Command Prompt
   - Mac/Linux: Run `ifconfig` or `ip addr`

2. On your mobile device (same WiFi):
   - Open browser and go to `http://[YOUR-IP]:8080`
   - Example: `http://192.168.1.100:8080`

## 🤖 AI-Assisted Development

This project was developed with significant assistance from **Claude.ai** (Anthropic's AI assistant). The collaboration involved:

- Architecture design and code generation
- User interface development
- PDF processing implementation
- Error handling and optimization

**Model Used**: Claude Sonnet 4.5  
**Platform**: [Claude.ai](https://claude.ai)

## 🛠️ Technology Stack

### Server Version
- **Backend**: Go (Golang) 1.21+
- **PDF Processing**: [pdfcpu](https://github.com/pdfcpu/pdfcpu) v0.11.1
- **Web Server**: Go net/http (standard library)
- **Frontend**: HTML, CSS, JavaScript (Vanilla)

### Client-Side Version
- **PDF Processing**: [pdf-lib](https://pdf-lib.js.org/) v1.17.1 (JavaScript)
- **Runtime**: Browser-based (no backend required)
- **Frontend**: HTML, CSS, JavaScript (Vanilla)

### Development
- **AI Assistant**: Claude.ai by Anthropic (Claude Sonnet 4.5)

## 📄 License

This project is licensed under the MIT License - see the [LICENSE.txt](LICENSE.txt) file for details.

### Third-Party Licenses

**Server Version:**
- **pdfcpu**: Apache License 2.0 - PDF processing library for Go
- **Go**: BSD 3-Clause License - Programming language

**Client-Side Version:**
- **pdf-lib**: MIT License - JavaScript PDF library

See [NOTICE.txt](NOTICE.txt) for full third-party attributions.

## 🔒 Privacy & Security

### Server Version
- All PDF processing happens on the server temporarily
- Files are automatically deleted immediately after processing
- No data storage, tracking, or analytics
- No cookies or user registration required

### Client-Side Version
- **Maximum Privacy**: All processing happens in your browser
- Files **never** leave your device or get uploaded anywhere
- Works completely offline after initial page load
- No server, no tracking, no data collection
- Open source - verify the code yourself

**Both versions:**
- Open source code - audit it yourself
- No user registration or accounts required
- No cookies or analytics

## 🙏 Acknowledgments

- **Anthropic** - For creating Claude.ai, the AI assistant used to develop this project
- **Horst Rutter** - For developing and maintaining [pdfcpu](https://github.com/pdfcpu/pdfcpu)
- **Andrew Dillon** - For creating [pdf-lib](https://pdf-lib.js.org/)
- **Go Team** - For the excellent Go programming language
- **Open Source Community** - For making projects like this possible

## 📝 Disclaimer

This tool is provided "as is" without warranty of any kind. Use at your own risk.

## 🤝 Contributing

Contributions are welcome! Feel free to:

- Report bugs via GitHub Issues
- Submit feature requests
- Create pull requests
- Share the project

## 📧 Contact

Created by Raj

GitHub: rajmahavir

---

**Made with ❤️ using Claude.ai**

## 🌐 Live Demos

### Server Version
Visit: https://pdf-tools-production-bb6b.up.railway.app/

### Client-Side Version
Try it now (no backend needed!):
- **GitHub Pages**: `https://rajmahavir.github.io/PDF-Tools/pdf-tools-standalone.html`
- Or download and open locally: [`pdf-tools-standalone.html`](pdf-tools-standalone.html)

## 📸 Screenshots

### Home Page
Choose between removing pages, merging PDFs, or optimizing file size with a clean, intuitive interface.

### Remove Pages
Visual page selector makes it easy to choose which pages to delete from your PDF.

### Merge PDFs
Combine multiple PDFs into a single document seamlessly.

### Optimize PDF
Reduce file size while maintaining quality.

## 🐛 Known Issues

**Server Version:**
- Large PDFs (>50MB) may take longer to process
- Some encrypted PDFs may not be supported

**Client-Side Version:**
- Browser-based optimization has limitations compared to server-side
- Very large PDFs may cause browser memory issues
- Some advanced PDF features may not be supported

## 🔮 Future Features

- [ ] Rotate pages
- [ ] Split PDF into multiple files
- [ ] Add watermarks
- [ ] Extract specific page ranges
- [ ] PDF password protection/encryption
- [ ] Insert pages at specific positions (merge enhancement)
- [ ] Batch processing multiple files

## ⚙️ Configuration

### Server Version

**Default settings:**
- Port: 8080 (or from PORT environment variable)
- Max file size: 50MB
- Supported formats: PDF only

**Environment Variables:**
```bash
PORT=8080  # Port to run the server on (automatically set by Railway, Render, etc.)
```

The application automatically reads the `PORT` environment variable for cloud deployments.

### Client-Side Version

No configuration needed! Just open the HTML file and it works.

## 🧪 Testing

### Server Version
```bash
go run main.go
```
Open `http://localhost:8080` and upload test PDFs.

### Client-Side Version
Simply open `pdf-tools-standalone.html` in your browser and upload test PDFs.

## 📦 Building for Production

### Server Version

Build executable:
```bash
go build -o pdf-tools
```

Run the executable:
```bash
# Windows
pdf-tools.exe

# Linux/Mac
./pdf-tools
```

### Client-Side Version

No build needed! The HTML file is production-ready as-is.

## 🌍 Deployment

### Server Version (Go)

**Railway.app (Recommended):**
1. Connect your GitHub repository
2. Railway auto-detects Go and deploys
3. Environment variable `PORT` is automatically set

**Other platforms:**
- Render.com
- Fly.io
- Oracle Cloud
- Google Cloud Run
- Heroku

All major platforms support Go and will automatically use the `PORT` environment variable.

### Client-Side Version (HTML)

**GitHub Pages (Easiest):**
1. Enable GitHub Pages in repository settings
2. Select branch and root folder
3. Access at: `https://rajmahavir.github.io/PDF-Tools/pdf-tools-standalone.html`

**Other platforms:**
- Netlify (drag & drop the HTML file)
- Vercel
- Cloudflare Pages
- Any static web hosting
- Or just share the file directly!

---

**Star ⭐ this repo if you find it useful!**