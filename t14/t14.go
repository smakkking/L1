package t14

import (
	"fmt"
	"reflect"
)

func foo([]interface{}) {}

func get_type(obj interface{}) {
	fmt.Println(reflect.TypeOf(obj).Kind())
}

func DoTask() {
	var val []int
	get_type(val)
}
