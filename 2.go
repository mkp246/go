package main

import "fmt"
import "math"


func main(){
	var x,y int =3,4
	var f float64 = math.Sqrt(float64(x*x+y*y+1))
	var z= uint(f)
	fmt.Println(x,y,f,z)

	//---
	const i int = 32
	i =43
	j:=i
	fmt.Printf("j is type of %T\n",j)
}
