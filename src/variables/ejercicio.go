package main

import "fmt"

/*
 * En este ejercicio vamos a calcular el area de un cuadrado. Deberemos especificar
 * la longitud de sus lado en una variable llamada `lado` y deberemos imprimir en la consola
 * su area con el siguiente formato: "La superficie del cuadrado es de XXm2"
 */
func main() {
	// Declaramos la longitud de los lados del cuadrado
	lado := 12.4

	// Calculamos su area
	area := lado * lado

	// Imprimimos por consola su superficie
	fmt.Printf("La superficie del cuadrado es de %vm2\n", area)
}
