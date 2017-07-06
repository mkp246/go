package main
import "fmt"

func main(){
	sum:=0
	for i:=0;i<=10;i++ {
		sum+=i
	}
	fmt.Printf("%d\n",sum)
	
	for sum<200 {
		fmt.Println(sum)
		sum++
	}
	if sum<200 {
		fmt.Println("number less than 200")
	} else {
		fmt.Println("number greater than 200")
	}
}