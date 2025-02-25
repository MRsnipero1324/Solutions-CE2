
FROM golang:1.20-alpine

WORKDIR /app

COPY ../Go/Script.go /app/

RUN go build -o Script Script.go

CMD ["./Script"]
