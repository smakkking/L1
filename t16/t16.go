package t16

import "fmt"

// https://blog.boot.dev/golang/quick-sort-golang/

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			// все элементы большие pivot выкидываем во 2-ую часть
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	// в итоге получается, что элемент по индексу i - граница: слева от него значения < pivot, справа - >=
	arr[i], arr[high] = arr[high], arr[i]
	return i
}

func quickSort(arr []int, low, high int) {
	if low < high {
		var p int
		p = partition(arr, low, high)

		// сортируем 2 частии по полученной границе
		quickSort(arr, low, p-1)
		quickSort(arr, p+1, high)
	}
}

func DoTask() {
	arr := []int{1, 234, 23, 1, 15, 284, 38, 41}

	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
