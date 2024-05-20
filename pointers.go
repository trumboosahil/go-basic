package main

import "fmt"

func main() {

	var i int32
	i = 11
	fmt.Printf("the value of i is: %v", i)
	var p *int32 = new(int32)
	p = &i
	*p = 88
	fmt.Println()
	fmt.Printf("the value of i is: %v", i)
}
