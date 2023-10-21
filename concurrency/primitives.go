package main

import (
	"fmt"
)

func main() {

	// Channel: Canal de comunicação entre rotinas - área de transferência
	myChannel1 := make(chan string)
	myChannel2 := make(chan string)

	// Go Routines: Executa funções em threads separadas
	go func () {
		myChannel1 <- "data1"
	}()
	go func () {
		// time.Sleep(time.Second * 1)
		myChannel2 <- "data2"
	}()

	// Select: Utiliza o canal que estiver pronto primeiro
	select {
	case msgFromMyChannel1 := <- myChannel1:
		fmt.Println(msgFromMyChannel1)
	case msgFromMyChannel2 := <- myChannel2:
		fmt.Println(msgFromMyChannel2)
	}

}