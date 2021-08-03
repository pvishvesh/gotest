package main

import "fmt"

func main() {
	fmt.Println(add(6, 4))
	//fmt.Println(split(64))
	//fmt.Println("reversing my string...", add(7,8))
	//a,b := swap(".....","??????")
	//fmt.Println(a,b)
}

func split(num int) (a, b int) {
	a = num % 10
	b = num - a
	return
}

func swap(a, b string) (string, string) {

	return b, a
}

func add(a, b int) int {
	return a + b
}

func mult(s, d int) int {
	return s * d
}
