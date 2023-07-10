package t23

import "fmt"

func del_elem(arr []int, index int) []int {
	if !(0 <= index && index < len(arr)) {
		return arr
	}

	for i := index; i < len(arr)-1; i++ {
		arr[i] = arr[i+1]
	}

	return arr[:len(arr)-1]
}

func main() {
	arr := []int{1, 2, 3, 4, 5}

	fmt.Println(arr)

	arr = del_elem(arr, 0)
	fmt.Println(arr)
}
