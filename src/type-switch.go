package main 
import "fmt"

func do(i interface{}) {
	switch v:= i.(type){
	case int:
		fmt.Println("twice is :",2*v)
	case string:
		fmt.Println("length is :",len(v))
	default:
		fmt.Printf("unknown type %T",v)
	}
}

func main() {
	do(21)
	do("hello")
	do(66.09)
}
