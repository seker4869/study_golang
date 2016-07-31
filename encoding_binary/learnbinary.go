package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	var pi float64 = 3.141592653589793
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, pi)
	if err != nil {
		fmt.Println("binary write failed :", err)
	}
	fmt.Println(buf.Bytes())
}
