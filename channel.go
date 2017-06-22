package main

import "fmt"

func sum(s []int, c chan int) {
    sum := 0
    for _, v := range (s) {
        sum += v
    }
    c <- sum
}
func main() {
    s := []int{5, 6, 2, 7, -8, 9}
    c := make(chan int)
    go sum(s[:len(s) / 2], c)
    go sum(s[len(s) / 2:], c)
    x, y := <-c, <-c
    fmt.Println(x, y, x + y)
    d:= make(chan string)
    go func(){d <- "all routines done"}() //self execute anonymous function
    fmt.Println(<-d)
}
