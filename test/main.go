package main

import (
	"fmt"

	"github.com/sprdgx/gots" // Import your package here
)

func main() {
	// Define paths for certificate and key files
	certPath := "certs/certificate.pem"
	keyPath := "certs/private_key.pem"

	// Generate and save certificate and private key using your package
	if err := gots.GenerateAndSaveCert(certPath, keyPath); err != nil {
		fmt.Println("Error generating or saving certificate:", err)
		return
	}

	fmt.Println("Certificate and Private Key generation completed successfully.")
}
