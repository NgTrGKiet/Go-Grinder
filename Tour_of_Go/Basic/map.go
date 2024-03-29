package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func main() {
	m := map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	m["alal"] = Vertex{30, -99.225}
	delete(m, "Google")
	fmt.Println(m)

}
