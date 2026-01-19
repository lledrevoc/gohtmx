# Build stage
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .

# Run stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
# Copy HTMX/static files if they aren't embedded in your binary
COPY --from=builder /app/static ./static 
COPY --from=builder /app/templates ./templates
EXPOSE 8080
CMD ["./main"]