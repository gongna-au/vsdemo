package main

import "fmt"

func main() {

	scoreMap := make(map[string]int)
	scoreMap["zhangsan"] = 100
	scoreMap["lisi"] = 200
	for k, v := range scoreMap {

		fmt.Println(k, v)

	}

}
