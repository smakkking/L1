package t1

import "fmt"

func (a *Action) do_intherance(...interface{}) {

}

type Human struct {
	weight int
	name   string
} // структура может быть с произвольным числом методов и полей

func (a *Human) walk() {
	fmt.Printf("%s walks\n", a.name)
}

type Action struct {
	human Human
} // чтобы можно было использовать a.walk(), а не a.human.walk()

func DoTask() {

}
