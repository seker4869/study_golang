package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)

func reader(chin, chout chan int) {
	wg.Add(1)
	defer wg.Done()
	leftToDo := 5
	for {
		x := <-chin
		if x < 0 {
			return
		}
		fmt.Printf("reading - left to do = %d\n", leftToDo)
		leftToDo--
		time.Sleep(1 * time.Second)
		if leftToDo <= 0 {
			chout <- 0
			return
		}
		chout <- x
	}
}

func writer(chin, chout chan int) {
	wg.Add(1)
	defer wg.Done()
	for {
		x := <-chin

		if x < 0 {
			return
		}
		fmt.Printf("writing - left to do = %d\n", x)
		time.Sleep(1 * time.Second)
		if x == 0 {
			return
		}
		chout <- x
	}
}

func doit() {
	chin := make(chan int)
	chout := make(chan int)
	go writer(chout, chin)
	go reader(chin, chout)
	//go writer(chin, chout)

	chin <- 1
	wg.Wait()
}

func main() {
	fmt.Println("go\n")
	doit()
}
