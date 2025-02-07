#include <iostream>
#include <vector>
#include <cstdlib>

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

    vector<vector<int>> result = multiplyMatrices(A, B);
}
