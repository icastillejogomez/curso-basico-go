package main

import "fmt"

func main() {
	// En este archivo vamos a definir variables normales. Es decir, aunque el tipo de dato que van
	// a guardar será fijo, su valor si podrá cambiar a lo largo del programa.

	// Para definir una variable usaremos la keyword `var` seguido el tipo de dato que va a guardar
	var username string = "Satoshi Nakamoto"
	fmt.Printf("El nombre de usuario seleccionado es: %s\n", username)

	// Otra forma de definir una variable es de la forma `<nombre> := valor`
	// En esta caso estamos dejando a Go inferir el tipo de dato que va a guardar la variable
	// Es importante tener en cuenta que cuando usamos el operador `:=` la variable no ha sido
	// declarada con anterioridad ya que la expresión `:=` solo nos sirve para declarar por
	// primera vez una variable
	password := "I'm the CIA"
	fmt.Printf("La contraseña del usuario es: %s\n", password)

	// Como estas variables son "normales", podemos cambiar el valor de las mismas cuando queramos.
	// Para hacer esto usaremos el operador `=`. Como podrás observar no usa los dos puntos `:`
	password = "Who knows..."
	fmt.Printf("La nueva contraseña del usuario es: %s\n", password)

	// Es importante tener en cuenta que si declaramos una variable que no usamos en
	// ningún punto de nuestro programa el compilador de Go fallará. (Prueba a descomentar el siguiente código)
	// var variableInutil int = 12
}
