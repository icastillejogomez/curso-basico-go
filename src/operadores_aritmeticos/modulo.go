package main

import "fmt"

func main() {
	const x = 5
	const y = 2
	const modulo = x % y

	fmt.Printf("El resto de dividir X / Y es: %v + %v = %v\n", x, y, modulo)
}
