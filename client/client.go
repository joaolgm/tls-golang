package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	cont := 0
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs: readCertPool("server.crt"),
		},
	}

	client := &http.Client{Transport: tr}

	for cont < 5 {
		cont++
		time.Sleep(6 * time.Second)
		// server-service é o nome que escolhi para o meu service do kubernetes
		resp, err := client.Get("https://server-service.default.svc.cluster.local:5000/")
		if err != nil {
			fmt.Println("Client Error:", err)
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Client Error:", err)
			return
		}

		fmt.Println("Response:", string(body))

	}
}

func readCertPool(certFile string) *x509.CertPool {
	caCert, err := ioutil.ReadFile(certFile)
	if err != nil {
		panic(err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	return caCertPool
}
