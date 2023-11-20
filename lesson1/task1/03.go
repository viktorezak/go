package main

import (
	"fmt"
	"os"
)

func main() {
	var number int
	fmt.Print("Введите число до 10000:")
	fmt.Scan(&number)

	if number > 10000 {
		fmt.Println("К сожалению число больше 10000")
		os.Exit(1)
	}
	fmt.Println("Последнее число:")
	fmt.Println(number % 10)

	for number > 9 {
		number /= 10
	}
	fmt.Println("Первое число:")
	fmt.Println(number)

}
