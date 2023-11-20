package main

import (
	"fmt"
	"math"
)

func main() {
	var number1 float64
	fmt.Print("Введите первое число:")
	fmt.Scan(&number1)

	var number2 float64
	fmt.Print("Введите второе число:")
	fmt.Scan(&number2)

	if number2 == 0 || number1 == 0 {
		fmt.Println("Вы ввели 0 или Скорее всего вы ввели не число:")
		result := math.Pow(number1, 2) + math.Pow(number2, 2)
		fmt.Println(result)
	} else {
		result := math.Pow(number1, 2) + math.Pow(number2, 2)
		fmt.Println(result)
	}
}
