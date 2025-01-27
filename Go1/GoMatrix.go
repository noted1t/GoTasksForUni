package main

import (
	"fmt"
	"sync"
	"time"
)

// Функция умножения матриц
func multiplyMatrices(A, B [][]int) [][]int {
	n := len(A)
	m := len(B[0])
	p := len(B)
	C := make([][]int, n)
	for i := range C {
		C[i] = make([]int, m)
	}

	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			wg.Add(1) // Увеличиваем счетчик WaitGroup для каждой горутины
			go func(i, j int) {
				defer wg.Done() // Уменьшаем счетчик WaitGroup после завершения работы горутины
				for k := 0; k < p; k++ {
					C[i][j] += A[i][k] * B[k][j] // Вычисляем значение элемента результирующей матрицы
				}
			}(i, j)
		}
	}

	wg.Wait() // Ожидаем завершения всех горутин
	return C
}

func main() {
	A := [][]int{
		{1, -1},
		{2, 0},
		{3, 0},
	}

	B := [][]int{
		{1, 1},
		{2, 0},
	}

	start := time.Now() // Запускаем таймер
	C := multiplyMatrices(A, B)
	elapsed := time.Since(start) // Останавливаем таймер

	fmt.Println("Resultant Matrix:") // Выводим результирующую матрицу
	for _, row := range C {
		fmt.Println(row)
	}
	fmt.Printf("Time taken: %s\n", elapsed) // Выводим время выполнения
}
