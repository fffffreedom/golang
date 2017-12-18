package main

import "fmt"

func main () {
	m1 := make(map[string] interface{}, 16)
	m1["name"] = "jonny"
	m1["age"] = 99.0

	fmt.Println(m1["name"].(string))
	fmt.Println(m1["age"].(float64))
}
