package main

func t_13_f1(a *int, b *int) {
	*a += *b
	*b -= *a
	*a += *b
	*b = -*b
}

func t_13_f2(a *int, b *int) {
	(*a) ^= (*b)
	(*b) ^= (*a)
	(*a) ^= (*b)
}

func t_13_f3(a int, b int) (int, int) {
	return b, a
}
