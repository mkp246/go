package main
import "fmt"
// import "math"
var c, python, java = true, true, "dmg"

func main(){
	fmt.Println("Sum is:", add(66,4329))
	// c :=321
	a,b := swap("hello", "world")
	fmt.Println("swapping in go:",a,b)
	fmt.Println("named returns: ")
	fmt.Println(split(17))
	fmt.Println(c, python, java)
	//=======
	var i int
	var f float64
	var k bool
	var s string
	fmt.Println(i,f,k,s)
}

func add(x , y int) int {
	return x+y
}
func swap(x,y string)(string,string){
	return y,x
}
func split(sum int)(x,y int){
	x = sum*4/9
	y = sum - x
	return
}