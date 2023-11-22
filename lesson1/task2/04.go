package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Scan(&number)

	left := number/100000%10 + number/10000%10 + number/1000%10
	right := number/100%10 + number/10%10 + number%10
	if left == right {
		fmt.Println("Счастливчик")
	} else {
		fmt.Println("В сделующий раз повезет")
	}
}
