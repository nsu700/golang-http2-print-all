# Stage 1: Build the Go application
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod init test && go mod tidy
RUN go build -o app .

# Stage 2: Create the final image
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/app .
EXPOSE 8080
CMD ["./app"]
