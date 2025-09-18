package main

import "fmt"

func main() {
	fmt.Println(hello())
	fmt.Println("hello")
	fmt.Println("go")
}

func hello() string {
	return "Hello Go" // эту строку исправили Go теперь не go
}
