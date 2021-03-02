package arbol

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type Nodo struct {
	Valor     int
	Izquierda *Nodo
	Derecha   *Nodo
}

type Arbol struct {
	raiz *Nodo
}

func NewArbol() *Arbol {
	return &Arbol{nil}
}

func NewNodo(valor int) *Nodo {
	return &Nodo{valor, nil, nil}
}

func insertar(raiz *Nodo, valor int) *Nodo {
	if raiz == nil {
		raiz = NewNodo(valor)
	} else if valor < raiz.Valor {
		NuevoIzq := insertar(raiz.Izquierda, valor)
		raiz.Izquierda = NuevoIzq
	} else if valor > raiz.Valor {
		NuevoDer := insertar(raiz.Derecha, valor)
		raiz.Derecha = NuevoDer
	}

	return raiz
}

func (this *Arbol) Insertar(valor int) {
	this.raiz = insertar(this.raiz, valor)
}

func (this *Arbol) GraficarArbol() {
	var cadena strings.Builder
	fmt.Fprintf(&cadena, "digraph G{\n")
	fmt.Fprintf(&cadena, "rankdir=UD\n")
	fmt.Fprintf(&cadena, "node[shape=circle]\n")
	fmt.Fprintf(&cadena, "concentrate=true\n")
	graficarArbol(this.raiz, &cadena)
	fmt.Fprintf(&cadena, "}")
	guardarArchivo(cadena.String())
}

func graficarArbol(raiz *Nodo, s *strings.Builder) {
	fmt.Fprintf(s, "nodo%p[label=\"%v\"];\n", &(*raiz), raiz.Valor)
	if raiz.Izquierda != nil {
		graficarArbol(raiz.Izquierda, s)
		fmt.Fprintf(s, "nodo%p->nodo%p;\n", &(*raiz), &(*raiz.Izquierda))
	}
	if raiz.Derecha != nil {
		graficarArbol(raiz.Derecha, s)
		fmt.Fprintf(s, "nodo%p->nodo%p;\n", &(*raiz), &(*raiz.Derecha))
	}
}

func (this *Arbol) ImplimirPreOrden() string {
	var s strings.Builder
	fmt.Fprintf(&s, strconv.Itoa(this.raiz.Valor)+", ")
	if this.raiz.Izquierda != nil {
		this.raiz.Izquierda.implimirHoja(&s)
	}
	if this.raiz.Derecha != nil {
		this.raiz.Derecha.implimirHoja(&s)
	}
	return s.String()
}

func (this *Arbol) ImplimirInorden() string {
	var s strings.Builder
	if this.raiz.Izquierda != nil {
		this.raiz.Izquierda.implimirHoja(&s)
	}
	fmt.Fprintf(&s, strconv.Itoa(this.raiz.Valor)+" ,")
	if this.raiz.Derecha != nil {
		this.raiz.Derecha.implimirHoja(&s)
	}
	return s.String()
}

func (this *Arbol) ImplimirPostorden() string {
	var s strings.Builder
	if this.raiz.Izquierda != nil {
		this.raiz.Izquierda.implimirHoja(&s)
	}
	if this.raiz.Derecha != nil {
		this.raiz.Derecha.implimirHoja(&s)
	}
	fmt.Fprintf(&s, strconv.Itoa(this.raiz.Valor))
	return s.String()
}

func (this *Nodo) implimirHoja(s *strings.Builder) {
	if this.Izquierda != nil {
		this.Izquierda.implimirHoja(s)
	}
	fmt.Fprintf(s, strconv.Itoa(this.Valor)+" ,")
	if this.Derecha != nil {
		this.Derecha.implimirHoja(s)
	}
}

func guardarArchivo(cadena string) {
	f, err := os.Create("ArbolBinario.dot")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString(cadena)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	fmt.Println(l, "Escritura Exitosa")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tpdf", "./"+"ArbolBinario.dot").Output()
	mode := int(0777)
	ioutil.WriteFile("ArbolBinario.pdf", cmd, os.FileMode(mode))
}
