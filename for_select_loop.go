package main

import (
	"fmt"
	"time"
)

func main() {

	myChannel := make(chan string, 3)
	chars := []string { "a", "b", "c" }

	go func() {
		for _, s := range chars {
			select {
			case myChannel <- s:
			}
		}
		close(myChannel)
	}()

	go func() {
		for result := range myChannel{
			fmt.Println(result)
		}
	}()

	time.Sleep(time.Second * 10)

}