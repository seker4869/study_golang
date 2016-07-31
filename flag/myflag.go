package main

import "flag"
import "fmt"

func main() {
	haha := flag.String("canshu", "My first flag", "Usage <go run myslag.go canshu XXXXXX>")
	flag.Parse()
	fmt.Println(*haha)
}
