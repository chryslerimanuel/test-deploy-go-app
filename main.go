package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println("server started")

	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}

	//root
	http.HandleFunc("/", handlerHello)

	http.HandleFunc("/berhasil", handlerSuccess)
	http.HandleFunc("/sukses", handlerSuccess)

	http.ListenAndServe(":"+port, nil)
}

func handlerHello(w http.ResponseWriter, r *http.Request) {
	var message = "Hello world!"
	w.Write([]byte(message))
}

func handlerSuccess(w http.ResponseWriter, r *http.Request) {
	var message = "Hore Ini update terbaru!!"
	w.Write([]byte(message))
}
