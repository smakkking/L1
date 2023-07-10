package t13

import "fmt"

func t_13_f1(a *int, b *int) {
	*a += *b // a = a + b
	*b -= *a // b = b - a - b = -a
	*a += *b // a = a + b - a = b
	*b = -*b // b = -(-a) = a
}

func t_13_f2(a *int, b *int) {
	(*a) ^= (*b) // a = a ^ b
	(*b) ^= (*a) // b = b ^ a ^ b = a
	(*a) ^= (*b) // a = a ^ b ^ a = b
}

func t_13_f3(a int, b int) (int, int) {
	return b, a
}

func DoTask() {
	a, b := 1, 100
	fmt.Println(a, b)
	t_13_f1(&a, &b)
	fmt.Println(a, b)
}
