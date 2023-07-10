package t14

import (
	"fmt"
	"reflect"
)

func foo([]interface{}) {}

func get_type(obj interface{}) {
	fmt.Println(reflect.TypeOf(obj).Kind())

	// https://habr.com/ru/articles/415171/
	// интерфейс содержит в себе 2 части:
	// 1. само значение
	// 2. базовый тип этого значения
	// пустой интерфейс может содержать любое значение и всю информацию, которая нам когда-либо понадобится от него.
}

func DoTask() {
	var val []int
	get_type(val)
}
