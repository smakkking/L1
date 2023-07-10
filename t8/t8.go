package t8

// пояснение
// x ^ 1 = ~x,
// x ^ 0 = x

func f1(num int64, bit_number int, bit_value int) int64 {
	var b int64 = 1 << bit_number

	if (num&b)>>bit_number == int64(bit_value) {
		return num
	} else {
		return num ^ b
	}
}

func f2(num int64, bit_number int) int64 {
	// инвертирует бит в позиции N
	var b int64 = 1 << bit_number
	return num ^ b
}

func DoTask() {

}
