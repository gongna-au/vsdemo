package main

import "fmt"

func main() {
	n := 18
	p := &n

	fmt.Println(p)
	fmt.Printf("%T\n", p)
	m := *p
	fmt.Println(m)

}
