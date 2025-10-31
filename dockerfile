# ---------- BUILD STAGE ----------
    FROM golang:1.23-alpine AS builder

    WORKDIR /app
    
    # Copy dependency files first
    COPY go.mod go.sum ./
    RUN go mod download
    
    # Copy the full source
    COPY . .
    
    # Build the Go binary
    RUN CGO_ENABLED=0 GOOS=linux go build -o server .
    
    # ---------- RUN STAGE ----------
    FROM alpine:latest
    
    
    WORKDIR /root/
    
    # Copy the pre-built binary from the builder stage
    COPY --from=builder /app/server .
    
    # Expose the app port
    EXPOSE 8080
    
    # Run the app
    CMD ["./server"]
    