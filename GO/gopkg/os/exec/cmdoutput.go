package main

import (
	"fmt"
	"os/exec"
)

func main() {
	arg := []string{"123", "abc"}
	cmd := exec.Command("mv", arg...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("output file : %s\n", output)
}
