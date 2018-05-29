package main

import (
	"fmt"
)

type Test1 interface {
	Show()
}
type Test2 struct {
	Name string
}

func (t Test2) Show() {
	fmt.Println("laoliao")
}

func main() {
	//	t1 := Test2{"hahaah"}
	var t2 *Test2
	var t2 Test1

	f(t2)
}

func f(t Test1) {
	if t != nil {
		t.Show()
		fmt.Print("not nil")
		return
	}
	fmt.Println("nil")
	return
}
