import random
import time  # Importar el módulo para medir tiempo

def multiply_matrices(A, B):
    rows_A, cols_A = len(A), len(A[0])
    rows_B, cols_B = len(B), len(B[0])
    
    if cols_A != rows_B:
        raise ValueError("Las matrices no se pueden multiplicar")

    result = [[0] * cols_B for _ in range(rows_A)]
    
    for i in range(rows_A):
        for j in range(cols_B):
            for k in range(cols_A):
                result[i][j] += A[i][k] * B[k][j]
    
    return result

# Medir tiempo de ejecución
start_time = time.time()

# Prueba con matrices 100x100
A = [[random.randint(1, 10) for _ in range(100)] for _ in range(100)]
B = [[random.randint(1, 10) for _ in range(100)] for _ in range(100)]
result = multiply_matrices(A, B)

end_time = time.time()
execution_time = (end_time - start_time) * 1000  # Convertir a milisegundos

# Guardar tiempo en un archivo .txt
with open("execution_time.txt", "w") as file:
    file.write(f"Tiempo de ejecución: {execution_time:.2f} ms")
