package t16

import "fmt"

// https://blog.boot.dev/golang/quick-sort-golang/

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func quickSort(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}
	return arr
}

func DoTask() {
	arr := []int{1, 234, 23, 1, 15, 284, 38, 41}

	arr = quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}