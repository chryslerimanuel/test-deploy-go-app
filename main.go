package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("server started")

	// for _, env := range os.Environ() {
	// 	fmt.Println(env)
	// }

	fmt.Println(os.Getenv("HOMEPATH"))
	// //root
	// http.HandleFunc("/", handlerHello)

	// http.HandleFunc("/berhasil", handlerSuccess)

	// http.ListenAndServe("localhost:9091", nil)
}

// func handlerHello(w http.ResponseWriter, r *http.Request) {
// 	var message = "Hello world!"
// 	w.Write([]byte(message))
// }

// func handlerSuccess(w http.ResponseWriter, r *http.Request) {
// 	var message = "Hore!!"
// 	w.Write([]byte(message))
// }
