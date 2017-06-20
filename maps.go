package main

import "fmt"
var m map[string]Vertex

type Vertex struct {
	Lat,Long float64
}

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{40.65,-74.44}

	fmt.Println(m)
}