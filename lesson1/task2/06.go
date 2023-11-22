package main

import "fmt"

func main() {
	var x, p, y, years int
	fmt.Print("Введите сумму вклада:")
	fmt.Scan(&x)

	fmt.Print("Введите процент:")
	fmt.Scan(&p)

	fmt.Print("Хотите сколько получить:")
	fmt.Scan(&y)

	for ; x < y; years++ {
		x += x * p / 100
	}

	fmt.Println(years)
}
