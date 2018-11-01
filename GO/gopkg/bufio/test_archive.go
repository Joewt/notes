package main

import (
	"archive/zip"
	"fmt"
)

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	rc, err := zip.OpenReader("Archive.zip")
	handleError(err)
	defer rc.Close()
	for _, f := range rc.File {
		fmt.Println(f.FileInfo().Name())
	}
}
