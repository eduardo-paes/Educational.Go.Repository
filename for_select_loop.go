package main

import (
	"fmt"
	"time"
)

func main() {

	// Input data
	chars := []string { "a", "b", "c" }
	
	// Fan-In
	myChannel := make(chan string, 3)
	go func() {
		for _, s := range chars {
			select {
			case myChannel <- s:
			}
		}
		close(myChannel)
	}()

	// Fan-Out
	go func() {
		for result := range myChannel{
			fmt.Println(result)
		}
	}()

	// Wait process
	time.Sleep(time.Second * 10)

}