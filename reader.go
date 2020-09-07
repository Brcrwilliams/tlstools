package tlstools

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io"
	"io/ioutil"
)

// ReadPEM will read PEM-encoded certificates from reader
// and then parse them into *x509.Certificates.
// It will return an error if the input contains non-certificate
// PEMs, or if one of the PEMs is invalid.
func ReadPEM(reader io.Reader) ([]*x509.Certificate, error) {
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("could not read data: %w", err)
	}

	certs := []*x509.Certificate{}
	for len(data) > 0 {
		var block *pem.Block
		block, data = pem.Decode(data)
		if block == nil {
			break
		}

		if block.Type != "CERTIFICATE" {
			return nil, fmt.Errorf("file contains a PEM of unexpected type: %s", block.Type)
		}

		cert, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			return nil, fmt.Errorf("failed to parse PEM: %w", err)
		}

		certs = append(certs, cert)
	}

	return certs, nil
}

// ReadDER will a DER-encoded certificate from reader and parse
// it into an *x509.Certificate. It expects the input to contain
// only one certificate, since DER does not have delimiters.
func ReadDER(reader io.Reader) (*x509.Certificate, error) {
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("could not read data: %w", err)
	}
	return x509.ParseCertificate(data)
}
