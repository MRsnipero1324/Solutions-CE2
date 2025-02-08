package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func multiplyMatrices(A, B [][]int) ([][]int, error) {
	rowsA, colsA := len(A), len(A[0])
	rowsB, colsB := len(B), len(B[0])

	if colsA != rowsB {
		return nil, errors.New("las matrices no se pueden multiplicar")
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
	// Semilla para el generador de números aleatorios
	rand.Seed(time.Now().UnixNano())

	// Crear matrices 100x100
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

	// Medir tiempo de ejecución
	start := time.Now()
	_, err := multiplyMatrices(A, B)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	duration := time.Since(start).Milliseconds()

	// Crear carpeta output y escribir archivo
	outputDir := "/output"
	os.MkdirAll(outputDir, os.ModePerm)

	filePath := fmt.Sprintf("%s/execution_time_go.txt", outputDir)
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error al crear el archivo:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("Tiempo de ejecución: %d ms\n", duration))
	if err != nil {
		fmt.Println("Error al escribir en el archivo:", err)
		return
	}

	fmt.Printf("Tiempo de ejecución guardado en: %s\n", filePath)
}
