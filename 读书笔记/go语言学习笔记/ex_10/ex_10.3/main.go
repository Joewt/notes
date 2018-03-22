package main

import (
	"fmt"
	"./ex_10.3/structPack"
)

func main() {
	struct1 := new(structPack.ExpStruct)
	struct1.Mi1 = 10
	struct1.Mf1 = 16
	fmt.Printf("Mi1 = %d\n", struct1.Mi1)
}