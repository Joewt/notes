package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("sleep", "10")
	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("waiting....")
	err = cmd.Wait()

	fmt.Println("finish")
}
