package main

import (
	"fmt"
	"os"
)

func main() {
	attr := &os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout},
		Env:   os.Environ(),
	}

	p, err := os.StartProcess("/bin/pwd", []string{"/bin/pwd"}, attr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", p)
}
