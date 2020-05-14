package main

import (
	"fmt"
	"os/exec"
)

func main() {
	arg := []string{"Hello", "World"}
	cmd := exec.Command("echo", arg...)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("output %s\n", output)
}
