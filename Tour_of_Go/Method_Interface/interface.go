package main

import (
	"fmt"
)

type I interface {
	M() string
}

type T struct {
	S string
}

func (t *T) M() string {
	return t.S
}

type F string

func (f F) M() string {
	return string(f)
}

func describe(i I) {
	fmt.Println(i.M())
}

func main() {
	t := &T{"Hello"}
	describe(t)

	f := F("ssflksd")
	describe(f)
}
