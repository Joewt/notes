package main

import (
	"fmt"
	"os"
)

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("current pwd is : %s\n", pwd)
	if err = os.Chdir("exec"); err != nil {
		fmt.Println(err)
		return
	}

	pwd, err = os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("current pwd is : %s\n", pwd)

}
