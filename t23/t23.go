package t23

import "fmt"

// https://golang-blog.blogspot.com/2020/04/remove-element-from-slice-in-golang.html

// с сохранением порядка
func del_elem(arr []int, index int) []int {
	if !(0 <= index && index < len(arr)) {
		return arr
	}

	for i := index; i < len(arr)-1; i++ {
		arr[i] = arr[i+1]
	}

	return arr[:len(arr)-1]
}

// без сохранения порядка
func del_elem2(arr []int, index int) []int {
	arr[index] = arr[len(arr)-1]
	arr[len(arr)-1] = 0
	return arr[:len(arr)-1]
}

func DoTask() {
	arr := []int{1, 2, 3, 4, 5}

	fmt.Println(arr)

	arr = del_elem(arr, 0)
	fmt.Println(arr)
}
