package main

import "fmt"

func myprint(x, y int) {
	fmt.Printf("X is %d\n", x)
	fmt.Printf("Y is %d\n", y)
}

func main() {

	x := 10
	y := 11
	myprint(y, x)
}
