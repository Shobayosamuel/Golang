package main

import (
	"time"
	"fmt"
	// "runtime"
)

func main() {
	// go start()
	// fmt.Println("Started")
    // time.Sleep(1 * time.Second)
    // fmt.Println("Finished")

	// fmt.Println("Started")
	// for i := 0; i < 10; i++ {
	// 	go execute(i)
	// }
	// time.Sleep(time.Second * 2)
	// fmt.Println("Finished")
	// fmt.Println(runtime.NumCPU())

	// Anonymous goroutines
	go func() {
        fmt.Println("In Goroutine")
    }()

    fmt.Println("Started")
    time.Sleep(1 * time.Second)
    fmt.Println("Finished")

}

func start() {
	go start2()
	fmt.Println("In Goroutine")
}
func start2() {
	fmt.Println("In Goroutine2")
}
func execute(id int) {
	fmt.Println("id:", id)
}