package main

import (
	"fmt"
	"time"
)

func main() {
	err := test("123")
	fmt.Println(err)
}

func test(str string) (err error) {
	type Test struct{}
	defer func() {
		switch p := recover(); p {
		case nil:
		case Test{}:
			err = fmt.Errorf("panic")
		default:
			panic(p)
		}
	}()

	if str == "" {
		panic(Test{})
	}
	panic("return")
	return nil
}
