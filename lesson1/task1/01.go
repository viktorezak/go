package main

import "fmt"

func main() {
	var number int
	fmt.Print("Введите число:")
	fmt.Scan(&number)
	if number == 0 {
		fmt.Println("Вы ввели 0 или Скорее всего вы ввели не число:")
		number = number * 2
		fmt.Println(number)
		number = number + 100
		fmt.Println(number)
	} else {
		number = number * 2
		fmt.Println(number)
		number = number + 100
		fmt.Println(number)
	}
}
