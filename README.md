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

## Módulo 2: Variables, funciones y documentación

### Variables, constantes y zero values

### Operadores aritméticos

### Tipos de datos primitivos

### Paquete fmt: algo más que imprimir en consola

### Uso de funciones

### Go doc: La forma de ver la documentación