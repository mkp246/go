package main 
import "fmt"

type Person struct {
	Name string
	Age int
}

func (p Person) String() string{
	return fmt.Sprintf("%v - %v years",p.Name,p.Age)
}
func main() {
	a:=Person{"aurther reed", 62}
	z:=Person{"led zipplin", 43}
	fmt.Println(a,z)
}
