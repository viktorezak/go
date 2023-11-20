package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Print("Введите число от 0 до 10000:")
	fmt.Scan(&number)

	if number > 10000 || number < 0 {
		fmt.Println("К сожалению число больше 10000 или меньше 0")
		main()
	}
	fmt.Println("Число десятков:")
	fmt.Println((number % 100) / 10)

}
