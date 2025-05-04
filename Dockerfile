# syntax=docker/dockerfile:1

FROM golang:1.24.2 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o failgen ./cmd/failgen

FROM gcr.io/distroless/static-debian11:nonroot
USER nonroot:nonroot
COPY --from=builder /app/failgen /
EXPOSE 8080
ENTRYPOINT ["/failgen"]
