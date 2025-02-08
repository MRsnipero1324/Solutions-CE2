# Dockerfile para Go
FROM golang:1.20

WORKDIR /app

COPY ../Go/Script.go /app/

RUN go build -o Script Script.go

CMD ["./Script"]
