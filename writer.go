package certreader

import (
	"crypto/x509"
	"encoding/json"
	"io"
)

func prettyEncoder(out io.Writer) *json.Encoder {
	e := json.NewEncoder(out)
	e.SetIndent("", "  ")
	return e
}

// WriteCert takes an *x509.Certificate and
// writes it to out in OpenSSL JSON format.
func WriteCert(out io.Writer, cert *x509.Certificate) error {
	return prettyEncoder(out).Encode(CertToOpenSSL(cert))
}

// WriteCerts takes a slice of *x509.Certificates and
// writes them to out in OpenSSL JSON format.
func WriteCerts(out io.Writer, certs []*x509.Certificate) error {
	all := []*OpenSSLFormat{}
	for _, cert := range certs {
		all = append(all, CertToOpenSSL(cert))
	}
	return prettyEncoder(out).Encode(&all)
}
