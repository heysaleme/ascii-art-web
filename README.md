# ASCII Art Web 🖥️🎨

![Главная страница](https://github.com/heysaleme/ascii-art-web/raw/main/screenshots/main_page.png)

## Description
**ASCII Art Web** is a Go web application that provides a browser-based graphical interface for generating ASCII art. It allows users to input text, choose a banner style, and instantly generate ASCII art directly from their web browser.

The application serves an interactive web page where users can:
- Type any text they want to convert to ASCII art
- Select from available banner styles: `standard`, `shadow`, or `thinkertoy`
- View the generated ASCII art in real-time
- Copy or download the results

---

## Features ✨

### 🎨 Web Interface
- Clean, responsive web interface with intuitive controls
- Real-time ASCII art generation
- Horizontal scrolling for wide ASCII art outputs
- Dark/light theme friendly design

### 🔒 Robust Error Handling
- **200 OK** – Successful ASCII art generation
- **400 Bad Request** – Invalid input or unsupported characters
- **404 Not Found** – Banner file not found or template missing
- **405 Method Not Allowed** – Invalid HTTP method used
- **500 Internal Server Error** – Server-side errors or corrupted banner files

### 🛡️ Banner Integrity Protection
- **SHA-256 hash verification** of banner files
- **Strict validation** of ASCII art character structure
- **Automatic detection** of modified or corrupted banner files
- **Character-by-character integrity checks**

### 🌐 Multi-Banner Support
- **Standard** – Classic ASCII art style
- **Shadow** – Text with shadow effect
- **Thinkertoy** – Bold, blocky style

---

## Prerequisites 📋

### Required Software
- **Go 1.21** or higher
- Modern web browser (Chrome, Firefox, Safari, Edge)

### File Requirements
The following banner files must exist in the `banners/` directory:
- `standard.txt`
- `shadow.txt`
- `thinkertoy.txt`

---

## Installation & Setup 🚀

### 1. Clone the Repository
```bash
git clone https://github.com/yourusername/ascii-art-web.git
cd ascii-art-web
```

### 2. Generate Banner Hashes (First Time Only)
```bash
# Edit main.go to call GenerateBannerHashes()
# Then run:
go run main.go 
# Copy the generated hashes into ascii/hashes.go
```

### 3. Start the Web Server
```bash
go run main.go web.go
```

### 4. Access the Application
Open your browser and navigate to:
```
http://localhost:8080
```

---

## API Endpoints 🔌

### `GET /`
- **Description**: Serves the main HTML interface
- **Responses**:
  - `200 OK` – HTML page served successfully
  - `404 Not Found` – HTML template file missing
  - `500 Internal Server Error` – Template parsing error

### `GET /generate`
- **Description**: Generates ASCII art from provided text
- **Query Parameters**:
  - `text` (required) – Text to convert to ASCII art
  - `banner` (optional, default: `standard`) – Banner style to use
- **Responses**:
  - `200 OK` – ASCII art returned as plain text
  - `400 Bad Request` – Missing text parameter or unsupported characters
  - `404 Not Found` – Banner file not found
  - `405 Method Not Allowed` – If called with POST instead of GET
  - `500 Internal Server Error` – Banner file corrupted or server error

---

## Character Support 🔤

### ✅ Supported Characters
- All printable ASCII characters (codes 32-126)
- Line breaks (`\n`)
- Space characters

### ❌ Unsupported Characters
- Cyrillic letters (Russian, Ukrainian, etc.)
- Asian characters (Chinese, Japanese, Korean)
- Arabic script
- Emoji and special symbols
- Characters with diacritical marks (é, ñ, ü, etc.)

---

## Error Scenarios & Handling ⚠️

### Common Error Responses

#### 400 Bad Request
- Empty text input
- Non-ASCII characters detected
- Malformed input parameters

#### 404 Not Found
- Requested banner file doesn't exist
- HTML template file missing

#### 405 Method Not Allowed
- Using POST method on `/` endpoint
- Using GET method where POST is expected

#### 500 Internal Server Error
- Banner file has been modified or corrupted
- Banner file has incorrect structure
- Internal server processing errors
- Banner hash verification failed

### Banner Integrity Errors
The system performs multiple checks on banner files:
1. **SHA-256 Hash Verification** – Detects any file modifications
2. **Structure Validation** – Ensures 9 lines per character
3. **Character Consistency** – Verifies internal character formatting
4. **Content Validation** – Ensures only valid ASCII art characters are used

---

## Technical Details 🔧

### ASCII Art Generation Algorithm
1. **Banner Loading** – Font file is loaded and validated
2. **Text Processing** – Input is split by line breaks
3. **Character Mapping** – Each character maps to its 9-line representation
4. **Line Assembly** – Characters are combined horizontally
5. **Output Formatting** – Proper spacing and line breaks are added

### Banner File Requirements
- Each character must occupy exactly 9 lines
- 95 total characters (ASCII 32-126)
- Space character (ASCII 32) must be completely empty
- All lines within a character must have consistent width
- Only valid ASCII art drawing characters are allowed

---

## Development 🛠️

### Adding New Banners
1. Place new banner file in `banners/` directory
2. Add banner name to validation logic
3. Generate and add SHA-256 hash to `hashes.go`
4. Update banner selection in `index.html`

### Testing Banner Integrity
```bash
# Modify any banner file (add/remove space)
# Restart server and attempt generation
# Should receive 500 error with specific details
```

### Running Tests
```bash
# Test with valid input
curl "http://localhost:8080/generate?text=Hello&banner=standard"

# Test error cases
curl "http://localhost:8080/generate?text="
curl "http://localhost:8080/generate?text=Привет"
curl "http://localhost:8080/generate?text=test&banner=invalid"
```

---

## Security Considerations 🔐

### File Integrity
- Banner files are verified with SHA-256 hashes
- Any modification to banner files triggers 500 error
- Prevents malicious or accidental file corruption

### Input Validation
- Strict character range checking (32-126 ASCII)
- Parameter sanitization
- Size limits on input text

### Error Information
- Detailed error messages for developers
- Generic error messages for end users
- No sensitive information leakage in errors

---


### Code Standards
- Follow Go standard formatting
- Include appropriate error handling
- Add comments for complex logic
- Update documentation as needed

---

**Happy ASCII Art Creating! 🎨✨**