package main

// Si realizaramos este import y no usaramos las librerias importadas Go nos
// daria un error, es muy estricto en este punto
import (
  "fmt"
  "os"
)

func main() {
  // Se esperan dos argumentos por que el primero siempre es la ruta del
  // ejecutable actual
  if len(os.Args) != 2 {
  // len es una funcion incorporada por Go que retorna el largo de un String,
  // asi como el numero de valores de un diccionario o el numero de elementos
  // de un array
    os.Exit(1)
  // Si quisieramos profundizar sobre lo que hace el metodo Exit o algun otro
  // podemos dirigirnos al sitio de go y leer la documentacion o ejecutar el
  // comando ** godoc -http=:6060 ** y luego ver la documentacion en
  // http://localhost:6060
  }

  // Se comprueba que en el primer argumento se guarda la ruta del ejecutable actual
  fmt.Println(os.Args[0])
  fmt.Println("It´s over", os.Args[1])
}
