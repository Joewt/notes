package main

import (
	"fmt"
	"strings"
)

func rot13(r rune) rune {
	switch {
	case r >= 'A' && r <= 'Z':
		return 'A' + (r-'A'+13)%26
	case r >= 'a' && r <= 'z':
		return 'a' + (r-'a'+13)%26
	}
	return r
}

func main() {
	fmt.Println(strings.Contains("joejoesssfsdfasdf", "jo"))
	fmt.Println(strings.Contains("joewteaslfkjasldf", "jowtttt"))
	fmt.Println(strings.ContainsRune("乔fsdf", rune('乔')))
	fmt.Println(strings.Count("joewwwww", "w"))
	fmt.Println("%q", strings.Fields("jo sd sdf as fs   sdf   "))
	fmt.Println(strings.HasPrefix("joesdfasfa", "joe"))
	fmt.Println(strings.HasSuffix("joesdfasfa", "fa"))
	s := []string{"a", "b", "c"}
	fmt.Println(strings.Join(s, " , "))
	fmt.Println(strings.Map(rot13, "'Twas brillig and the slithy gopher..."))
}
