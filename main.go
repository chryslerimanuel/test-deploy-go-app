package main

import (
	"fmt"
	"os"
	"net/http"
)

func main() {
	fmt.Println("server started")

	// for _, env := range os.Environ() {
	// 	fmt.Println(env)
	// }

	// fmt.Println(os.Getenv("HOMEPATH"))

	port := os.Getenv("PORT")
	if port == "" {
    port = "9000" // Default port if not specified
	}

	// //root
	http.HandleFunc("/", handlerHello)

	http.HandleFunc("/berhasil", handlerSuccess)

	http.ListenAndServe(":"+port, nil)
}

func handlerHello(w http.ResponseWriter, r *http.Request) {
	var message = "Hello world!"
	w.Write([]byte(message))
}

func handlerSuccess(w http.ResponseWriter, r *http.Request) {
	var message = "Hore!!"
	w.Write([]byte(message))
}
