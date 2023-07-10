package t17

import (
	"errors"

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

	if arr[q] < x {
		return bin_find(arr, x, q+1, r)
	} else {
		return bin_find(arr, x, l, q-1)
	}
}

func DoTask() {

}
