package t26

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
	is_unique([]rune("abcde"))
}
