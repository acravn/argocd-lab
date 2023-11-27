package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func getHome(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, fmt.Sprintf(os.Getenv("MESSAGE")))
}
func main() {
	http.HandleFunc("/", getHome)

	err := http.ListenAndServe(":3333", nil)
	log.Println(err)
}
