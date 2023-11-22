package main

import (
	"fmt"
	"strconv"
)

func main() {
	var number int
	fmt.Scan(&number)

	nStr := strconv.Itoa(number)

	fmt.Println(string(nStr[0]))
}
