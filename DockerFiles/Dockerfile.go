# Dockerfile para Go
FROM golang:1.20

WORKDIR /app

COPY script.go /app/

RUN go build -o script script.go

CMD ["./script"]
