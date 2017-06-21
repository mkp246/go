package main
import "fmt"
import "math"

type Abser interface{
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a Abser
	f:=MyFloat(-math.Sqrt2)
	v:=Vertex{4,3}
	a=f
	a=&v
	//a=v fails
	fmt.Println(a.Abs())
}