package main

//go get github.com/githubnemo/CompileDaemon

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"./listas"

	"github.com/gorilla/mux"
)

type Men struct {
	Mensajes []Mensajes `json:"Mensajes"`
}

type Mensajes struct {
	Origen  string `json:"Origen"`
	Destino string `json:"Destino"`
	Msg     []Msg  `json: "json:"Msg`
}

type Msg struct {
	Fecha string `json:"Fecha"`
	Texto string `json:"Texto"`
}

func cargar(w http.ResponseWriter, r *http.Request) {
	var ms Men
	var mss listas.Nodo
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error de entrada")
	}
	json.Unmarshal(reqBody, &ms)
	fmt.Println(mss.To_string())
	json.NewEncoder(w).Encode(ms)

}

func ver(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "hola")
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/cargar", cargar).Methods("POST")
	router.HandleFunc("/ver", ver).Methods("GET")
	log.Fatal(http.ListenAndServe(":3000", router))

}
