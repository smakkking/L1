package t17

import (
	"errors"
	"fmt"

	"golang.org/x/exp/constraints"
)

func bin_find[T constraints.Ordered](arr []T, x T, l int, r int) (int, error) {
	if l == r {
		if arr[l] != x {
			return l, errors.New("no found error!")
		} else {
			return l, nil
		}
	}

	q := (l + r) / 2

	// в зависимости от значения медианы ищем в разных частях
	if arr[q] < x {
		return bin_find(arr, x, q+1, r)
	} else if arr[q] > x {
		return bin_find(arr, x, l, q-1)
	} else {
		return q, nil
	}
}

func DoTask() {
	arr := []int{1, 1, 15, 23, 38, 41, 234, 284}
	fmt.Println(bin_find(arr, 23, 0, len(arr)-1))
}
