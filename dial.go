package tlstools

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net"
	"os"
	"time"
)

var dialer = &net.Dialer{
	Timeout: 3 * time.Second,
}

var conf = &tls.Config{
	InsecureSkipVerify: true,
	VerifyConnection:   warnOnVerificationFailure,
}

// Dial opens a TLS connection to the given address over TCP,
// and returns the peer certificates. It will return an error
// if there was an error opening the TLS connection.
func Dial(addr string) ([]*x509.Certificate, error) {
	conn, err := tls.DialWithDialer(dialer, "tcp", addr, conf)
	if err != nil {
		return nil, fmt.Errorf("connection error: %w", err)
	}
	defer conn.Close()

	return conn.ConnectionState().PeerCertificates, nil
}

// warnOnVerificationFailure re-implements the regular certificate verification process,
// but never returns an error. It is used so that we may continue processing invalid certificates,
// but warn the user that the chain is not valid.
func warnOnVerificationFailure(cs tls.ConnectionState) error {
	opts := x509.VerifyOptions{
		DNSName:       cs.ServerName,
		Intermediates: x509.NewCertPool(),
	}
	for _, cert := range cs.PeerCertificates[1:] {
		opts.Intermediates.AddCert(cert)
	}
	_, err := cs.PeerCertificates[0].Verify(opts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "WARNING: certificate verification failed: %s\n", err.Error())
	}
	return nil
}
