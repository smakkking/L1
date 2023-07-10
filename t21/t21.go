package t21

import (
	"errors"
	"fmt"
	"math"
)

// исходный интерфейс "неизменяемый"
type MyInterface interface {
	SendRequest() float32
	ThrowException() error
	DoSomeLog()
}

// условный тип, удовлетворяющий интерфейсу
type S1 struct{}

func (s *S1) SendRequest() float32 {
	fmt.Println("request sended!")
	return math.Pi
}

func (s *S1) ThrowException() error {
	fmt.Println("exception generated!")
	return errors.New("shit")
}

func (s *S1) DoSomeLog() {
	fmt.Println("log created!")
}

// интерфейс, к которому нужно привести
type MyOtherInterface interface {
	Send() int
	Throw() error
	Log()
}

// преобразователь типа, в него мы будем засовывать сущ тип,
// на выходе получать тип, удовл нужному интерфейсу
type ToMyOtherInterface struct {
	*S1
}

func (t *ToMyOtherInterface) Send() int {
	return int(t.S1.SendRequest())
}

func (t *ToMyOtherInterface) Throw() error {
	return t.S1.ThrowException()
}

func (t *ToMyOtherInterface) Log() {
	t.S1.DoSomeLog()
}

func DoTask() {
	var _ MyInterface = &S1{}
	// var _ MyOtherInterface = &S1{} // error!
	var _ MyOtherInterface = &ToMyOtherInterface{&S1{}}
}
