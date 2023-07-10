package t15

import "os"

var justString string

func someFunc() {
	// проблема - строка может быть очень большая по размеру и не поместиться в память
	file, err := os.Create("very_big_string.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// _, err := createHugeString(1<<10, file) // здесь мы постепенно записываем ее в файл
}

func DoTask() {
	someFunc()
}
