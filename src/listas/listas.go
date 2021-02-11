package listas

import (
	"fmt"
)

type Nodo struct {
	Mensajes []string
	Origen   string
	Destino  string
	Msg      string
	Sig      *Nodo
	Ant      *Nodo
}

type List struct {
	primero, ultimo *Nodo
}

func NewList() *List {
	return &List{nil, nil}
}

func (this *List) Insertar(nuevo *Nodo) {
	if this.primero == nil {
		this.primero = nuevo
		this.ultimo = nuevo
	} else {
		this.ultimo.Sig = nuevo
	}
}

func (this *Nodo) To_string() string {
	return "Mensajes: " + this.Origen 
}

func (this *List) To_string() string {
	var cadena string
	aux := this.primero
	for aux != nil {
		cadena += aux.To_string()
		aux = aux.Sig
	}
	return cadena
}

func (this *List) print() {
	fmt.Println("List")
	fmt.Println(this.To_string())
}
