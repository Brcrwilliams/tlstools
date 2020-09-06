package tlstools

import (
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"io"
)

func prettyEncoder(out io.Writer) *json.Encoder {
	e := json.NewEncoder(out)
	e.SetIndent("", "  ")
	return e
}

// WriteX509Meta takes an *x509.Certificate and
// writes the x509 metadata to out in OpenSSL JSON format.
func WriteX509Meta(out io.Writer, cert *x509.Certificate) error {
	return prettyEncoder(out).Encode(CertToOpenSSL(cert))
}

// WriteX509Metas takes a slice of *x509.Certificates and
// writes the x509 metadata to out in OpenSSL JSON format.
func WriteX509Metas(out io.Writer, certs []*x509.Certificate) error {
	all := []*OpenSSLFormat{}
	for _, cert := range certs {
		all = append(all, CertToOpenSSL(cert))
	}
	return prettyEncoder(out).Encode(&all)
}

// WritePEM encodes the given certificate to PEM
// and writes it to out.
func WritePEM(out io.Writer, cert *x509.Certificate) error {
	block := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: cert.Raw,
	}
	return pem.Encode(out, block)
}
