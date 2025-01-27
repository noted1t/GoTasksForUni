package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Функция пузырьковой сортировки
func bubbleSort(arr []int, wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшаем счетчик WaitGroup после завершения работы
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j] // Меняем элементы местами
			}
		}
	}
}

// Функция сортировки вставками
func insertionSort(arr []int, wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшаем счетчик WaitGroup после завершения работы
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key // Вставляем ключевой элемент на правильное место
	}
}

// Функция быстрой сортировки
func quickSort(arr []int, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done() // Уменьшаем счетчик WaitGroup после завершения работы
	}
	if len(arr) < 2 {
		return
	}
	left, right := 0, len(arr)-1
	pivot := rand.Int() % len(arr) // Выбираем случайный элемент в качестве опорного
	arr[pivot], arr[right] = arr[right], arr[pivot]
	for i := range arr {
		if arr[i] < arr[right] {
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}
	}
	arr[left], arr[right] = arr[right], arr[left] // Разделяем массив на две части

	// Используем дополнительные WaitGroup только для параллельных вызовов
	if wg != nil {
		wg.Add(2)
		go quickSort(arr[:left], wg)   // Рекурсивная сортировка левой части
		go quickSort(arr[left+1:], wg) // Рекурсивная сортировка правой части
	} else {
		quickSort(arr[:left], nil)   // Рекурсивная сортировка левой части
		quickSort(arr[left+1:], nil) // Рекурсивная сортировка правой части
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, 1000)
	for i := range arr {
		arr[i] = rand.Intn(1000) // Заполняем массив случайными числами
	}

	// Создаем копии массива для каждой сортировки
	arr1 := make([]int, len(arr))
	arr2 := make([]int, len(arr))
	arr3 := make([]int, len(arr))
	copy(arr1, arr)
	copy(arr2, arr)
	copy(arr3, arr)

	var wg sync.WaitGroup
	wg.Add(3) // Устанавливаем счетчик WaitGroup для трех горутин

	// Параллельно запускаем три алгоритма сортировки
	go bubbleSort(arr1, &wg)
	go insertionSort(arr2, &wg)
	go quickSort(arr3, &wg)

	wg.Wait() // Ожидаем завершения всех горутин

	fmt.Println("Bubble Sorted: ", arr1)
	fmt.Println("Insertion Sorted: ", arr2)
	fmt.Println("Quick Sorted: ", arr3)
}
