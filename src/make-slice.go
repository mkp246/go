package main
import "fmt"

func main() {
	var s []int
	printSlice(s)
	s= append(s,0)
	printSlice(s)
	s= append(s,1)
	printSlice(s)
	s= append(s, 2,3,5,8)
	printSlice(s)
	for i,v :=range(s) {
		fmt.Printf("2**%d=%d\n",i,v)
	}
	for _,v :=range(s) {
		fmt.Printf("%d\n",v)
	}
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s),cap(s),s)
}