package main

import (
	"fmt"
	"time"
)

func receive(c <-chan int) {
	for {
		time.Sleep(10 * time.Second)
		a, ok := <-c
		if !ok {
			fmt.Println("done")
			break
		}
		fmt.Printf("|| receive %d ||\n", a)
	}
}

func send(c chan<- int) {
	for i := range [10]int{} {
		fmt.Printf(">> sending %d <<\n", i)
		c <- i
		fmt.Printf(">> sent %d <<\n", i)
	}
	close(c)
}

func main() {
	// defer db.Close()
	// cli.Start()
	c := make(chan int, 10)
	go send(c)
	receive(c)
}
