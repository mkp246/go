package main 
import "fmt"
import "math"

func compute(fn func(float64,float64)float64) float64{
	return fn(3,4)
}

func main() {
	hypot:=func(x,y float64) float64 {
		return math.Sqrt(x*x+y*y)
	}
	// fmt.Println(compute(hypot(8,12)))
	fmt.Println(hypot(8,12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}