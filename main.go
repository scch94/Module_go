package main

import (
	"fmt"

	"example.com/go-modelos/slices"
)

func main() {
	fmt.Println("hello gophers")
	lista := []uint{1, 2, 3}
	fmt.Println(slices.Includes(lista, 4))
}
