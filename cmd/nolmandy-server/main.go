package main

import (
	"crypto/x509"
	"encoding/pem"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/aktsk/nolmandy/server"
	"github.com/aktsk/nolmandy/version"
)

const name = "nolmandy-server"

var GitCommit string

func main() {
	var (
		port         int
		certFileName string
		versionFlag  bool
	)

	flag.IntVar(&port, "port", 8000, "Port to listen")
	flag.StringVar(&certFileName, "certFile", "", "Certificate file")
	flag.BoolVar(&versionFlag, "version", false, "print version string")

	flag.Parse()

	if versionFlag {
		fmt.Printf("%s version: %s (rev: %s)", name, version.Get(), GitCommit)
		os.Exit(0)
	}

	var cert *x509.Certificate

	if certFileName != "" {
		certFile, err := os.Open(certFileName)
		if err != nil {
			log.Fatal(err)
		}

		defer certFile.Close()

		certPEM, err := ioutil.ReadAll(certFile)
		if err != nil {
			log.Fatal(err)
		}

		certDER, _ := pem.Decode(certPEM)
		cert, err = x509.ParseCertificate(certDER.Bytes)
		if err != nil {
			log.Fatal(err)
		}
	}

	server.Serve(port, cert)
}
