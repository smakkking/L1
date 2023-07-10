package t19

import "fmt"

func reverseString(chars []rune) {
	left, right := 0, len(chars)-1
	for left < right {
		// обмениваемся символами, пока не дойдем до середины
		chars[left], chars[right] = chars[right], chars[left]
		left++
		right--
	}
}

func DoTask() {
	t := []rune("aaabbb")

	reverseString(t)
	fmt.Println(string(t))
}
