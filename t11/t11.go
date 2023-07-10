package t11

import "fmt"

func intersect[T comparable](arr1 []T, arr2 []T) []T {
	var result []T
	if len(arr1) < len(arr2) {
		result = make([]T, 0, len(arr1))
	} else {
		result = make([]T, 0, len(arr2))
	}

	is_used1 := make([]bool, len(arr1))
	is_used2 := make([]bool, len(arr2))

	for i := 0; i < len(arr1); i++ {
		for j := 0; !is_used1[i] && j < len(arr2); j++ {
			if !is_used2[j] && arr1[i] == arr2[j] {
				result = append(result, arr1[i])
				is_used1[i] = true
				is_used2[j] = true
			}
		}
	}

	return result

}

func DoTask() {
	r := []int{1, 2, 3, 4, 5}
	s := []int{1, 2, 100, 4, 6}

	fmt.Println(intersect(r, s))
}
