package main

import (
	"fmt"
	"net/http"

	"github.com/sprdgx/gots"
)

func main() {
	// Define paths for certificate and key files
	certPath := "certs/certificate.pem"
	keyPath := "certs/private_key.pem"

	// Generate and save certificate and private key using gots library
	if err := gots.GenerateAndSaveCert(certPath, keyPath); err != nil {
		fmt.Println("Error generating or saving certificate:", err)
		return
	}

	fmt.Println("Certificate and Private Key generation completed successfully.")

	// Start HTTPS server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hi, welcome to HTTPS!")
	})

	fmt.Println("HTTPS Server running on port 9999")
	err := http.ListenAndServeTLS(":9999", certPath, keyPath, nil)
	if err != nil {
		fmt.Println("Error starting HTTPS server:", err)
	}
}
