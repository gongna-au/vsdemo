package main

import "fmt"

func main() {
	a1 := [...]int{1, 2, 3, 4, 5, 6}
	s1 := a1[:]
	s1 = append(s1[:1], s1[2:]...)
	fmt.Println(s1)
	fmt.Println(a1)

}
