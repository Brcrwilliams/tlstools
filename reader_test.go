package tlstools_test

import (
	"github.com/brcrwilliams/tlstools"
	"encoding/hex"
	"fmt"
	"os"
)

func ExampleReadPEM() {
	f, err := os.Open("/path/to/cert.pem")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	certs, err := tlstools.ReadPEM(f)
	if err != nil {
		panic(err)
	}

	for _, cert := range certs {
		fmt.Printf("Serial: %s\n", hex.EncodeToString(cert.SerialNumber.Bytes()))
	}
}

func ExampleReadDER() {
	f, err := os.Open("/path/to/cert.der")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	cert, err := tlstools.ReadDER(f)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Serial: %s\n", hex.EncodeToString(cert.SerialNumber.Bytes()))
}
