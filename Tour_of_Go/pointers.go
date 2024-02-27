package main

import "fmt"

func main() {
	i, j := 55, 27

	p := &i
	fmt.Println(*p)

	*p = 12
	fmt.Println(i)

	p = &j
	*p = *p / 3
	fmt.Println(j)
}
