package main

import "fmt"

func print_help() {
	fmt.Println("just look")
}

func Afunction(ch chan int) {
	fmt.Println("finish")
	fmt.Println("test")
	print_help()
	<-ch
}

func main() {
	ch := make(chan int)

	go Afunction(ch)
	ch <- 1
}
