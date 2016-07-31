package main

import (
	"fmt"
	"log"
	"os"
)

var debug = true

func submit(n int) {
	sum := 0
	for i := 0; i < n; i++ {
		sum += i
	}
	if debug {
		fmt.Printf("summit(%d) = %d\n", n, sum)
	}

}

func submitAndReturn(n int) int {
	sum := 0
	for i := 0; i < n; i++ {
		sum += i
	}
	return sum
}

func who_Am_I_Now() {
	host, _ := os.Hostname()
	if debug {
		fmt.Printf("1 host = %s\n", host)
	}
}

func main() {
	submit(10)
	//	who_Am_I_Now()
	x := submitAndReturn(101)
	host, errCode := os.Hostname()
	if errCode != nil {
		log.Panic("don't know my own computer name")
	}
	if debug {
		fmt.Printf("%s %d \n", host, x)
	}
}
