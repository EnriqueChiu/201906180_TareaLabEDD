package main

import (
	"fmt"
	"math/rand"

	"./arbol"
)

func main() {
	var aleatorio int
	a := arbol.NewArbol()
	a.Insertar(19)
	for n := 0; n < 20; n++ {
		aleatorio = rand.Intn(30)
		a.Insertar(aleatorio)
	}
	a.Insertar(30)
	a.GraficarArbol()

	fmt.Println()
	fmt.Println("----------------------------------------Preorden---------------------------------------------------")
	fmt.Println(a.ImplimirPreOrden())
	fmt.Println("----------------------------------------------------------------------------------------------------")
	fmt.Println()
	fmt.Println("-----------------------------------------Inorden---------------------------------------------------")
	fmt.Println(a.ImplimirInorden())
	fmt.Println("----------------------------------------------------------------------------------------------------")
	fmt.Println()
	fmt.Println("-----------------------------------------Postorden--------------------------------------------------")
	fmt.Println(a.ImplimirPostorden())
	fmt.Println("----------------------------------------------------------------------------------------------------")
	fmt.Println()

}
