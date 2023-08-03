package main

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
)

func main() {
	cert, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		fmt.Println("Error loading certificates:", err)
		return
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	server := &http.Server{
		Addr:      ":5000",
		TLSConfig: tlsConfig,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, TLS!")
	})

	l, err := net.Listen("tcp", server.Addr)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer l.Close()

	fmt.Println("Server listening on", server.Addr)
	err = server.Serve(tls.NewListener(l, server.TLSConfig))
	if err != nil {
		fmt.Println("Error serving:", err)
	}
}
