package main

import (
	"fmt"
	"strconv"
	"sync"
)

const MAX_ITEMS = 2

/**
 *	Funcion para agregar elementos al canal item
 *
 *	@param item canal donde agregar los elementos
 *	@param wg   puntero a una estructura WaitGroup para indicar cuando termina de ejecutarse
 */
func produce(items chan<- string, wg *sync.WaitGroup) {
	for i := 1; i <= MAX_ITEMS; i++ {
		items <- "item " + strconv.Itoa(i)
	}
	wg.Done()
}

/**
 *	Funcion para sacar elementos del canal items
 *
 *	@param items canal donde sacar elementos
 *	@param wg    puntero a una estructura WaitGroup para indicar cuando termina de ejecutarse
 */
func consume(items <-chan string) {
	for {
		item := <-items
		fmt.Println(item)
	}
}

func main() {
	wg := &sync.WaitGroup{}
	items := make(chan string)
	wg.Add(1)
	go produce(items, wg)
	go consume(items)
	wg.Wait()
}
