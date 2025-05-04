# FailGen

**FailGen** is a simple HTTP API that simulates various HTTP errors. Useful for testing error handling in frontends or APIs.

## Usage

```bash
curl http://localhost:8080/fail?status=500&delay=1000&message=Server+Overloaded
```

## Query Parameters
- `status`: HTTP status code (default: 500)
- `delay`: Delay in ms before sending the response
- `message`: Custom message in the response

## Run locally
```bash
go run ./cmd/failgen
```

## Docker
```bash
docker build -t failgen .
docker run -p 8080:8080 failgen
```
