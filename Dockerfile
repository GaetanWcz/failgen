FROM golang:1.24.2 AS builder
WORKDIR /app
COPY . .
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init -g cmd/failgen/main.go -o docs && \
    CGO_ENABLED=0 GOOS=linux go build -o failgen ./cmd/failgen

FROM gcr.io/distroless/static-debian11:nonroot
USER nonroot:nonroot
COPY --from=builder /app/failgen /
COPY --from=builder /app/docs /docs
EXPOSE 8080
ENTRYPOINT ["/failgen"]