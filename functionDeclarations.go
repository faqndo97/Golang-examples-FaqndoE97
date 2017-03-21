package main

func main() {

}

// Esta es la manera en la que definimos funciones en Go
// func + nombre + parametros + tipo del retorno
// Se puede observar que este valor se puede omitir
func log(message string) {
}

// Para definir los parametros lo hacemos de la sigueinte manera
// nombre + tipo
func add(a int, b int) int {
}

// Otra caracteristica que nos permite Go en el tema de los parametros
// es que si por ejemplo tenemos dos parametros y ambos son del mismo tipo
// ponemos el tipo una unica vez y se infiere para ambos parametros
func remove(a, b int) int {
}

// Go nos perimte que una funcion nos devuelva multiples valores, lo que por ejemplo
// nos sirve para la asignacion multiple de variables
// En el caso de la siguiente funcion, podriamos hacer en algun momento:
// value, exist := power("algo")

// En ocaciones solo nos interesa conservar un unico valor de los que retorna
// una funcion asi, lo que podemos hacer para realizar esto es:
// _, variable := power("algo")
//
func power(name string) (int, bool) {
}
