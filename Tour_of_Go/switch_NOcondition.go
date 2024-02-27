package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Chao buoi sang")
	case t.Hour() < 17:
		fmt.Println("Chao buoi chieu")
	default:
		fmt.Println("Chao buoi toi")
	}
}
