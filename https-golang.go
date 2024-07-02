package main

import (
	"log"
	"net/http"
)

const (
	domainAndPort = "domain-name:443"
	serverCert    = "path/to/server.crt"
	serverKey     = "path/to/server.key"
)

func main() {
	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Test HTTPS server in Golang using net/http package\n"))
	})

	log.Fatal(http.ListenAndServeTLS(domainAndPort, serverCert, serverKey, nil))
}
