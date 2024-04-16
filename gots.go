package gots

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"path/filepath"
	"time"
)

// GenerateAndSaveCert generates a self-signed certificate and private key and saves them to files.
func GenerateAndSaveCert(certPath, keyPath string) error {
	// Generate certificate and private key
	certPEM, privateKeyPEM, err := generateCert()
	if err != nil {
		return err
	}

	// Create the directory if it doesn't exist
	dir := filepath.Dir(certPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("Error creating directory: %v", err)
	}

	// Save certificate to file
	if err := ioutil.WriteFile(certPath, certPEM, 0644); err != nil {
		return fmt.Errorf("Error saving certificate: %v", err)
	}
	fmt.Printf("Certificate saved to: %s\n", certPath)

	// Save private key to file
	if err := ioutil.WriteFile(keyPath, privateKeyPEM, 0644); err != nil {
		return fmt.Errorf("Error saving private key: %v", err)
	}
	fmt.Printf("Private Key saved to: %s\n", keyPath)

	return nil
}

func generateCert() ([]byte, []byte, error) {
	// Generate RSA private key
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}

	// Create certificate template
	template := x509.Certificate{
		SerialNumber: big.NewInt(time.Now().Unix()),
		Subject: pkix.Name{
			CommonName:         "localhost",
			Organization:       []string{"SPRD"}, // Set Organization (O) name to "SPRD"
			OrganizationalUnit: []string{"SPRD"}, // Set Organizational Unit (OU) to placeholder
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(1, 0, 0), // Valid for 1 year
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		IsCA:                  true, // Certificate is a CA certificate
	}

	// Generate certificate
	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &privateKey.PublicKey, privateKey)
	if err != nil {
		return nil, nil, err
	}

	// Encode certificate to PEM format
	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: derBytes})

	// Encode private key to PEM format
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)})

	return certPEM, privateKeyPEM, nil
}
