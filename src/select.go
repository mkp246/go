package main

import "fmt"

func fib(c,quit chan int){
    x,y:=0,1
    for {
        select {
            case c<-x:
                x,y=y,x+y
            case <-quit:
                fmt.Println("quit")
                return
        }
    }
}
func main() {
    c:= make(chan int)
    quit:= make(chan int)
    go func(){
        for i := 0; i < 20; i++ {
            fmt.Println(<-c)
        }
        quit<-0
    }()
    fib(c,quit)
}
