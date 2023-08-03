package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"os"
)

func main() {
	// Client configuration
	clientConfig := &tls.Config{
		RootCAs:      loadServerCACertificate(), // Load server CA certificate for server certificate validation
		Certificates: loadClientCertificates(),  // Load the client's certificate and private key
	}

	conn, err := tls.Dial("tcp", "server-service.default.svc.cluster.local:5000", clientConfig)
	if err != nil {
		log.Fatalf("Error connecting to server: %s", err)
	}
	defer conn.Close()

	// Example of client-side data exchange
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Printf("Error reading from connection: %s", err)
		return
	}

	fmt.Println("Server says:", string(buf[:n]))
}

// Function to load server CA certificate for server certificate validation
func loadServerCACertificate() *x509.CertPool {
	caCert, err := os.ReadFile("ca.crt")
	if err != nil {
		log.Fatalf("Error loading server CA certificate: %s", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	return caCertPool
}

// Function to load client certificate and private key
func loadClientCertificates() []tls.Certificate {
	cert, err := tls.LoadX509KeyPair("client.crt", "client.key")
	if err != nil {
		log.Fatalf("Error loading client certificate and key: %s", err)
	}
	return []tls.Certificate{cert}
}
