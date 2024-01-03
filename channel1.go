package main

import (
	"fmt"
	"time"
	// "time"
)

func main() {
	// ch := make(chan int)
	// go send(ch)
	// go receive(ch)

	// time.Sleep(time.Second * 1)
	// fmt.Println(ch)
	// ch <- 8
	// v := <-ch	
	// fmt.Println("received", v)
	// ch := make(chan int, 5)
    // ch <- 3
	// ch <- 3
	// fmt.Println("length:", len(ch))
    // fmt.Println("Sending value to channnel complete")
    // val := <-ch
    // fmt.Printf("Receiving Value from channel finished. Value received: %d\n", val)
	// fmt.Println("capacity:", cap(ch))
	// fmt.Println("length:", len(ch))

	ch := make(chan int, 3)
	ch <- 6
	ch <- 6
	ch <- 6
	close(ch)
	go sum(ch)
	time.Sleep(time.Second * 1)

}

func send(ch chan int)  {
	ch <- 1
	fmt.Println("End")
}

func receive(ch chan int)  {
	val := <-ch
	fmt.Println("Value Received=%d in receive function\n", val)
}

func sum(ch chan int) {
	sum := 0
	for val := range ch {
		sum += val
	}
	fmt.Println("Sum:", sum)
}