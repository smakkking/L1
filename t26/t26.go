package t26

import "fmt"

func is_unique(str []rune) bool {
	table := make(map[rune]bool)

	for _, elem := range str {
		_, ok := table[elem]
		if ok {
			return false
		}
		table[elem] = true
	}

	return true
}

func DoTask() {
	fmt.Println(is_unique([]rune("abcde")))
}
