FROM golang:1.16-alpine

RUN apk add chromium

WORKDIR /app

COPY ../../src /app

RUN go build -o prerender ./cmd/server/main.go

EXPOSE 50051

CMD ["./prerender"]
