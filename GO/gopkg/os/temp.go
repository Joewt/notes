package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("The temp dir is: %s\n", os.TempDir())

}
