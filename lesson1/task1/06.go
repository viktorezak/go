package main

import (
	"fmt"
)

func main() {
	var n, number uint
	fmt.Print("Введите количество элементов от 0 до 100:")
	fmt.Scan(&n)

	if n > 100 || n < 0 {
		fmt.Println("Мы не можем такое считать =(")
		main()
	}

	items := make([]int, n)
	for i := range items {
		fmt.Scan(&items[i])
	}

	for _, x := range items {
		if x > 0 {
			number++
		}
	}

	fmt.Printf("количество положительных элементов в последовательности %d ", number)
}
