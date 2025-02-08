# Dockerfile para C++
FROM gcc:12

WORKDIR /app

COPY ../Cpp/Script.cpp /app/

RUN g++ Script.cpp -o Script

CMD ["./Script"]
