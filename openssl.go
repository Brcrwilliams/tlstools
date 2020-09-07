package tlstools

import (
	"crypto/dsa"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"strings"
	"time"
)

// OpenSSLFormat is a transformation of x509.Certificate.
// It's intended to output JSON which looks similar to the output
// of `openssl x509 -text`.
type OpenSSLFormat struct {
	Version                 int
	SerialNumber            string
	Issuer                  string
	Subject                 string
	Validity                *Validity
	SubjectPublicKeyInfo    *SubjectPublicKeyInfo
	X509v3Extensions        *X509v3Extensions
	AuthorityInformation    *AuthorityInformation
	SubjectAlternativeNames []string
	CertificatePolicies     []string
	CRLDistributionPoints   []string
	SignatureAlgorithm      string
	Signature               string
}

// CertToOpenSSL converts an *x509.Certificate into OpenSSLFormat.
func CertToOpenSSL(cert *x509.Certificate) *OpenSSLFormat {
	policies := []string{}
	for _, policy := range cert.PolicyIdentifiers {
		policies = append(policies, policy.String())
	}

	return &OpenSSLFormat{
		Version:                 cert.Version,
		SerialNumber:            encodeColonSeparatedHex(cert.SerialNumber.Bytes()),
		Issuer:                  cert.Issuer.String(),
		Subject:                 cert.Subject.String(),
		Validity:                newValidity(cert),
		SubjectPublicKeyInfo:    newSubjectPublicKeyInfo(cert),
		X509v3Extensions:        newX509v3Extensions(cert),
		AuthorityInformation:    newAuthorityInformation(cert),
		SubjectAlternativeNames: getAltNames(cert),
		CertificatePolicies:     policies,
		CRLDistributionPoints:   cert.CRLDistributionPoints,
		SignatureAlgorithm:      cert.SignatureAlgorithm.String(),
		Signature:               encodeColonSeparatedHex(cert.Signature),
	}
}

func getAltNames(cert *x509.Certificate) []string {
	names := []string{}
	for _, name := range cert.DNSNames {
		names = append(names, "DNS:"+name)
	}

	for _, name := range cert.EmailAddresses {
		names = append(names, "Email:"+name)
	}

	for _, name := range cert.IPAddresses {
		names = append(names, "IP:"+name.String())
	}

	for _, name := range cert.URIs {
		names = append(names, "URI:"+name.String())
	}

	return names
}

// Validity contains the NotBefore and NotAfter timestamps.
type Validity struct {
	NotBefore time.Time
	NotAfter  time.Time
}

func newValidity(cert *x509.Certificate) *Validity {
	return &Validity{
		NotBefore: cert.NotBefore,
		NotAfter:  cert.NotAfter,
	}
}

const timeFormat = "Jan 2 15:04:05 2006 MST"

// MarshalJSON turns Valdity into a JSON object, with
// the timestamps in "Jan 2 15:04:05 2006 MST" format.
func (v *Validity) MarshalJSON() ([]byte, error) {
	j := map[string]string{
		"NotBefore": v.NotBefore.Format(timeFormat),
		"NotAfter":  v.NotAfter.Format(timeFormat),
	}
	return json.Marshal(j)
}

// SubjectPublicKeyInfo contains information about the certificate public key.
type SubjectPublicKeyInfo struct {
	PublicKeyAlgorithm string
	Parameters         interface{}
}

func newSubjectPublicKeyInfo(cert *x509.Certificate) *SubjectPublicKeyInfo {
	s := new(SubjectPublicKeyInfo)
	switch k := cert.PublicKey.(type) {
	case *rsa.PublicKey:
		s.PublicKeyAlgorithm = "RSA"
		s.Parameters = newRSAPublicKeyInfo(k)
	case *dsa.PublicKey:
		s.PublicKeyAlgorithm = "DSA"
		s.Parameters = newDSAPublicKeyInfo(k)
	case *ecdsa.PublicKey:
		s.PublicKeyAlgorithm = "ECDSA"
		s.Parameters = newECDSAPublicKeyInfo(k)
	case ed25519.PublicKey:
		s.PublicKeyAlgorithm = "Ed25519"
		s.Parameters = newEd25519PublicKeyInfo(k)
	default:
		s.PublicKeyAlgorithm = "Unknown"
	}
	return s
}

type rsaPublicKeyInfo struct {
	KeySize  int
	Modulus  string
	Exponent int
}

func newRSAPublicKeyInfo(key *rsa.PublicKey) *rsaPublicKeyInfo {
	return &rsaPublicKeyInfo{
		KeySize:  key.Size() * 8,
		Modulus:  encodeColonSeparatedHex(key.N.Bytes()),
		Exponent: key.E,
	}
}

type dsaPublicKeyInfo struct {
	P string
	Q string
	G string
	Y string
}

func newDSAPublicKeyInfo(key *dsa.PublicKey) *dsaPublicKeyInfo {
	return &dsaPublicKeyInfo{
		P: encodeColonSeparatedHex(key.P.Bytes()),
		Q: encodeColonSeparatedHex(key.Q.Bytes()),
		G: encodeColonSeparatedHex(key.G.Bytes()),
		Y: encodeColonSeparatedHex(key.Y.Bytes()),
	}
}

type ecdsaPublicKeyInfo struct {
	Curve string
	X     string
	Y     string
}

func newECDSAPublicKeyInfo(key *ecdsa.PublicKey) *ecdsaPublicKeyInfo {
	return &ecdsaPublicKeyInfo{
		Curve: key.Params().Name,
		X:     encodeColonSeparatedHex(key.X.Bytes()),
		Y:     encodeColonSeparatedHex(key.Y.Bytes()),
	}
}

type ed25519PublicKeyInfo struct {
	PublicKey string
}

func newEd25519PublicKeyInfo(key ed25519.PublicKey) *ed25519PublicKeyInfo {
	return &ed25519PublicKeyInfo{
		PublicKey: encodeColonSeparatedHex([]byte(key)),
	}
}

// TODO: Golang ignores whether or not these extensions are critical.
// Do we care enough to go and parse that out of the ASN.1 data?
// RFC 5280 says the following:
// Key Usage - CAs must include, extension should be critical
// Basic Constraints - CAs must include, extension must be critical
// Extended Key Usage - Extension may be critical or non-critical

// X509v3Extensions contains the x509 v3 Extenions
type X509v3Extensions struct {
	KeyUsage               []string
	ExtendedKeyUsage       []string
	BasicConstraints       *BasicConstraints
	SubjectKeyIdentifier   string
	AuthorityKeyIdentifier string
}

func newX509v3Extensions(cert *x509.Certificate) *X509v3Extensions {
	extendedUsage := []string{}
	for _, use := range cert.ExtKeyUsage {
		extendedUsage = append(extendedUsage, extendedKeyUsageToString(use))
	}

	return &X509v3Extensions{
		KeyUsage:               keyUsageToStrings(cert.KeyUsage),
		ExtendedKeyUsage:       extendedUsage,
		BasicConstraints:       newBasicConstraints(cert),
		SubjectKeyIdentifier:   encodeColonSeparatedHex(cert.SubjectKeyId),
		AuthorityKeyIdentifier: encodeColonSeparatedHex(cert.AuthorityKeyId),
	}
}

// TODO: The output of this seems wrong. How to fix?
func keyUsageToStrings(use x509.KeyUsage) []string {
	usages := []string{}
	if (use & x509.KeyUsageCRLSign) == 0 {
		usages = append(usages, "CRL Sign")
	}

	if (use & x509.KeyUsageCertSign) == 0 {
		usages = append(usages, "Cert Sign")
	}

	if (use & x509.KeyUsageContentCommitment) == 0 {
		usages = append(usages, "Content Commitment")
	}

	if (use & x509.KeyUsageDataEncipherment) == 0 {
		usages = append(usages, "Data Encipherment")
	}

	if (use & x509.KeyUsageDecipherOnly) == 0 {
		usages = append(usages, "Decipher Only")
	}

	if (use & x509.KeyUsageDigitalSignature) == 0 {
		usages = append(usages, "Digital Signature")
	}

	if (use & x509.KeyUsageEncipherOnly) == 0 {
		usages = append(usages, "Encipher Only")
	}

	if (use & x509.KeyUsageKeyAgreement) == 0 {
		usages = append(usages, "Key Agreement")
	}

	if (use & x509.KeyUsageKeyEncipherment) == 0 {
		usages = append(usages, "Key Encipherment")
	}

	return usages
}

func extendedKeyUsageToString(use x509.ExtKeyUsage) string {
	switch use {
	case x509.ExtKeyUsageAny:
		return "Any"
	case x509.ExtKeyUsageClientAuth:
		return "Client Auth"
	case x509.ExtKeyUsageCodeSigning:
		return "Code Signing"
	case x509.ExtKeyUsageEmailProtection:
		return "Email Protection"
	case x509.ExtKeyUsageIPSECEndSystem:
		return "IPSEC End System"
	case x509.ExtKeyUsageIPSECTunnel:
		return "IPSEC Tunnel"
	case x509.ExtKeyUsageIPSECUser:
		return "IPSEC User"
	case x509.ExtKeyUsageMicrosoftCommercialCodeSigning:
		return "Microsoft Commercial Code Signing"
	case x509.ExtKeyUsageMicrosoftKernelCodeSigning:
		return "Microsoft Kernel Code Signing"
	case x509.ExtKeyUsageMicrosoftServerGatedCrypto:
		return "Microsoft Server Gated Crypto"
	case x509.ExtKeyUsageNetscapeServerGatedCrypto:
		return "Netscape Server Gated Crypto"
	case x509.ExtKeyUsageOCSPSigning:
		return "OCSP Signing"
	case x509.ExtKeyUsageServerAuth:
		return "Server Auth"
	case x509.ExtKeyUsageTimeStamping:
		return "Time Stamping"
	}
	return "Unknown"
}

// BasicConstraints represents the Basic Constraints extension.
type BasicConstraints struct {
	CA                bool
	MaxPathLength     int
	maxPathLengthZero bool
}

func newBasicConstraints(cert *x509.Certificate) *BasicConstraints {
	return &BasicConstraints{
		CA:                cert.IsCA,
		MaxPathLength:     cert.MaxPathLen,
		maxPathLengthZero: cert.MaxPathLenZero,
	}
}

// MaxPathIsNil is used to determine if MaxPathLength is zero or unset.
func (b *BasicConstraints) MaxPathIsNil() bool {
	return !b.maxPathLengthZero && b.MaxPathLength == 0
}

// MarshalJSON converts the BasicConstraints into JSON.
// If MatxPathLength is nil, then it will be ommitted.
func (b *BasicConstraints) MarshalJSON() ([]byte, error) {
	j := map[string]interface{}{}
	j["CA"] = b.CA
	if !b.MaxPathIsNil() {
		j["MaxPathLength"] = b.MaxPathLength
	}
	return json.Marshal(j)
}

// AuthorityInformation is...
// only grouped here because that's what openssl does.
type AuthorityInformation struct {
	OCSP      []string
	CAIssuers []string
}

func newAuthorityInformation(cert *x509.Certificate) *AuthorityInformation {
	return &AuthorityInformation{
		OCSP:      cert.OCSPServer,
		CAIssuers: cert.IssuingCertificateURL,
	}
}

func encodeColonSeparatedHex(b []byte) string {
	s := hex.EncodeToString(b)
	buff := new(strings.Builder)
	l := len(s)
	for i := 0; i < l-2; i = i + 2 {
		buff.WriteString(s[i : i+2])
		buff.WriteString(":")
	}
	buff.WriteString(s[l-2:])
	return buff.String()
}
