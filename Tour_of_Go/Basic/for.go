package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	//for continued
	index := 1
	for index < 2 {
		index += index
	}
	fmt.Println(index)
}
