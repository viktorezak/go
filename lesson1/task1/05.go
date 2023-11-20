package main

import (
	"fmt"
)

func main() {
	var d int
	fmt.Print("Введите градусы от 0 до 360:")
	fmt.Scan(&d)

	if d > 360 || d < 0 {
		fmt.Println("К сожалению у нас нет таких градусов")
		main()
	}
	hours := d * 2 / 60
	minutes := d * 2 % 60
	fmt.Printf("Это %d часов %d минут.", hours, minutes)

}
