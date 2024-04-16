# 🛡️ Gots - Go TLS Library

Gots is a Go library for generating self-signed TLS certificates and managing TLS connections. It simplifies the process of handling TLS certificates, especially for development and testing purposes.

## ✨ Features

- **Certificate Generation:** Easily generate self-signed TLS certificates with a simple function call.
- **Key Pair Creation:** Create private key and public key pairs for secure communication.
- **Certificate Storage:** Save generated certificates and keys to specified directories.
- **Flexible Configuration:** Customize certificate and key file names and output directories as per your needs.
- **Error Handling:** Comprehensive error handling for smooth troubleshooting and debugging.


## 🔧 Usage

Here's an example of how you can utilize Gots to generate self-signed TLS certificates and set up an HTTPS server:

1. **Import Gots: Import Gots into your Go project:**

   ```bash
   import (
        "fmt"
        "net/http"

        "github.com/sprdgx/gots"
    )

2. **Generate TLS Certificate: Generate a self-signed TLS certificate and private key:**
   
   ```bash
    certPath := "certs/certificate.pem"
    keyPath := "certs/private_key.pem"
    rootCACertPath := "certs/root_ca.crt"

    err := gots.GenerateAndSaveCert(certPath, keyPath, rootCACertPath)
    if err != nil {
        fmt.Println("Error generating or saving certificate:", err)
        return
    }

    fmt.Println("Certificate and Private Key generation completed successfully.")
    fmt.Println("Root CA Certificate exported to:", rootCACertPath)
   
3. **Set Up HTTPS Server: Create an HTTPS server using the generated certificate and private key:**

   ```bash
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "Hi, welcome to HTTPS!")
    })

    fmt.Println("HTTPS Server running on port 9999")
    err := http.ListenAndServeTLS(":9999", certPath, keyPath, nil)
    if err != nil {
        fmt.Println("Error starting HTTPS server:", err)
    }

## 🤔 Why Use Gots? 

  **Simplicity:**
    Easily generate self-signed TLS certificates without complex configurations.

  **Convenience:**
    Simplify TLS certificate management for development and testing purposes.

  **Flexibility:** 
    Integrate TLS functionality seamlessly into Go applications.

## 🤝 Contributing

Contributions to Gots are welcome! If you encounter any issues or have suggestions for improvements, please open an issue or submit a pull request on [GitHub repository ](https://github.com/sprdgx/Gots).

## 📄 License

This project is licensed under the MIT License.


**Thank you for choosing Gots for your TLS certificate generation and management! If you have any questions, need assistance, or want to contribute to the project, feel free to reach out. Happy secure communication with Gots! 🌟**