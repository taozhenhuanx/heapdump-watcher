FROM golang:1.21-alpine AS builder  
WORKDIR /app  
COPY go.mod go.sum ./  
RUN go mod download  
COPY . .  
RUN CGO_ENABLED=0 GOOS=linux go build -o /heapdump-watcher  
  
FROM alpine:3.18  
RUN apk add --no-cache ca-certificates  
WORKDIR /app  
COPY --from=builder /heapdump-watcher ./heapdump-watcher  
CMD ["/heapdump-watcher"]