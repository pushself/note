package main

import (
	"fmt"
	"os"
)

type p func(string)

func main() {
	var pr p
	pr = Show
	Print("laoliao", pr)

}

func Show(str string) {
	fmt.Println(str)
	os.Stdout
}

func Print(str string, show func(string)) {
	show(str)
}
