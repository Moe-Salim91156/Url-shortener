# 🔗 URL Shortener

A lightweight, high-performance URL shortener built with **Go**. This service generates short codes for long URLs and redirects users seamlessly.

## ✨ Features

- **URL Shortening**: Convert long URLs into compact 4-character codes
- **Fast Redirects**: Instant redirection from short codes to original URLs
- **In-Memory Storage**: Lightning-fast lookups using Go maps
- **RESTful API**: Clean HTTP endpoints for shortening and resolving URLs
- **Simple Architecture**: Organized codebase following Go best practices

## 🚀 Quick Start

### Prerequisites

- Go 1.x or higher installed on your system

### Installation

1. Clone the repository:
```bash
git clone https://github.com/Moe-Salim91156/Url-shortener.git
cd Url-shortener
```

2. Run the server:
```bash
go run cmd/main.go
```

The server will start on `http://localhost:8000`

## 📖 Usage

### Step 1: Shorten a URL

Send a POST request to create a short code:

**Using curl:**
```bash
curl -X POST http://localhost:8000/shorten \
  -H "Content-Type: application/json" \
  -d '{"Url": "https://www.google.com"}'
```

**Using Postman or Insomnia:**
- Method: `POST`
- URL: `http://localhost:8000/shorten`
- Headers: `Content-Type: application/json`
- Body:
  ```json
  {
    "Url": "https://www.google.com"
  }
  ```

**Response:** A 4-character short code, for example:
```
abcd
```

### Step 2: Use the Short Code

**Option A - Browser (Recommended):**

Simply paste the short URL in your browser's address bar:
```
http://localhost:8000/abcd
```

You'll be automatically redirected to `https://www.google.com`

**Option B - Command Line:**
```bash
curl -L http://localhost:8000/abcd
```

The `-L` flag tells curl to follow the redirect.

### Complete Example

```bash
# 1. Shorten a URL and capture the code
SHORT_CODE=$(curl -s -X POST http://localhost:8000/shorten \
  -H "Content-Type: application/json" \
  -d '{"Url": "https://github.com"}')

echo "Your short code is: $SHORT_CODE"

# 2. Test the redirect
curl -L http://localhost:8000/$SHORT_CODE
```

## 🏗️ Project Structure

```
Url-shortener/
├── cmd/
│   └── main.go           # Application entry point & HTTP handlers
├── internal/
│   ├── models/
│   │   └── model.go      # UrlData struct definition
│   ├── shortener/
│   │   └── shorten.go    # Short code generation logic
│   └── store/
│       └── store.go      # In-memory storage operations
├── go.mod                # Go module definition
└── README.md
```

## 🔧 API Reference

### `POST /shorten`

Create a short code for a long URL.

**Request Body:**
```json
{
  "Url": "https://www.example.com"
}
```

**Response:** Plain text short code (e.g., `wxyz`)

**Status Codes:**
- `200 OK` - Successfully created short code
- `400 Bad Request` - Invalid JSON payload

---

### `GET /{shortCode}`

Redirect to the original URL associated with the short code.

**Example:**
```
GET http://localhost:8000/wxyz
```

**Response:** HTTP 301 redirect to the original URL

**Status Codes:**
- `301 Moved Permanently` - Successful redirect to original URL
- `404 Not Found` - Short code doesn't exist

**Note:** This is how URL shorteners work! The user:
1. Gets a short code from `/shorten`
2. Manually constructs the short URL: `http://localhost:8000/{code}`
3. Visits that URL in a browser or shares it
4. Gets automatically redirected to the original long URL

## 💡 How It Works

```
┌─────────┐      POST /shorten       ┌─────────┐
│         │  ───────────────────────> │         │
│  Client │  {"Url": "long-url"}     │  Server │
│         │  <─────────────────────── │         │
└─────────┘      "abcd"               └─────────┘
                                            │
     │                                      │ Stores in map:
     │                                      │ "abcd" → "long-url"
     │                                      ▼
     │
     │         GET /abcd              ┌─────────┐
     └──────────────────────────────> │  Server │
                                      │         │
       301 Redirect → "long-url"     │ Lookup  │
     <────────────────────────────── │  "abcd" │
                                      └─────────┘
```

1. **Shorten**: Client sends long URL → Server generates random 4-char code → Returns code
2. **Store**: Server saves mapping in memory (`shortCode → longUrl`)
3. **Redirect**: User visits `localhost:8000/{shortCode}` → Server looks up URL → Returns 301 redirect

## ⚠️ Current Limitations

- **In-Memory Storage**: All shortened URLs are lost when the server restarts
- **Manual URL Construction**: Users must manually add the short code to `localhost:8000/` 
- **No Domain**: Only works on localhost (not accessible from other machines)
- **No Collision Handling**: Duplicate short codes may theoretically occur
- **No Custom Codes**: Short codes are randomly generated
- **No Analytics**: No tracking of clicks or usage statistics
- **No URL Validation**: Accepts any string as a URL
- **No Expiration**: URLs stored indefinitely until restart

## 🎯 Testing Tips

**Quick Test Script:**
```bash
#!/bin/bash
echo "Starting URL shortener test..."

# Start the server in background (if not already running)
# go run cmd/main.go &

# Shorten a URL
echo -e "\n1. Shortening URL..."
CODE=$(curl -s -X POST http://localhost:8000/shorten \
  -H "Content-Type: application/json" \
  -d '{"Url": "https://github.com/Moe-Salim91156"}')

echo "   Short code: $CODE"
echo "   Short URL: http://localhost:8000/$CODE"

# Test redirect
echo -e "\n2. Testing redirect..."
curl -I http://localhost:8000/$CODE 2>&1 | grep -i location

echo -e "\n✅ Test complete! Try visiting: http://localhost:8000/$CODE"
```
## 👤 Author

**Moe-Salim91156**

- GitHub: [@Moe-Salim91156](https://github.com/Moe-Salim91156)

## 🙏 Acknowledgments

- Built with Go's standard library
- Inspired by URL shortening services like bit.ly and tinyurl.com

---

⭐ Star this repository if you find it helpful!
