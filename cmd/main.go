package main

import (
	"fmt"
	"icon/api"
	"net/http"
)

func main() {

	fmt.Println("http://localhost:8080")
	http.HandleFunc("/api/icon", api.IconHandler)
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.ListenAndServe(":8080", nil)
}
