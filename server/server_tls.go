package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	// Server configuration
	serverConfig := &tls.Config{
		Certificates: loadServerCertificates(),       // Load the server's certificate and private key
		ClientAuth:   tls.RequireAndVerifyClientCert, // Enable client authentication
		ClientCAs:    loadClientCACertificate(),      // Load client CA certificate for client certificate validation
	}

	// Start the server
	listener, err := tls.Listen("tcp", ":5000", serverConfig)
	if err != nil {
		log.Fatalf("Error creating listener: %s", err)
	}
	defer listener.Close()
	fmt.Println("Server started. Listening on port 5000...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %s", err)
			continue
		}

		// Handle each incoming connection in a goroutine
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Example of server-side data exchange
	msg := "Hello, client! This is a secure connection."
	_, err := conn.Write([]byte(msg))
	if err != nil {
		log.Printf("Error writing to connection: %s", err)
	}
}

// Function to load server certificate and private key
func loadServerCertificates() []tls.Certificate {
	cert, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		log.Fatalf("Error loading server certificate and key: %s", err)
	}
	return []tls.Certificate{cert}
}

// Function to load client CA certificate for client certificate validation
func loadClientCACertificate() *x509.CertPool {
	caCert, err := os.ReadFile("ca.crt")
	if err != nil {
		log.Fatalf("Error loading client CA certificate: %s", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	return caCertPool
}
