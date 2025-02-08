# Dockerfile para Java
FROM openjdk:17-slim

WORKDIR /app

COPY ../Java/Script.java /app/

RUN javac Script.java

CMD ["java", "Script"]
