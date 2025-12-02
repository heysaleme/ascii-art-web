
# ASCII Art Web 🖥️🎨

## Description
**ASCII Art Web** is a Go web application that provides a browser-based graphical interface for the `ascii-art` project.  
It allows users to input text, choose a banner (font style), and generate ASCII art directly from their browser.

The app serves a simple web page with an input form, where users can:
- Type any text they want to convert.
- Select one of the available banners: `standard`, `shadow`, or `thinkertoy`.
- Click a button to generate the result using the Go backend.

---

## Usage

### 🧩 Prerequisites
Make sure you have **Go** installed:
```bash
go version
````

### 🚀 How to Run

Clone the repository and start the web server:

```bash
git clone https://github.com/yourusername/ascii-art-web.git
cd ascii-art-web
go run .
```

After starting, open your browser and go to:

```
http://localhost:8080
```

You’ll see the ASCII Art generator page.

---

## Implementation Details

### 📁 Project Structure

```
ascii-art-web/
├── ascii/
│   ├── ascii.go           # ASCII art generation logic
├── banners/
│   ├── standard.txt
│   ├── shadow.txt
│   ├── thinkertoy.txt
├── static/
│   ├── index.html         # Main web page
├── web.go                 # HTTP server and route handlers
└── README.md
```

### ⚙️ Endpoints

#### `GET /`

* **Description:** Serves the main HTML page.
* **Response:** 200 OK if successful.
* **Errors:**

  * 404 Not Found – if the HTML template is missing.
  * 500 Internal Server Error – if an unexpected error occurs.

#### `POST /ascii-art`

* **Description:** Accepts user input (`text` and `banner`), generates ASCII art, and returns it.
* **Request Body:**

  * `text` — The input string to be converted.
  * `banner` — One of the available fonts (`standard`, `shadow`, `thinkertoy`).
* **Response:**

  * 200 OK – Returns ASCII art as plain text.
  * 400 Bad Request – Invalid or unsupported input.
  * 404 Not Found – Banner not found.
  * 500 Internal Server Error – For unhandled server errors.

---

## 🧠 Algorithm Overview

1. The banner file (font) is loaded from the `/banners` directory.
2. The input text is split into lines (`\n`).
3. Each printable ASCII character (from 32 to 126) is matched to its 8-line representation in the banner.
4. Lines are combined into a full ASCII art output using string builders.
5. The result is sent back as plain text to the web page, which displays it inside a `<pre>` block with horizontal scrolling.

---

## 💡 Features

* Web GUI with a clean and simple interface.
* Support for three different banners: `standard`, `shadow`, and `thinkertoy`.
* Graceful error handling with proper HTTP status codes:

  * `200 OK`
  * `400 Bad Request`
  * `404 Not Found`
  * `500 Internal Server Error`
* Horizontal scroll support for long ASCII art lines.
* Built entirely in **Go**, using `html/template` and standard libraries.

---

## 🧭 Example

### Request:

```
POST /ascii-art
text=Hello
banner=shadow
```

### Response:

```
_|    _|          _| _|
_|    _|   _|_|   _| _|    _|_|
_|_|_|_| _|_|_|_| _| _|  _|    _|
_|    _| _|       _| _|  _|    _|
_|    _|   _|_|_| _| _|    _|_|
```

# ascii-art-web
