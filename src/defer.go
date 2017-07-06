package main
import "fmt"

func main() {
	defer fmt.Println("Hello world 2")
	defer fmt.Println("Hello world 1")
	fmt.Println("hello")
	defer fmt.Println("Hello world 3")
}