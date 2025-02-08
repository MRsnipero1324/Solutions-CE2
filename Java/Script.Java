import java.io.File;
import java.io.FileWriter;
import java.io.IOException;
import java.util.Random;

public class MatrixMultiplication {
    public static int[][] multiplyMatrices(int[][] A, int[][] B) {
        int rowsA = A.length, colsA = A[0].length;
        int rowsB = B.length, colsB = B[0].length;

        if (colsA != rowsB) {
            throw new IllegalArgumentException("Las matrices no se pueden multiplicar");
        }

        int[][] result = new int[rowsA][colsB];

        for (int i = 0; i < rowsA; i++) {
            for (int j = 0; j < colsB; j++) {
                for (int k = 0; k < colsA; k++) {
                    result[i][j] += A[i][k] * B[k][j];
                }
            }
        }
        return result;
    }

    public static void main(String[] args) {
        Random rand = new Random();
        int[][] A = new int[100][100];
        int[][] B = new int[100][100];

        for (int i = 0; i < 100; i++) {
            for (int j = 0; j < 100; j++) {
                A[i][j] = rand.nextInt(10) + 1;
                B[i][j] = rand.nextInt(10) + 1;
            }
        }

        // Medir tiempo de ejecución
        long startTime = System.currentTimeMillis();
        int[][] result = multiplyMatrices(A, B);
        long endTime = System.currentTimeMillis();

        long executionTime = endTime - startTime; // Tiempo en milisegundos

        // Crear carpeta /output y escribir archivo execution_time.txt
        File outputDir = new File("/output");
        if (!outputDir.exists()) {
            outputDir.mkdirs();
        }

        File outputFile = new File(outputDir, "execution_time_java.txt");
        try (FileWriter writer = new FileWriter(outputFile)) {
            writer.write("Tiempo de ejecución: " + executionTime + " ms\n");
        } catch (IOException e) {
            System.err.println("Error escribiendo el archivo: " + e.getMessage());
        }
    }
}
