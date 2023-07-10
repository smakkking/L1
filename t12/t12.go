package t12

import (
	"fmt"
	"math/rand"
	"time"
)

func generate_subset(arr []string, salt int64) []string {
	result := make([]string, 0, len(arr)-1)

	s1 := rand.NewSource(salt - 100)
	r1 := rand.New(s1)
	n := r1.Intn(1<<len(arr) - 1)

	i := 0
	for n > 0 {
		if n%2 == 1 {
			result = append(result, arr[i])
		}
		i++
		n /= 2
	}

	return result

}

func DoTask() {
	arr := []string{"cat", "dog", "cat", "cat", "tree"}

	t := generate_subset(arr, int64(time.Now().Unix()))

	fmt.Println(t)
}
