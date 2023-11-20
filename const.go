package main

func main() {
	const pi float64 = 3.14159
	pi = 3.1 // ошибка
	const (
		pi float64 = 3.14159
		e  float64 = 2.7182
	)
	const pi, e = 3.14, 2.72
	const n int // Ошибка
	var m int = 7
	const k = m // Ошибка
}
