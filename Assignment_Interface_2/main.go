package main

import (
	"fmt"
	"os"
)

func main() {
	// f, err := os.Open(os.Args[1])
	// if err != nil {
	// 	fmt.Print("Error", err)
	// 	os.Exit(1)
	// }
	// io.Copy(os.Stdout, f)
	fmt.Println(os.Args[1])
}
