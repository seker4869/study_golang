package main

import "fmt"
import "io/ioutil"

func main() {
	myconf := "/etc/passwd"
	bs, err := ioutil.ReadFile(myconf)
	if err != nil {
		fmt.Println("there is something wrong %s", err.Error())
	}
	for _, value := range bs {
		fmt.Printf(string(value))
	}
}
