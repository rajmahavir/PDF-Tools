# 🔧 PDF Tools

Free & Open Source web application for merging and editing PDF files.

![License](https://img.shields.io/badge/license-MIT-blue.svg)
![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go)
![Built with Claude.ai](https://img.shields.io/badge/Built%20with-Claude.ai-8A2BE2)

## ✨ Features

- **Remove Pages**: Delete specific pages from PDFs with visual page selector
- **Merge PDFs**: Insert all pages from one PDF into another at any position
- **PDF Preview**: View uploaded PDFs and see page counts before processing
- **Privacy Focused**: Files are processed temporarily and deleted immediately
- **No Registration**: Use freely without any account or tracking
- **Mobile Friendly**: Works on desktop and mobile devices

## 🚀 Quick Start

### Prerequisites

- Go 1.21 or higher
- Git (optional)

### Installation

1. Clone or download this repository
```bash
git clone https://github.com/rajmahavir/PDF-Tools.git
cd PDF-Tools
```

2. Install dependencies
```bash
go mod init pdf-tools
go get github.com/pdfcpu/pdfcpu/pkg/api
go get github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model
```

3. Run the application
```bash
go run main.go
```

4. Open your browser
```
http://localhost:8080
```

### Access from Mobile Devices

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

- **Backend**: Go (Golang)
- **PDF Processing**: [pdfcpu](https://github.com/pdfcpu/pdfcpu)
- **AI Assistant**: Claude.ai by Anthropic
- **Frontend**: HTML, CSS, JavaScript (Vanilla)

## 📄 License

This project is licensed under the MIT License - see the [LICENSE.txt](LICENSE.txt) file for details.

### Third-Party Licenses

- **pdfcpu**: Apache License 2.0
- **Go**: BSD 3-Clause License

See [NOTICE.txt](NOTICE.txt) for full third-party attributions.

## 🔒 Privacy & Security

- All PDF processing happens on the server temporarily
- Files are automatically deleted after processing
- No data storage, tracking, or analytics
- No cookies or user registration required
- Open source - verify the code yourself

## 🙏 Acknowledgments

- **Anthropic** - For creating Claude.ai
- **Horst Rutter** - For developing pdfcpu
- **Go Team** - For the Go programming language
- **Open Source Community** - For making this possible

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

## 📸 Screenshots

### Home Page
Choose between removing pages or merging PDFs with a clean, intuitive interface.

### Remove Pages
Visual page selector makes it easy to choose which pages to delete.

### Merge PDFs
Insert one PDF into another at any position with live preview.

## 🌐 Live Demo

Visit: https://pdf-tools-production-bb6b.up.railway.app/

## 🐛 Known Issues

- Large PDFs (>50MB) may take longer to process
- Some encrypted PDFs may not be supported

## 🔮 Future Features

- [ ] Rotate pages
- [ ] Split PDF into multiple files
- [ ] Compress PDF files
- [ ] Add watermarks
- [ ] Merge multiple PDFs at once
- [ ] Extract specific page ranges

## ⚙️ Configuration

Default settings:
- Port: 8080
- Max file size: 50MB
- Supported formats: PDF only

To change port, edit `main.go`:
```go
port := "8080"  // Change to your preferred port
```

## 🧪 Testing

Run locally and test with sample PDFs:
```bash
go run main.go
```

Open `http://localhost:8080` and upload test PDFs.

## 📦 Building for Production

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

## 🌍 Deployment

See deployment guides for:
- Railway.app
- Render.com
- Fly.io
- Oracle Cloud
- Google Cloud Run

Detailed instructions available in the [Deployment Guide](DEPLOYMENT.md) (if you create one).

---

**Star ⭐ this repo if you find it useful!**