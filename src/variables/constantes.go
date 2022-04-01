package main

import "fmt"

func main() {
	// Vamos a definir a continuación una variable constante. Una variable constante
	// es una variable cuyo valor no puede cambiar NUNCA.

	// Para declarar una constante vamos a usar la keyword `const` seguida del tipo de dato que va a almacenar
	const pi float64 = 3.1416
	fmt.Printf("El valor de PI es: %f\n", pi)

	// Aunque también podemos definir el tipo "al vuelo" dejando a Go que lo infiera.
	const e = 2.61
	fmt.Printf("El valor del número e es: %f\n", e)
}
