package main

import "fmt"

func loop() {
	for i := 0; i < 1000; i++ {
		fmt.Printf("%d", i)

	}
	fmt.Println("\n")

}

func main() {
	for i := 0; i < 500; i++ {
		loop()

	}

}
