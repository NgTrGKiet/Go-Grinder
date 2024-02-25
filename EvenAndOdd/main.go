package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	for _, num := range numbers {
		if num%2 == 0 {
			fmt.Println(num, " is a even number")
		} else {
			fmt.Println(num, " is a odd number")
		}
	}
}
