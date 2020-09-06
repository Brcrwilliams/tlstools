package certreader

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io"
	"io/ioutil"
)

func Dial(addr string) ([]*x509.Certificate, error) {
	conn, err := tls.Dial("tcp", addr, &tls.Config{InsecureSkipVerify: true})
	if err != nil {
		return nil, fmt.Errorf("connection error: %w", err)
	}

	return conn.ConnectionState().PeerCertificates, nil
}

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

func ReadDER(reader io.Reader) (*x509.Certificate, error) {
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("could not read data: %w", err)
	}
	return x509.ParseCertificate(data)
}
