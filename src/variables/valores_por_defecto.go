package main

import "fmt"

func main() {
	// Hemos visto en el archivo anterior `variables.go` que no podemos declarar
	// variables que no usemos. Ahora vamos a ver que pasa si creamos variables que
	// SI usamos pero que no establecemos un valor inicial.

	// En una clase posterior vamos a ver detalladamente todos los tipos de datos
	// que podemos manejar con Go pero ahora nos vamos a centrar en los siguientes:
	/*
	 * 1. Valores enteros: estos son números sin decimales
	 * 2. Valores en coma flotante: estos son valores con decimales
	 * 3. Valores booleanos: los valores booleanos (que vienen del álgebra de Bool) solo pueden ser verdadero (true) o falso (false)
	 * 4. Strings: estos son cadenas de texto o literales
	 */

	var entero int
	var flotante float64
	var booleano bool
	var texto string

	fmt.Printf("El valor entero por defecto es: %d\n", entero)
	fmt.Printf("El valor de coma flotante por defecto es: %f\n", flotante)
	fmt.Printf("El valor booleano por defecto es: %t\n", booleano)
	fmt.Printf("La cadena de texto por defecto es: \"%s\"\n", texto)
}
