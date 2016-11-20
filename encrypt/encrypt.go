package main

import (
	"log"
	"net/http"
	"rsc.io/letsencrypt"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, HTTP/2 TLS world!"))
	})
	var m letsencrypt.Manager
	if err := m.CacheFile("letsencrypt.cache"); err != nil {
		log.Fatal(err)
	} 
	log.Fatal(m.Serve())
}
