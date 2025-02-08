# Multiplicación de Matrices en Distintos Lenguajes

Este repositorio contiene implementaciones de un algoritmo de multiplicación de matrices en los siguientes lenguajes de programación:

- Python
- Java
- C++
- JavaScript
- Go
- Rust

Cada implementación genera dos matrices de 100x100 con valores aleatorios entre 1 y 10 y realiza la multiplicación.

## Estructura del Proyecto

El código de cada lenguaje se encuentra en su respectiva carpeta:

```
├── python/
│   ├── Script.py
├── java/
│   ├── Script.java
├── cpp/
│   ├── Script.cpp
├── javascript/
│   ├── Script.js
├── go/
│   ├── Script.go
├── rust/
│   ├── Script.rs
├── Dockerfiles/
│   ├── Dockerfile.python
│   ├── Dockerfile.java
│   ├── Dockerfile.cpp
│   ├── Dockerfile.javascript
│   ├── Dockerfile.go
│   ├── Dockerfile.rust
└── README.md
```


## Cómo Ejecutar los Programas con Docker

Para cada lenguaje hay un `Dockerfile` en la carpeta `Dockerfiles/`. Sigue los siguientes pasos para ejecutar cada versión en su propio contenedor.

### Python
```bash
docker build -t matrix_python -f Dockerfiles/Dockerfile.python .
docker run --rm matrix_python
```

### Java
```bash
docker build -t matrix_java -f Dockerfiles/Dockerfile.java .
docker run --rm matrix_java
```

### C++
```bash
docker build -t matrix_cpp -f Dockerfiles/Dockerfile.cpp .
docker run --rm matrix_cpp
```

### JavaScript (Node.js)
```bash
docker build -t matrix_js -f Dockerfiles/Dockerfile.javascript .
docker run --rm matrix_js
```

### Go
```bash
docker build -t matrix_go -f Dockerfiles/Dockerfile.go .
docker run --rm matrix_go
```

### Rust
```bash
docker build -t matrix_rust -f Dockerfiles/Dockerfile.rust .
docker run --rm matrix_rust
```

## Requisitos
- Tener instalado Docker.
- Clonar este repositorio:
  ```bash
  git clone https://github.com/tu-usuario/multiplicacion-matrices.git
  cd multiplicacion-matrices
  ```
