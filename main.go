package main

import (
        "fmt"
        "log"
        "net/http"
        "os"
	"net/http/httputil"
)

func main() {
        certFile := os.Getenv("CERT_FILE")
        keyFile := os.Getenv("KEY_FILE")

        if certFile == "" || keyFile == "" {
                log.Fatal("Missing certificate or key file. Please set the CERT_FILE and KEY_FILE environment variables.")
        }

        http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
                // Print all request headers
		reqDump, err := httputil.DumpRequest(req, true)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("REQUEST:\n%s", string(reqDump))
                // Send a simple response
                fmt.Fprintf(w, "Hello, HTTP/2!")
        })

        // Enable HTTP/2
        server := &http.Server{
                Addr:    ":8080",
                Handler: nil,
        }

        log.Fatal(server.ListenAndServeTLS(certFile, keyFile))
}
