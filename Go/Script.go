package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func multiplyMatrices(A, B [][]int) ([][]int, error) {
	rowsA, colsA := len(A), len(A[0])
	rowsB, colsB := len(B), len(B[0])

	if colsA != rowsB {
		return nil, errors.New("Las matrices no se pueden multiplicar")
	}

	result := make([][]int, rowsA)
	for i := range result {
		result[i] = make([]int, colsB)
	}

	for i := 0; i < rowsA; i++ {
		for j := 0; j < colsB; j++ {
			for k := 0; k < colsA; k++ {
				result[i][j] += A[i][k] * B[k][j]
			}
		}
	}

	return result, nil
}

func main() {
	rand.Seed(time.Now().UnixNano())

	A := make([][]int, 100)
	B := make([][]int, 100)
	for i := range A {
		A[i] = make([]int, 100)
		B[i] = make([]int, 100)
		for j := range A[i] {
			A[i][j] = rand.Intn(10) + 1
			B[i][j] = rand.Intn(10) + 1
		}
	}

	result, _ := multiplyMatrices(A, B)
}
