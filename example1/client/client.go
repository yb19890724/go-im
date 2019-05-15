package main

import (
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":2000", http.FileServer(http.Dir("./go-im/example1/client/"))))
}
