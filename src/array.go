package main

import "fmt"

func main() {
	var a [2]string
	a[0]= "hello"
	a[1]= "world"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
	primes := []int{4,43,24,99,653,543}
	fmt.Println(primes)
	s := primes[0:2]
	t := primes[1:5]
	fmt.Println(s,t)
	s[1]= 999
	fmt.Println(s,t)	
	fmt.Println(primes)
	fmt.Printf("%T\n", primes)
	primes = primes[1:4]
	fmt.Println(primes)
	primes = primes[0:5]
	fmt.Println(primes)
}
