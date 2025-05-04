# FailGen

**FailGen** is a simple HTTP API that simulates various HTTP errors. It's ideal for testing frontend error handling, chaos engineering, and HTTP client resilience.

## 🔧 Features
- Simulate HTTP error responses (status code, delay, message)
- Customizable delay in milliseconds
- JSON response with metadata
- Ready-to-deploy with Docker (distroless)

## 🚀 Usage

curl http://localhost:8080/fail?status=500&delay=1000&message=Server+Overloaded

### 🔍 Query Parameters
| Parameter | Type   | Default | Description                     |
|-----------|--------|---------|---------------------------------|
| status    | int    | 500     | HTTP status code to return     |
| delay     | int    | 0       | Delay in milliseconds          |
| message   | string | status text | Custom error message     |

## 🔬 Development

### Run locally

To generate the Swagger documentation automatically, you need the swag CLI tool:

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```
Make sure your $PATH includes $HOME/go/bin:

```bash
export PATH="$HOME/go/bin:$PATH"
```
Once installed, swag init will run automatically when starting the app.

```bash
go mod init failgen
go run ./cmd/failgen
```

### Build and run with Docker
docker build -t failgen .
docker run -p 8080:8080 failgen

## 📘 API Documentation (Swagger)

After running the app, the Swagger documentation is available under the docs/ directory.

You can preview it at:
- JSON spec: http://localhost:8080/swagger/doc.json
- Swagger UI: http://localhost:8080/swagger/index.html

## 📄 License
MIT
