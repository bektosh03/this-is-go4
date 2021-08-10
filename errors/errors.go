package main

import (
	"fmt"
)

type MyError struct {
	value string
}

func (m MyError) Error() string {
	return m.value
}

func returnErr(num int) error {
	var (
		err1 = MyError{value: "my error"}
		err2 = MySecondError{}
		err3 MyStringError = "some error"
	)
	if num > 0 {
		return err1
	}
	if 4 > 5 {
		return err3
	}
	return err2
}

func main() {
	err := returnErr(5)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = returnErr(-1)
	fmt.Println(err)

	var a MyStringError = "string"
	var str string = "string"

	if a == MyStringError(str) {

	}
	fmt.Println(a, string(56))
}

type MySecondError struct{
}

func (m MySecondError) Error() string {
	return "error"
}

// =========


type MyStringError string

func (m MyStringError) Error() string {
	return string(m)
}

const (
	MyConstError = MyStringError("some error")
)
