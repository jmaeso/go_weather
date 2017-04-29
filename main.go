package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("`curl http://localhost:8080` --> Hello from server")
	fmt.Println("Ctrl+c to stop the server")

	http.HandleFunc("/", helloFunc) //Call to root path calls to helloFunc
	http.ListenAndServe(":8080", nil)
}

func helloFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("You asked for nothing! Here you have a Hello World so. :D\n"))
}
