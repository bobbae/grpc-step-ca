package main

import (
	"log"
	"net/http"
)

func HiHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Hello, world!\n"))
}
func main() {
	http.HandleFunc("/hi", HiHandler)
	err := http.ListenAndServeTLS(":9443", "srv.crt", "srv.key", nil)
	if err != nil {
		log.Fatal(err)
	}
}
