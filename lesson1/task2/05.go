package main

import "fmt"

func main() {
	var number, num int
	fmt.Scan(&number)

	sum := 0
	for i := 0; i < number; i++ {
		fmt.Scan(&num)

		if num >= 10 && num < 100 && num%8 == 0 {
			sum += num
		}
	}
	fmt.Println(sum)
}
