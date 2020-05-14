package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t.Format(time.UnixDate))
	fmt.Println(time.UnixDate)
	d, _ := time.ParseDuration("3h4m2s")
	fmt.Println(d.Hours())
	fmt.Println(t.In(time.UTC))
	t = time.Date(2013, time.March, 14, 15, 23, 4, 100, time.UTC)
	fmt.Println(t)
	fmt.Println(t.Local())
	fmt.Println(time.Now().Location())
	fmt.Println(time.Now().Date())
	d, _ = time.ParseDuration("2s3ms")
	fmt.Println(d.Nanoseconds())

}
