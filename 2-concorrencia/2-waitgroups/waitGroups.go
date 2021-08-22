package main

import (
	"fmt"
	"sync"
	"time"
)


func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go func() {
		escrever("Olá mundo!")
		waitGroup.Done()
	}()


	go func() {
		escrever("Tudo bem ?")
		waitGroup.Done()
	}()

	waitGroup.Wait()
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}