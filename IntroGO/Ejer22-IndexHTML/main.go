package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	//http.HandleFunc("/login", home) responder a diferentes intrucciones.
	http.ListenAndServe(":3000", nil)
}

//            respuesta             request
func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html") //(response writer, *request, name)
}
