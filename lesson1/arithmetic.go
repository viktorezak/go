package main

import "fmt"

func main() {
	a := 100
	b := 10
	c := a + b
	c = a * b
	c = a - b
	c = a / b
	c++
	c--
	var k int = 10 / 7
	var m float32 = 10.0 / 6
	var y int = 7 % 3
	fmt.Println(k, m, y)
}
