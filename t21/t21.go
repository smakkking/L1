package t21

import (
	"errors"
	"fmt"
	"math"
)

type MyInterface interface {
	SendRequest() float32
	ThrowException() error
	DoSomeLog()
}

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

type MyOtherInterface interface {
	Send() int
	Throw() error
	Log()
}

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

func foo(s MyInterface) {}

func DoTask() {
	var _ MyInterface = &S1{}
	// var _ MyOtherInterface = &S1{} // error!
	var _ MyOtherInterface = &ToMyOtherInterface{&S1{}}
}
