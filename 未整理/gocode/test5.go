package main

import (
	"fmt"
	"os"
	//	"time"
)

var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		fmt.Println("done")
		return true
		//	default:
		//		fmt.Println("default")
		//		return false

	}
}

func main() {
	go func() {
		os.Stdin.Read(make([]byte, 1))
		fmt.Println("ok")
		close(done)
	}()
	cancelled()

	for {
		select {
		case <-done:
			fmt.Println("main")
			return
		}

	}

	//	time.Sleep(10000)
}
