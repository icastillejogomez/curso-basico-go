# Curso básico de Go

En este repositorio veremos los conceptos más básicos de Go (Golang) que cubren en este curso básico de Go. Los principales puntos que veremos serán:

Este contenido se comenzó a redactar el 1 de Abril del 2022. Presta atención pues aunque seguramente los contenidos aquí explicados no hayan sufrido cambios el *estado del arte* del lenguaje puede haber cambiado.

### Temario
1. Variables y constantes
2. Funciones y estructuras de datos básicas
3. Go-rutines
4. Go-modules
5. Librerías para backend
6. Tips

#### Autor del proyecto

Este repositorio con el curso básico de Go ha sido elaborado por Ignacio F. Castillejo Gómez.

#### Gestión de cambios de este proyecto

Si deseas proponer algún cambio en la documentación o código de este proyecto no dudes en abrir una *pull request* o *issue*.


## Módulo 1: Hola mundo en Go

### ¿Qué es, por qué y quienes utilizan Go?

Go es un lenguaje compilado y estáticamente tipado. ¿Qué quiere decir esto? Bueno veámoslo por partes:

1. **Compilado**: cuando decimos que un lenguaje es compilado nos referimos a que la CPU de nuestro ordenador no es capaz de entender *perse* el código fuente que nosotros escribimos. Para poder ejecutar nuestros programas escritos en Go será necesario que **otro programa** transforme nuestro código fuente a otro tipo de código que la CPU **si es capaz de entender**, y por lo tanto de ejecutar. A este proceso se le llamada compilar. Es por esto que decimos que Go, es un lenguaje compilado.
2. **Estáticamente tipado**: cuando veamos las variables en Go, veremos que para guardar un dato (en una variable) será necesario especificar que necesitamos dicha variable **y su tipo**. Pero... veamos esto más detenídamente. Imagina que necesitas que un usuario se está registrando en tu aplicación y que para poder hacerlo debe enviarte su email y la contraseña. Tu necesitaras guardar estos dos datos en dos variables. Que un lenguaje sea estáticamente tipado quiere decir que tu **no solo** deberás programar estas dos variables sino que **tendrás que especificar el tipo de dato que va a guardar**. En el ejemplo anterior, ambas variables serán de tipo texto o como se dice en ingles (y en el mundo del software) serán de tipo **String**. Hay más tipos de datos como números, etc. pero eso lo veremos más adelante.

Pero... **¿Por qué muchas veces leemos Go como "Golang"?** Esto fue debido a que cuando pensaron que los desarrolladores buscarían información sobre Go en los buscadores, la palabra "Go" podría llevara  confusiones y por eso aunque le lenguaje se llama "Go" en internet todos hacemos referencia a él como "**Golang**".

Fue creado en Google por [Robert Griesemer](https://github.com/griesemer), [Rob Pike](https://twitter.com/rob_pike?lang=es) y [Ken Thompson](https://es.wikipedia.org/wiki/Ken_Thompson). El motivo de su creación fue ayudar a los desarrolladores e ingenieros de Google a ser capaces de llevar a cabo tareas extremadamente pesadas. Es por eso que crearon un lenguaje con la potencia de C pero con una sintaxis y experiencia de desarrollo como más cómoda. Es con esta idea como los ingenieros de Google concibieron Go.

Go fue anunciado en Noviembre del 2009 pero no fue hasta dos años más tarde cuando lanzaron la primera versión del lenguaje, en Marzo del 2012.

¿Por qué usar Go?

1. **Gran velocidad de compilación**: como hemos visto anteriormente Go es un lenguaje compiado, por lo que nosotros, tendremos que compilar nuestros códigos fuentes a un lenguaje que la CPU sea capaz de entender (código máquina). Go es verdaderamente rápido haciendo esta tarea.
2. **Alto rendimiento para tareas pesadas**: Go esta diseñado desde su origen para usar todos los nucleos de tu CPU por lo que será capaz de exprimir al máximo toda la potencia de tu servidor.
3. **Soporte nativo para concurrencia**: la concurrencia es la capacidad de un lenguaje de programación de ejectuar varias tareas al mismo tiempo. Esto que parece algo obvio no lo es tanto; no todos los lenguajes de programación tienen esta capacidad. Go fue diseñado para soportar nativamente los retos concurrentes y esto es una gran ventaja para los desarrolladores de este lenguaje.
4. **Obliga a implementar buenas practicas**: esto es algo que irás viendo en más detalle a los largo de este curso pero podriamos decir que Go fue diseñador para que sus desarrolladores, las personas que crean software en Go, se vean de alguna manera *forzados* a usar buenas prácticas de programación lo que hará que tu código fuente sea más fácil de entender por otros ingenieros o desarrolladores.

### Instalar Go 

Para instalar Go en tu máquina te recomiendo que visites la [página web oficial de Go](https://go.dev/) y sigas los pasos que allí se te indiquen para realizar la instalación de Go correctamente.

También es posible instalar Go a través de un gestor de paquetes. Por ejemplo en el caso de usar macOS puedes usar `brew`:

```
$ brew install go
```

Para comprobar que has instalado correctamente el binario de Go puedes ejecutar en tu terminal el comando `version`:
```
$ go version
go version go1.18 darwin/amd64
```

### Nuestras primeras lineas de código

Es hora de hacer nuestro primer programa escrito en Go. Para ello vamos a desarrollar el programa Hola Mundo que simplementa va a escribir en consola la palabra "Hello World".

```go
package main

import "fmt"

func main() {
  fmt.Println("Hello World")
}
```
Puedes ver el archivo en la ruta [src/hello_world/main.go](src/hello_world/main.go)

Veamos detalladamente el código anterior:

1. En primer lugar vemos que tenemos en la primera linea `package main` esto es para indicar el package de dicho archivo. Como el archivo se llama `main.go` usaremos la palabra **main** para describir el paquete.
2. A continuación vemos en la linea #3 `import "fmt"`. Esta linea indica a nuestro programa que tiene que importar la librería `fmt` para que podamos usarla posteriormente. 
3. En el resto del archivo vemos que hemos definido una función llamada `main`. Por convención, al ejecutar un programa escrito en Go, el programa iniciará su ejecución por dicha función `main` por lo que si cambiáramos este nombre nuestro programa no funcionaría ya que no sabría por donde debe de empezar.
4. Dentro de nuestra función `main` ejecutamos la siguiente instrucción: `fmt.Println("Hello World")`. Lo que estamos haciendo en esta linea es ejecutar la función `Println` dentro de la librería `fmt` y pasarla como primer argumento la cadena de texto `Hello World`

Para poder ejectuar el siguiente código debemos compilarlo primero. Para ello vamos a ejecutar el siguiente comando (ten en cuenta que si has creado una estructura de carpetas diferente deberas adaptar el siguiente comando a ella):

```
$ go build -o src/hello_world/main src/hello_world/main.go
```

Para ejecutar nuestro programa lo haremos de la siguiente manera:
```
$ src/hello_world/main
Hello World
```

¿Aburrido esto de estar compilando a cada rato para probar piezas de código tan pequeñas verdad? Go nos ofrece el comando `run` que nos compila en una carpeta temporal nuestro código y a la vez lo ejecuta para hacernos la vida más fácil. La menera de hacerlo es la siguiente:

```
$ go run src/hello_world/main.go
Hello World
```

Para terminar esta sección veamos un ejemplo de algo que vimos anteriormente... ¿Te acuerdas que te dije que Go nos *obliga* a usar buenas práctica de programación? Como sabrás muchos programadores en vez de escribir las funciones de esta manera:

```go
func functionName() {
  // Cuerpo de la función
}
```

lo hacen de esta manera:

```go
func functionName()
{
  // Cuerpo de la función
}
```

Te dejo como ejercicio que pruebes a crear un programa Hola Mundo en Go de esa manera y veas el resultado que obtienes por consola a la hora de ejecutarlo o compilarlo.

## Módulo 2: Variables, funciones y documentación
### Variables, constantes y zero values

#### Constantes

En esta sección vamos a ver esas pequeñas cajas de memoria que nos van a permitir guardar datos en ellas. Para comenzar veamos un ejemplo de constantes.

Si tienes dudas a cerca de la función `Printf` puedes visitar [este enlace](https://www.educative.io/edpresso/how-to-use-the-printf-function-in-golang).

Las constantes son variables cuyo valor no puede cambiar a lo largo de la ejecución de un programa. Puede parecer contraintuitivo *limitarnos* a no poder cambiar el valor de una variable pero hay innumerables ocasiones en las que esto tiene muchísimo sentido. Piensa por ejemplo... en un programa que hace algún tipo de cálculo matemático que requiere el valor de PI (3.14). Este es un valor constante en las matemáticas ya que PI absolutamente nunca cambiará de valor. En esta caso te es útil poder declarar la variable como constante para evitarte la posibilidad de cometer un error en el futuro y editar el valor de PI a un número incorrecto.

```go
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
```
Puedes encontrar este archivo en [/src/variables/constantes.go](/src/variables/constantes.go)

#### Variables

En este archivo vamos a definir variables normales. Es decir, aunque el tipo de dato que van a guardar será fijo, su valor si podrá cambiar a lo largo del programa.

```go
package main

import "fmt"

func main() {
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
}
```
Puedes encontrar este archivo en [/src/variables/variables.go](/src/variables/variables.go)


Es importante tener en cuenta que si declaramos una variable que no usamos en ningún punto de nuestro programa el compilador de Go fallará. (Prueba a descomentar el siguiente código)

```go
package main

func main() {
  var variableInutil int = 12
}
```
#### Valores por defecto

Hemos visto en el archivo anterior `variables.go` que no podemos declarar variables que no usemos. Ahora vamos a ver que pasa si creamos variables que SI usamos pero que no establecemos un valor inicial.

En una clase posterior vamos a ver detalladamente todos los tipos de datos que podemos manejar con Go pero ahora nos vamos a centrar en los siguientes:	 
1. **Valores enteros**: estos son números sin decimales	 
2. **Valores en coma flotante**: estos son valores con decimales	 
3. **Valores booleanos**: los valores booleanos (que vienen del álgebra de Bool) solo pueden ser verdadero (true) o falso (false)	 
4. **Strings**: estos son cadenas de texto o literales

```go
package main

import "fmt"

func main() {
  var entero int
  var flotante float64
  var booleano bool
  var texto string

  fmt.Printf("El valor entero por defecto es: %d\n", entero)
  fmt.Printf("El valor de coma flotante por defecto es: %f\n", flotante)
  fmt.Printf("El valor booleano por defecto es: %t\n", booleano)
  fmt.Printf("La cadena de texto por defecto es: \"%s\"\n", texto)
}
```
Puedes encontrar este archivo en [/src/variables/valores_por_defecto.go](/src/variables/valores_por_defecto.go)


#### Ejercicio

En este ejercicio deberás calcular la superficie (area) de un cuadrado.

Puedes ver la solución a este ejercicio en el archivo [/src/variables/ejercicio.go](/src/variables/ejercicio.go)

Si quieres profundizar el los conceptos básicos de la algoritmia numérica te propongo la siguiente pregunta:

    En el ejercicio anterior, si estableces la longitud 
    de los lados del cuadrado en 12.4 verás que el resultado 
    que arroja Go al multiplicar dos veces dicho valor 
    es: 153.76000000000002 cuando el resultado realmente es 153.76

    ¿Por qué sucede esto? ¿Por qué Go nos añade 0.00000000000002 al resultado?

### Operadores aritméticos

Los operadores aritméticos son esos operadores que te permitirán operar valores numéricos como la suma, la resta, etc. En go podrás usar los siguientes operadores aritméticos:
1. **La suma**: Para sumar dos números usaremos el operador `+`
2. **La resta**: Para restar dos números usaremos el operador `-`
3. **La multiplicación**: Para multiplicar dos números usaremos el operador `*`
4. **La división**: Para dividir dos números usaremos el operador `/`
5. **El módulo**: Para obtener el módulo de una división usaremos el operador `%`

A continuación veamos unos ejemplos de los operadores anteriores:

```go
package main

import "fmt"

func main() {
	const x = 10
	const y = 10
	const suma = x + y

	fmt.Printf("La suma de X + Y es: %v + %v = %v\n", x, y, suma)
}
```
Puedes ver este archivo en la ruta [/src/operadores_aritmeticos/suma.go](/src/operadores_aritmeticos/suma.go)


### Tipos de datos primitivos

### Paquete fmt: algo más que imprimir en consola

### Uso de funciones

### Go doc: La forma de ver la documentación

## Módulo 3: Estructuras de control de flujo y condicionales

### El poder de los ciclos en Golang: for, for...while y for...forever

### Operadores lógicos y de comparación

### El condicional if

### Múltiples condiciones anidadas con Switch

### Keywords: defer, break y continue

## Módulo 4: Estructuras de datos básicas

### Arrays y Slices

### Recorrido de Slices con Range

### Clave-valor con Maps

### Structs: La forma de hacer clases en Go

### Modificadores de acceso en funciones y Structs

## Módulo 5: Métodos e interfaces

### Structs y Punteros

### Stringers: personalizar el output de Structs

### Interfaces y listas de interfaces

## Módulo 6: Concurrencia y Channels

### ¿Qué es la concurrencia?

### Primer contacto con las Goroutines
### Channels: la forma de organizar las goroutines

## Módulo 7: Manejo de paquetes y Go Modules

### go get: El manejador de paquetes

### go modules: Ir más allá del GoPath con Echo

### Modificando módulos con Go
