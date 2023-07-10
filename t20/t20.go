package t20

import (
	"fmt"
	"unicode"
)

func reverseWords(s string) string {
	chars := []rune(s)
	for unicode.IsSpace(chars[0]) {
		chars = chars[1:]
	}
	reverseString(chars, 0, len(chars)-1)
	reverseWordsInString(chars)
	return string(trimSpaces(chars))
}

func reverseString(chars []rune, left, right int) {
	for left < right {
		chars[left], chars[right] = chars[right], chars[left]
		left++
		right--
	}
}

func reverseWordsInString(chars []rune) {
	n := len(chars)
	start := 0

	for i := 0; i < n; i++ {
		if unicode.IsSpace(chars[i]) {
			reverseString(chars, start, i-1)
			start = i + 1
		} else if i == n-1 {
			reverseString(chars, start, i)
		}
	}
}

func trimSpaces(chars []rune) []rune {
	slow := 0
	fast := 0
	n := len(chars)

	for fast < n {
		for fast < n && unicode.IsSpace(chars[fast]) {
			fast++
		}
		for fast < n && !unicode.IsSpace(chars[fast]) {
			chars[slow] = chars[fast]
			slow++
			fast++
		}
		if fast < n {
			chars[slow] = ' '
			slow++
		}
	}

	return chars[:slow]
}

func DoTask() {
	s := "cat dog slow fast"
	s = reverseWords(s)
	fmt.Println(s)
}
