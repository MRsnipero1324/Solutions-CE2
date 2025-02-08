# Dockerfile para Java
FROM openjdk:17-alpine

WORKDIR /app

COPY ../Java/Script.Java /app/

RUN javac Script.Java

CMD ["java", "Script"]
