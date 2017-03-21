package main

import (
  "fmt"
)

func main() {
  // Con esto declaramos la variable power del tipo int
  // Go por defecto cuando declaramos una variable x de un tipo y le asigna
  // el valor por defecto de ese tipo y a la variable x hasta que se le asigne
  // otro valor
  // type      / default val
  // int       / 0
  // boolean   / false
  // Strings   / ""
  var power int
  power = 9000
  // podriamos unir ambas lineas haciendo un
  // var power int = 9000
  // otra forma seria realizar una asignacion sin tipo, el tipo lo infiere el
  // compilador cuando compila el archivo
  // power := 9000
  // Esto funcionaria tambien si le asignaramos una funcion a la variable power
  // si hicieramos esto sin tipo infiere el valor de la variable en base a lo
  // que retorne dicha funcion
  // Una variable no puede ser declarada mas de una vez en el mismo scope

  // Go nos permite realizar una asignacion multiple de variables
  name, age := "facundo", 20
  fmt.Printf("It´s over %d\n", power)
  fmt.Printf("Author: %s age: %d", name, age)

  // Es importante remarcar que go no nos permitira tener variables sin usar
}
