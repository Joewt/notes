package main

import (
	"fmt"
	"os/exec"
)

func main() {
	path, err := exec.LookPath("php")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s\n", path)
}
