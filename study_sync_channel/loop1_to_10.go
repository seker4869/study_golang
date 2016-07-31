package main

import "fmt"

var ch chan int = make(chan int, 1000)

func loop() {
	for i := 0; i < 1000; i++ {
		fmt.Printf("%d", i)

	}
	fmt.Println("\n")
	ch <- 0
}

func main() {
	for i := 0; i < 500; i++ {
		go loop()
		<-ch

	}

}
