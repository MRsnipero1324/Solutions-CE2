# Matrix Multiplication in Different Languages

This repository contains implementations of a matrix multiplication algorithm in the following programming languages:

- Python
- C++
- JavaScript
- Go
- Rust

Each implementation generates two 100x100 matrices with random values between 1 and 10 and performs the multiplication.

## Project Structure

The code for each language is located in its respective folder:

```
├── python/
│   ├── Script.py
├── cpp/
│   ├── Script.cpp
├── javascript/
│   ├── Script.js
├── go/
│   ├── Script.go
├── rust/
│   ├── Script.rs
├── DockerFiles/
│   ├── Dockerfile.python
│   ├── Dockerfile.cpp
│   ├── Dockerfile.javascript
│   ├── Dockerfile.go
│   ├── Dockerfile.rust
└── README.md
```


## How to Run the Programs with Docker

Each language has a corresponding `Dockerfile` located in the `DockerFiles/` folder. Follow these steps to run each version in its own container.

### Python
```bash
docker build -t matrix_python -f DockerFiles/Dockerfile.Python .
docker run --rm -v $(pwd)/output:/output matrix_python
```

### C++ 
```bash
docker build -t matrix_cpp -f DockerFiles/Dockerfile.cpp .
docker run --rm -v $(pwd)/output:/app/output matrix_cpp
```

### JavaScript (Node.js) 
```bash
docker build -t matrix_js -f DockerFiles/Dockerfile.javascript .
docker run --rm -v $(pwd)/output:/output matrix_js
```

### Go 
```bash
docker build -t matrix_go -f DockerFiles/Dockerfile.go .
docker run --rm -v $(pwd)/output:/output matrix_go
```

### Rust 
```bash
docker build -t matrix_rust -f DockerFiles/Dockerfile.rust .
docker run --rm -v $(pwd)/output:/app/output matrix_rust
```

## Output  
After running each container, an execution_time_* file will be created or modified in the output/ folder, where * is the programming language used. The file will contain the execution time for that particular solution in milliseconds.

For example:
- execution_time_python.txt
- execution_time_cpp.txt
- execution_time_js.txt
- execution_time_go.txt
- execution_time_rust.txt

These files will contain the recorded execution times for each respective language solution.

