function multiplyMatrices(A, B) {
    let rowsA = A.length, colsA = A[0].length;
    let rowsB = B.length, colsB = B[0].length;

    if (colsA !== rowsB) {
        throw new Error("Las matrices no se pueden multiplicar");
    }

    let result = Array.from({ length: rowsA }, () => Array(colsB).fill(0));

    for (let i = 0; i < rowsA; i++) {
        for (let j = 0; j < colsB; j++) {
            for (let k = 0; k < colsA; k++) {
                result[i][j] += A[i][k] * B[k][j];
            }
        }
    }
    return result;
}

// Prueba con matrices 100x100
let A = Array.from({ length: 100 }, () => Array.from({ length: 100 }, () => Math.floor(Math.random() * 10) + 1));
let B = Array.from({ length: 100 }, () => Array.from({ length: 100 }, () => Math.floor(Math.random() * 10) + 1));
let result = multiplyMatrices(A, B);
