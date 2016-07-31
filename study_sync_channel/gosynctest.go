package main

import (
	"fmt"
	"sync"
)

var (
	wg sync.WaitGroup
)

func Afunction(shownum int) {
	fmt.Println(shownum)
	wg.Done()
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go Afunction(i)
	}
	wg.Wait()
}
