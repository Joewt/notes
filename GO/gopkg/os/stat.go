package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Stat("chdir.go")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s(%o)\n", f.Mode(), f.Mode()&0777)

	err = os.Chmod("chdir.go", 0777)
	if err != nil {
		fmt.Println(err)
		return
	}

	f, err = os.Stat("chdir.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s(%o)\n", f.Mode(), f.Mode()&0777)
	fmt.Printf("%v\n", os.Environ())
	os.Exit(2)
}
