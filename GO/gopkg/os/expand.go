package main

import (
	"fmt"
	"os"
)

func main() {
	mapping := func(s string) string {
		m := map[string]string{"hello": "world", "go": "perfect program language"}
		return m[s]
	}

	str := "Golang is$not a $go in the ${hello}!"
	fmt.Printf("%s\n", os.Expand(str, mapping))
	s := "The language path is: $LANG"
	fmt.Printf("%s\n", os.ExpandEnv(s))

	f, err := os.Lstat("chdir.go")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%+v\n", f)

}
