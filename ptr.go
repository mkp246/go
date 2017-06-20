package main
import "fmt"

func main() {
	i,j := 42, 2701
	p := &i
	fmt.Println(p)
	fmt.Println(*p)
	*p=21
	fmt.Println(i,*p,p)

	p = &j
	fmt.Println(p)
	*p = *p/32
	fmt.Println(j,*p, p)
	// ======================
	v:= Vertex{1,2}
	fmt.Println(v)
	v.X = 3
	fmt.Println(v.X,v.Y,v)
	pv := &v
	(*pv).X =9
	pv.X =91
	fmt.Println(v,pv)
}

type Vertex struct {
	X int
	Y int
}