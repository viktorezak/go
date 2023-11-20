package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Print("Введите число:")
	fmt.Scan(&number)

	fmt.Println("Ответ:")
	c := number % 2
	fmt.Println(number + 2 - c)

}
