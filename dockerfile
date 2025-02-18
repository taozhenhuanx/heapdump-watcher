FROM golang:alpine AS builder
ENV GOPROXY=https://goproxy.cn/

WORKDIR /go/release

RUN apk update && apk add tzdata

COPY go.mod ./go.mod
RUN go mod tidy
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o heapdump-watcher .

# 
FROM alpine

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories \
    && apk update \
    && apk add --no-cache tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && mkdir -p  /app/conf/

WORKDIR /app

COPY --from=builder /go/release/heapdump-watcher /app/

COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

CMD ["/app/heapdump-watcher"]
