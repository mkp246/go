package main
import "fmt"

func adder() func(int)int {
	sum:=0
	return func (x int)int {
		sum+=x
		return sum
	}
}

func main() {
	pos,neg := adder(),adder()
	for i := 0; i < 11; i++ {
		fmt.Println(pos(i))
		fmt.Println(neg(-i))
	}
}