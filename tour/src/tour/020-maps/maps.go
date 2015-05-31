package main 

import (
	"fmt"
)

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

var n = map[string]Vertex {
	"Bell Labs": {40.684, -74.39967},
	"Google": {37.422, -122.084},
}

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{ 40.25666, -74.22555 }
	fmt.Println(m["Bell Labs"])
	fmt.Println(n)
}

