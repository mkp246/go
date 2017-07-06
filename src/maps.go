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
	fmt.Println(m1)
}

var m1 = map[string]Vertex{
	"Bell Labs": {
		40.68433, -74.39967,
	},
	"Google": {
		37.42202, -122.08408,
	},
}
