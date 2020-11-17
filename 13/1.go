package main

import "fmt"

func main() {

	a := make(map[string]int)
	a["zhangsan"] = 1
	a["lisi"] = 2

	k, ok := a["zhangsan"]

	if ok {
		fmt.Println(k)

	} else {

		fmt.Println("查无此人")

	}

}
