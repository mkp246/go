package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	ans := 1.0
	var delta float64 = (ans*ans -x) / 2*ans
	for ;delta>0.001; {
		ans-=delta
		delta=(ans*ans -x) / 2*ans
	}
	return ans
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(3))
	fmt.Println(Sqrt(4))
	fmt.Println(Sqrt(5))
}
