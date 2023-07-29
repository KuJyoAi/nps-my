package main

import "fmt"

func main() {
	m := make(map[string]int, 1)
	for k, v := range m {
		fmt.Println("11")
		fmt.Println(k, v)
	}

}
