use rand::Rng;

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

    let start = std::time::Instant::now();
    let result = multiply_matrices(&A, &B).unwrap();
    let duration = start.elapsed();

    // Crear carpeta de salida y guardar tiempo en un archivo
    std::fs::create_dir_all("output").unwrap();
    std::fs::write(
        "output/execution_time_rust.txt",
        format!("Tiempo de ejecución: {:?} ms", duration.as_millis()),
    )
    .unwrap();
}
