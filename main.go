package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("server started")
	//root
	http.HandleFunc("/", handlerHello)

	http.HandleFunc("/berhasil", handlerSuccess)

	http.ListenAndServe("localhost:9001", nil)
}

func handlerHello(w http.ResponseWriter, r *http.Request) {
	var message = "Hello world!"
	w.Write([]byte(message))
}

func handlerSuccess(w http.ResponseWriter, r *http.Request) {
	var message = "Hore!!"
	w.Write([]byte(message))
}
