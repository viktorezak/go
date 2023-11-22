package main

import "fmt"

func main() {
	var num int
	var seen [10]bool

	for {
		_, err := fmt.Scanf("%1d", &num)

		if err != nil {
			fmt.Println("Разные")
			break
		} else if seen[num] {
			fmt.Println("Не разные")
			break
		}

		seen[num] = true
	}
}
