FROM golang:1.16-alpine

WORKDIR /app

COPY ./src /app

RUN go build -o storage ./cmd/storage/main.go

EXPOSE 3000

CMD ["./storage"]
