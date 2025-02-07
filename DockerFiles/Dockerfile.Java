# Dockerfile para Java
FROM openjdk:17-slim

WORKDIR /app

COPY script.java /app/

RUN javac script.java

CMD ["java", "script"]
