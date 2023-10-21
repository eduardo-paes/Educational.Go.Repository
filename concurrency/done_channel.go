package main

import (
	"fmt"
	"time"
)

func doWork(done <- chan bool) {
	for {
		select {
		case <- done:
			fmt.Println("DONE")
			return
		default:
			fmt.Println("DOING WORK")
			time.Sleep(time.Second * 1)
		}
	}
}

func main() {

	done := make(chan bool)

	go doWork(done)

	time.Sleep(time.Second * 10)

	close(done)

	time.Sleep(time.Second * 1)

}