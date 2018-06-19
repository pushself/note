package main

import (
	"fmt"
)

type client chan<- string

var (
	messages = make(chan client)
)

func main() {
	go Test()
	//  var msg string
	//  msg = "abcdefgh"
	ch := make(chan string, 1)
	ch <- "dsfsdf"
	messages <- ch
	<-ch
	close(messages)
}

func Test() {
	msg := <-messages
	fmt.Println(<-msg)
	close(msg)
}
