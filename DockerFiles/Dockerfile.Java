# Usar una imagen base con Java
FROM openjdk:17-jdk-slim

# Establecer el directorio de trabajo
WORKDIR /app

# Copiar el código fuente al contenedor
COPY ../Java/Script.java /app/Script.java

# Crear un directorio de salida para almacenar los resultados
RUN mkdir -p /output

# Compilar el código fuente
RUN javac Script.java

# Ejecutar el programa
CMD ["java", "Script"]
