use rand::Rng;
use std::fs;
use std::time::Instant;

fn multiply_matrices(A: &Vec<Vec<i32>>, B: &Vec<Vec<i32>>) -> Result<Vec<Vec<i32>>, &'static str> {
    let rows_a = A.len();
    let cols_a = A[0].len();
    let rows_b = B.len();
    let cols_b = B[0].len();

    if cols_a != rows_b {
        return Err("Las matrices no se pueden multiplicar");
    }

    let mut result = vec![vec![0; cols_b]; rows_a];

    for i in 0..rows_a {
        for j in 0..cols_b {
            for k in 0..cols_a {
                result[i][j] += A[i][k] * B[k][j];
            }
        }
    }

    Ok(result)
}

fn main() {
    let mut rng = rand::thread_rng();
    let A: Vec<Vec<i32>> = (0..100).map(|_| (0..100).map(|_| rng.gen_range(1..=10)).collect()).collect();
    let B: Vec<Vec<i32>> = (0..100).map(|_| (0..100).map(|_| rng.gen_range(1..=10)).collect()).collect();

    // Medir tiempo de ejecución
    let start_time = Instant::now();
    let _result = multiply_matrices(&A, &B).unwrap();
    let duration = start_time.elapsed();

    // Convertir duración a milisegundos
    let execution_time_ms = duration.as_millis();

    // Crear carpeta /output y escribir el archivo execution_time.txt
    let output_dir = "/output";
    fs::create_dir_all(output_dir).unwrap();
    let output_file = format!("{}/execution_time_rust.txt", output_dir);
    fs::write(output_file, format!("Tiempo de ejecución: {} ms\n", execution_time_ms)).unwrap();
}
