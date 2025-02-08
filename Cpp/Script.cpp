#include <iostream>
#include <vector>
#include <cstdlib>
#include <fstream>
#include <chrono> // Para medir el tiempo de ejecución

using namespace std;

vector<vector<int>> multiplyMatrices(vector<vector<int>> &A, vector<vector<int>> &B) {
    int rowsA = A.size(), colsA = A[0].size();
    int rowsB = B.size(), colsB = B[0].size();

    if (colsA != rowsB) {
        throw invalid_argument("Las matrices no se pueden multiplicar");
    }

    vector<vector<int>> result(rowsA, vector<int>(colsB, 0));

    for (int i = 0; i < rowsA; i++) {
        for (int j = 0; j < colsB; j++) {
            for (int k = 0; k < colsA; k++) {
                result[i][j] += A[i][k] * B[k][j];
            }
        }
    }
    return result;
}

int main() {
    vector<vector<int>> A(100, vector<int>(100));
    vector<vector<int>> B(100, vector<int>(100));

    for (int i = 0; i < 100; i++) {
        for (int j = 0; j < 100; j++) {
            A[i][j] = rand() % 10 + 1;
            B[i][j] = rand() % 10 + 1;
        }
    }

    // Medir tiempo de ejecución
    auto start = chrono::high_resolution_clock::now();
    vector<vector<int>> result = multiplyMatrices(A, B);
    auto end = chrono::high_resolution_clock::now();

    // Calcular duración en milisegundos
    auto duration = chrono::duration_cast<chrono::milliseconds>(end - start).count();

    // Crear carpeta output y guardar el tiempo en un archivo
    ofstream outFile("output/execution_time_cpp.txt");
    if (outFile.is_open()) {
        outFile << "Tiempo de ejecución: " << duration << " ms" << endl;
        outFile.close();
    } else {
        cerr << "No se pudo abrir el archivo para guardar el tiempo de ejecución" << endl;
    }

    return 0;
}

