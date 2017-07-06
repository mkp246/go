package main
import (
	"fmt"
	"runtime"
	"time"
	)

func main() {
	fmt.Print("GO runs on ")
	switch os:=runtime.GOOS; os {
		case "darwin":
			fmt.Println("OS X")
		case "windows":
			fmt.Println("windows")
		case "linux":
			fmt.Println("linux")
		default:
			fmt.Println("Os not known")
	}
	fmt.Println(time.Now().Weekday())
	fmt.Println(time.Saturday)

	switch t:=time.Now(); {
		case t.Hour() < 12:
			fmt.Println("good morning")
		case t.Hour() < 17:
			fmt.Println("good afternoon")
		default:
			fmt.Println("good eveming..")
	}
	
}