package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cmd := "cat /etc/passwd"
	out, err := exec.Command("sh", "-c", cmd).CombinedOutput()
	if err != nil {
		fmt.Printf("err is %s", err)
		os.Exit(1)
	}

	fmt.Printf(string(out))
	os.Exit(0)

}
