# TLS Tools for Humans

[![PkgGoDev](https://pkg.go.dev/badge/github.com/brcrwilliams/tlstools)](https://pkg.go.dev/github.com/brcrwilliams/tlstools)

This repository provides some command line tools and libraries that are intended
to replace some of the common certificate-viewing commands from openssl, such as `openssl s_client` and `openssl x509`.

These tools are _not_ intended to perform security functions,
such as providing cryptographic primitives or APIs for use in TLS servers.

## Installation

You can install all of the tools using:

```
go get github.com/brcrwilliams/tlstools/cmd/...
```

## readpem

`readpem` is a tool to retrieve all of the peer certificate PEMs from a remote address.
It takes the address in the form of host:port as a single positional argument.
If no port is given, it will default to port 443.
It will output the certificates to stdout, and can be piped as needed.
Ex: `readpem example.com:443 > chain.pem`

Use x509meta if you want to see the x509 metadata.
Ex: `x509meta --pem chain.pem`

`readpem` will perform certificate verification when reading certificates using the golang standard library (`(*x509.Certificate).Verify()`).
If certificate verification fails, it will emit a warning, but continue operating as normal.
It does not check certificate revocation.

### Example usage

```
$ readpem github.com
-----BEGIN CERTIFICATE-----
MIIG1TCCBb2gAwIBAgIQBVfICygmg6F7ChFEkylreTANBgkqhkiG9w0BAQsFADBw
MQswCQYDVQQGEwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMRkwFwYDVQQLExB3
d3cuZGlnaWNlcnQuY29tMS8wLQYDVQQDEyZEaWdpQ2VydCBTSEEyIEhpZ2ggQXNz
dXJhbmNlIFNlcnZlciBDQTAeFw0yMDA1MDUwMDAwMDBaFw0yMjA1MTAxMjAwMDBa
MGYxCzAJBgNVBAYTAlVTMRMwEQYDVQQIEwpDYWxpZm9ybmlhMRYwFAYDVQQHEw1T
YW4gRnJhbmNpc2NvMRUwEwYDVQQKEwxHaXRIdWIsIEluYy4xEzARBgNVBAMTCmdp
dGh1Yi5jb20wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQC7MrTQ2J6a
nox5KUwrqO9cQ9STO5R4/zBUxxvI5S8bmc0QjWfIVAwHWuT0Bn/H1oS0LM0tTkQm
ARrqN77v9McVB8MWTGsmGQnS/1kQRFuKiYGUHf7iX5pfijbYsOkfb4AiVKysKUNV
UtgVvpJoe5RWURjQp9XDWkeo2DzGHXLcBDadrM8VLC6H1/D9SXdVruxKqduLKR41
Z/6dlSDdeY1gCnhz3Ch1pYbfMfsTCTamw+AtRtwlK3b2rfTHffhowjuzM15UKt+b
rr/cEBlAjQTva8rutYU9K9ONgl+pG2u7Bv516DwmNy8xz9wOjTeOpeh0M9N/ewq8
cgbR87LFaxi1AgMBAAGjggNzMIIDbzAfBgNVHSMEGDAWgBRRaP+QrwIHdTzM2WVk
YqISuFlyOzAdBgNVHQ4EFgQUYwLSXQJf943VWhKedhE2loYsikgwJQYDVR0RBB4w
HIIKZ2l0aHViLmNvbYIOd3d3LmdpdGh1Yi5jb20wDgYDVR0PAQH/BAQDAgWgMB0G
A1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjB1BgNVHR8EbjBsMDSgMqAwhi5o
dHRwOi8vY3JsMy5kaWdpY2VydC5jb20vc2hhMi1oYS1zZXJ2ZXItZzYuY3JsMDSg
MqAwhi5odHRwOi8vY3JsNC5kaWdpY2VydC5jb20vc2hhMi1oYS1zZXJ2ZXItZzYu
Y3JsMEwGA1UdIARFMEMwNwYJYIZIAYb9bAEBMCowKAYIKwYBBQUHAgEWHGh0dHBz
Oi8vd3d3LmRpZ2ljZXJ0LmNvbS9DUFMwCAYGZ4EMAQICMIGDBggrBgEFBQcBAQR3
MHUwJAYIKwYBBQUHMAGGGGh0dHA6Ly9vY3NwLmRpZ2ljZXJ0LmNvbTBNBggrBgEF
BQcwAoZBaHR0cDovL2NhY2VydHMuZGlnaWNlcnQuY29tL0RpZ2lDZXJ0U0hBMkhp
Z2hBc3N1cmFuY2VTZXJ2ZXJDQS5jcnQwDAYDVR0TAQH/BAIwADCCAXwGCisGAQQB
1nkCBAIEggFsBIIBaAFmAHUAKXm+8J45OSHwVnOfY6V35b5XfZxgCvj5TV0mXCVd
x4QAAAFx5ltprwAABAMARjBEAiAuWGCWxN/M0Ms3KOsqFjDMHT8Aq0SlHfQ68KDg
rVU6AAIgDA+2EB0D5W5r0i4Nhljx6ABlIByzrEdfcxiOD/o6//EAdQAiRUUHWVUk
VpY/oS/x922G4CMmY63AS39dxoNcbuIPAgAAAXHmW2nTAAAEAwBGMEQCIBp+XQKa
UDiPHwjBxdv5qvgyALKaysKqMF60gqem8iPRAiAk9Dp5+VBUXfSHqyW+tVShUigh
ndopccf8Gs21KJ4jXgB2AFGjsPX9AXmcVm24N3iPDKR6zBsny/eeiEKaDf7UiwXl
AAABceZbahsAAAQDAEcwRQIgd/5HcxT4wfNV8zavwxjYkw2TYBAuRCcqp1SjWKFn
4EoCIQDHSTHxnbpxWFbP6v5Y6nGFZCDjaHgd9HrzUv2J/DaacDANBgkqhkiG9w0B
AQsFAAOCAQEAhjKPnBW4r+jR3gg6RA5xICTW/A5YMcyqtK0c1QzFr8S7/l+skGpC
yCHrJfFrLDeyKqgabvLRT6YvvM862MGfMMDsk+sKWtzLbDIcYG7sbviGpU+gtG1q
B0ohWNApfWWKyNpquqvwdSEzAEBvhcUT5idzbK7q45bQU9vBIWgQz+PYULAU7KmY
z7jOYV09o22TNMQT+hFmo92+EBlwSeIETYEsHy5ZxixTRTvu9hP00CyEbiht5OTK
5EiJG6vsIh/uEtRsdenMCxV06W2f20Af4iSFo0uk6c1ryHefh08FcwA4pSNUaPyi
Pb8YGQ6o/blejFzo/OSiUnDueafSJ0p6SQ==
-----END CERTIFICATE-----
-----BEGIN CERTIFICATE-----
MIIEsTCCA5mgAwIBAgIQBOHnpNxc8vNtwCtCuF0VnzANBgkqhkiG9w0BAQsFADBs
MQswCQYDVQQGEwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMRkwFwYDVQQLExB3
d3cuZGlnaWNlcnQuY29tMSswKQYDVQQDEyJEaWdpQ2VydCBIaWdoIEFzc3VyYW5j
ZSBFViBSb290IENBMB4XDTEzMTAyMjEyMDAwMFoXDTI4MTAyMjEyMDAwMFowcDEL
MAkGA1UEBhMCVVMxFTATBgNVBAoTDERpZ2lDZXJ0IEluYzEZMBcGA1UECxMQd3d3
LmRpZ2ljZXJ0LmNvbTEvMC0GA1UEAxMmRGlnaUNlcnQgU0hBMiBIaWdoIEFzc3Vy
YW5jZSBTZXJ2ZXIgQ0EwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQC2
4C/CJAbIbQRf1+8KZAayfSImZRauQkCbztyfn3YHPsMwVYcZuU+UDlqUH1VWtMIC
Kq/QmO4LQNfE0DtyyBSe75CxEamu0si4QzrZCwvV1ZX1QK/IHe1NnF9Xt4ZQaJn1
itrSxwUfqJfJ3KSxgoQtxq2lnMcZgqaFD15EWCo3j/018QsIJzJa9buLnqS9UdAn
4t07QjOjBSjEuyjMmqwrIw14xnvmXnG3Sj4I+4G3FhahnSMSTeXXkgisdaScus0X
sh5ENWV/UyU50RwKmmMbGZJ0aAo3wsJSSMs5WqK24V3B3aAguCGikyZvFEohQcft
bZvySC/zA/WiaJJTL17jAgMBAAGjggFJMIIBRTASBgNVHRMBAf8ECDAGAQH/AgEA
MA4GA1UdDwEB/wQEAwIBhjAdBgNVHSUEFjAUBggrBgEFBQcDAQYIKwYBBQUHAwIw
NAYIKwYBBQUHAQEEKDAmMCQGCCsGAQUFBzABhhhodHRwOi8vb2NzcC5kaWdpY2Vy
dC5jb20wSwYDVR0fBEQwQjBAoD6gPIY6aHR0cDovL2NybDQuZGlnaWNlcnQuY29t
L0RpZ2lDZXJ0SGlnaEFzc3VyYW5jZUVWUm9vdENBLmNybDA9BgNVHSAENjA0MDIG
BFUdIAAwKjAoBggrBgEFBQcCARYcaHR0cHM6Ly93d3cuZGlnaWNlcnQuY29tL0NQ
UzAdBgNVHQ4EFgQUUWj/kK8CB3U8zNllZGKiErhZcjswHwYDVR0jBBgwFoAUsT7D
aQP4v0cB1JgmGggC72NkK8MwDQYJKoZIhvcNAQELBQADggEBABiKlYkD5m3fXPwd
aOpKj4PWUS+Na0QWnqxj9dJubISZi6qBcYRb7TROsLd5kinMLYBq8I4g4Xmk/gNH
E+r1hspZcX30BJZr01lYPf7TMSVcGDiEo+afgv2MW5gxTs14nhr9hctJqvIni5ly
/D6q1UEL2tU2ob8cbkdJf17ZSHwD2f2LSaCYJkJA69aSEaRkCldUxPUd1gJea6zu
xICaEnL6VpPX/78whQYwvwt/Tv9XBZ0k7YXDK/umdaisLRbvfXknsuvCnQsH6qqF
0wGjIChBWUMo0oHjqvbsezt3tkBigAVBRQHvFwY+3sAzm2fTYS5yh+Rp/BIAV0Ae
cPUeybQ=
-----END CERTIFICATE-----
```

## x509meta

`x509meta` is a reimplementation of `openssl x509`, with several improvements:

- JSON output
- The `--remote` flag can be used to retrieve certificate metadata directly from a remote server
- Can read multiple PEMs at a time
- Greatly simplified usage (5 options compared to `openssl x509`'s 53 options)

It operates in three modes:

- `--remote host:port` - Reads certificates from a remote server. By default, it will only output the server certificate.
Any additonal peer certificates can also be shown by passing the `--chain` flag.
- `--pem file` - Reads one or more certificate PEMs from a file and outputs the x509 metadata.
- `--der file` - Reads a DER encoded certificate from a file and outputs the x509 metadata.

`--pem` and `--der` can also read from stdin by givin them `-` as a value.

Ex:

```
echo '<pem>' | x509meta --pem -
```

`x509meta` will perform certificate verification when operating in `--remote` mode using the golang standard library (`(*x509.Certificate).Verify()`).
If certificate verification fails, it will emit a warning, but continue operating as normal.
It does not check certificate revocation.

### Example Usage

```
$ x509meta --remote github.com
{
  "Version": 3,
  "SerialNumber": "05:57:c8:0b:28:26:83:a1:7b:0a:11:44:93:29:6b:79",
  "Issuer": "CN=DigiCert SHA2 High Assurance Server CA,OU=www.digicert.com,O=DigiCert Inc,C=US",
  "Subject": "CN=github.com,O=GitHub\\, Inc.,L=San Francisco,ST=California,C=US",
  "Validity": {
    "NotAfter": "May 10 12:00:00 2022 UTC",
    "NotBefore": "May 5 00:00:00 2020 UTC"
  },
  "SubjectPublicKeyInfo": {
    "PublicKeyAlgorithm": "RSA",
    "Parameters": {
      "KeySize": 2048,
      "Modulus": "bb:32:b4:d0:d8:9e:9a:9e:8c:79:29:4c:2b:a8:ef:5c:43:d4:93:3b:94:78:ff:30:54:c7:1b:c8:e5:2f:1b:99:cd:10:8d:67:c8:54:0c:07:5a:e4:f4:06:7f:c7:d6:84:b4:2c:cd:2d:4e:44:26:01:1a:ea:37:be:ef:f4:c7:15:07:c3:16:4c:6b:26:19:09:d2:ff:59:10:44:5b:8a:89:81:94:1d:fe:e2:5f:9a:5f:8a:36:d8:b0:e9:1f:6f:80:22:54:ac:ac:29:43:55:52:d8:15:be:92:68:7b:94:56:51:18:d0:a7:d5:c3:5a:47:a8:d8:3c:c6:1d:72:dc:04:36:9d:ac:cf:15:2c:2e:87:d7:f0:fd:49:77:55:ae:ec:4a:a9:db:8b:29:1e:35:67:fe:9d:95:20:dd:79:8d:60:0a:78:73:dc:28:75:a5:86:df:31:fb:13:09:36:a6:c3:e0:2d:46:dc:25:2b:76:f6:ad:f4:c7:7d:f8:68:c2:3b:b3:33:5e:54:2a:df:9b:ae:bf:dc:10:19:40:8d:04:ef:6b:ca:ee:b5:85:3d:2b:d3:8d:82:5f:a9:1b:6b:bb:06:fe:75:e8:3c:26:37:2f:31:cf:dc:0e:8d:37:8e:a5:e8:74:33:d3:7f:7b:0a:bc:72:06:d1:f3:b2:c5:6b:18:b5",
      "Exponent": 65537
    }
  },
  "X509v3Extensions": {
    "KeyUsage": [
      "Digital Signature",
      "Key Encipherment"
    ],
    "ExtendedKeyUsage": [
      "Server Auth",
      "Client Auth"
    ],
    "BasicConstraints": {
      "CA": false,
      "MaxPathLength": -1
    },
    "SubjectKeyIdentifier": "63:02:d2:5d:02:5f:f7:8d:d5:5a:12:9e:76:11:36:96:86:2c:8a:48",
    "AuthorityKeyIdentifier": "51:68:ff:90:af:02:07:75:3c:cc:d9:65:64:62:a2:12:b8:59:72:3b"
  },
  "AuthorityInformation": {
    "OCSP": [
      "http://ocsp.digicert.com"
    ],
    "CAIssuers": [
      "http://cacerts.digicert.com/DigiCertSHA2HighAssuranceServerCA.crt"
    ]
  },
  "SubjectAlternativeNames": [
    "DNS:github.com",
    "DNS:www.github.com"
  ],
  "CertificatePolicies": [
    "2.16.840.1.114412.1.1",
    "2.23.140.1.2.2"
  ],
  "CRLDistributionPoints": [
    "http://crl3.digicert.com/sha2-ha-server-g6.crl",
    "http://crl4.digicert.com/sha2-ha-server-g6.crl"
  ],
  "SignatureAlgorithm": "SHA256-RSA",
  "Signature": "86:32:8f:9c:15:b8:af:e8:d1:de:08:3a:44:0e:71:20:24:d6:fc:0e:58:31:cc:aa:b4:ad:1c:d5:0c:c5:af:c4:bb:fe:5f:ac:90:6a:42:c8:21:eb:25:f1:6b:2c:37:b2:2a:a8:1a:6e:f2:d1:4f:a6:2f:bc:cf:3a:d8:c1:9f:30:c0:ec:93:eb:0a:5a:dc:cb:6c:32:1c:60:6e:ec:6e:f8:86:a5:4f:a0:b4:6d:6a:07:4a:21:58:d0:29:7d:65:8a:c8:da:6a:ba:ab:f0:75:21:33:00:40:6f:85:c5:13:e6:27:73:6c:ae:ea:e3:96:d0:53:db:c1:21:68:10:cf:e3:d8:50:b0:14:ec:a9:98:cf:b8:ce:61:5d:3d:a3:6d:93:34:c4:13:fa:11:66:a3:dd:be:10:19:70:49:e2:04:4d:81:2c:1f:2e:59:c6:2c:53:45:3b:ee:f6:13:f4:d0:2c:84:6e:28:6d:e4:e4:ca:e4:48:89:1b:ab:ec:22:1f:ee:12:d4:6c:75:e9:cc:0b:15:74:e9:6d:9f:db:40:1f:e2:24:85:a3:4b:a4:e9:cd:6b:c8:77:9f:87:4f:05:73:00:38:a5:23:54:68:fc:a2:3d:bf:18:19:0e:a8:fd:b9:5e:8c:5c:e8:fc:e4:a2:52:70:ee:79:a7:d2:27:4a:7a:49"
}

$ readpem github.com > chain.pem

$ x509meta --pem chain.pem
[
  {
    "Version": 3,
    "SerialNumber": "05:57:c8:0b:28:26:83:a1:7b:0a:11:44:93:29:6b:79",
    "Issuer": "CN=DigiCert SHA2 High Assurance Server CA,OU=www.digicert.com,O=DigiCert Inc,C=US",
    "Subject": "CN=github.com,O=GitHub\\, Inc.,L=San Francisco,ST=California,C=US",
    "Validity": {
      "NotAfter": "May 10 12:00:00 2022 UTC",
      "NotBefore": "May 5 00:00:00 2020 UTC"
    },
    "SubjectPublicKeyInfo": {
      "PublicKeyAlgorithm": "RSA",
      "Parameters": {
        "KeySize": 2048,
        "Modulus": "bb:32:b4:d0:d8:9e:9a:9e:8c:79:29:4c:2b:a8:ef:5c:43:d4:93:3b:94:78:ff:30:54:c7:1b:c8:e5:2f:1b:99:cd:10:8d:67:c8:54:0c:07:5a:e4:f4:06:7f:c7:d6:84:b4:2c:cd:2d:4e:44:26:01:1a:ea:37:be:ef:f4:c7:15:07:c3:16:4c:6b:26:19:09:d2:ff:59:10:44:5b:8a:89:81:94:1d:fe:e2:5f:9a:5f:8a:36:d8:b0:e9:1f:6f:80:22:54:ac:ac:29:43:55:52:d8:15:be:92:68:7b:94:56:51:18:d0:a7:d5:c3:5a:47:a8:d8:3c:c6:1d:72:dc:04:36:9d:ac:cf:15:2c:2e:87:d7:f0:fd:49:77:55:ae:ec:4a:a9:db:8b:29:1e:35:67:fe:9d:95:20:dd:79:8d:60:0a:78:73:dc:28:75:a5:86:df:31:fb:13:09:36:a6:c3:e0:2d:46:dc:25:2b:76:f6:ad:f4:c7:7d:f8:68:c2:3b:b3:33:5e:54:2a:df:9b:ae:bf:dc:10:19:40:8d:04:ef:6b:ca:ee:b5:85:3d:2b:d3:8d:82:5f:a9:1b:6b:bb:06:fe:75:e8:3c:26:37:2f:31:cf:dc:0e:8d:37:8e:a5:e8:74:33:d3:7f:7b:0a:bc:72:06:d1:f3:b2:c5:6b:18:b5",
        "Exponent": 65537
      }
    },
    "X509v3Extensions": {
      "KeyUsage": [
        "Digital Signature",
        "Key Encipherment"
      ],
      "ExtendedKeyUsage": [
        "Server Auth",
        "Client Auth"
      ],
      "BasicConstraints": {
        "CA": false,
        "MaxPathLength": -1
      },
      "SubjectKeyIdentifier": "63:02:d2:5d:02:5f:f7:8d:d5:5a:12:9e:76:11:36:96:86:2c:8a:48",
      "AuthorityKeyIdentifier": "51:68:ff:90:af:02:07:75:3c:cc:d9:65:64:62:a2:12:b8:59:72:3b"
    },
    "AuthorityInformation": {
      "OCSP": [
        "http://ocsp.digicert.com"
      ],
      "CAIssuers": [
        "http://cacerts.digicert.com/DigiCertSHA2HighAssuranceServerCA.crt"
      ]
    },
    "SubjectAlternativeNames": [
      "DNS:github.com",
      "DNS:www.github.com"
    ],
    "CertificatePolicies": [
      "2.16.840.1.114412.1.1",
      "2.23.140.1.2.2"
    ],
    "CRLDistributionPoints": [
      "http://crl3.digicert.com/sha2-ha-server-g6.crl",
      "http://crl4.digicert.com/sha2-ha-server-g6.crl"
    ],
    "SignatureAlgorithm": "SHA256-RSA",
    "Signature": "86:32:8f:9c:15:b8:af:e8:d1:de:08:3a:44:0e:71:20:24:d6:fc:0e:58:31:cc:aa:b4:ad:1c:d5:0c:c5:af:c4:bb:fe:5f:ac:90:6a:42:c8:21:eb:25:f1:6b:2c:37:b2:2a:a8:1a:6e:f2:d1:4f:a6:2f:bc:cf:3a:d8:c1:9f:30:c0:ec:93:eb:0a:5a:dc:cb:6c:32:1c:60:6e:ec:6e:f8:86:a5:4f:a0:b4:6d:6a:07:4a:21:58:d0:29:7d:65:8a:c8:da:6a:ba:ab:f0:75:21:33:00:40:6f:85:c5:13:e6:27:73:6c:ae:ea:e3:96:d0:53:db:c1:21:68:10:cf:e3:d8:50:b0:14:ec:a9:98:cf:b8:ce:61:5d:3d:a3:6d:93:34:c4:13:fa:11:66:a3:dd:be:10:19:70:49:e2:04:4d:81:2c:1f:2e:59:c6:2c:53:45:3b:ee:f6:13:f4:d0:2c:84:6e:28:6d:e4:e4:ca:e4:48:89:1b:ab:ec:22:1f:ee:12:d4:6c:75:e9:cc:0b:15:74:e9:6d:9f:db:40:1f:e2:24:85:a3:4b:a4:e9:cd:6b:c8:77:9f:87:4f:05:73:00:38:a5:23:54:68:fc:a2:3d:bf:18:19:0e:a8:fd:b9:5e:8c:5c:e8:fc:e4:a2:52:70:ee:79:a7:d2:27:4a:7a:49"
  },
  {
    "Version": 3,
    "SerialNumber": "04:e1:e7:a4:dc:5c:f2:f3:6d:c0:2b:42:b8:5d:15:9f",
    "Issuer": "CN=DigiCert High Assurance EV Root CA,OU=www.digicert.com,O=DigiCert Inc,C=US",
    "Subject": "CN=DigiCert SHA2 High Assurance Server CA,OU=www.digicert.com,O=DigiCert Inc,C=US",
    "Validity": {
      "NotAfter": "Oct 22 12:00:00 2028 UTC",
      "NotBefore": "Oct 22 12:00:00 2013 UTC"
    },
    "SubjectPublicKeyInfo": {
      "PublicKeyAlgorithm": "RSA",
      "Parameters": {
        "KeySize": 2048,
        "Modulus": "b6:e0:2f:c2:24:06:c8:6d:04:5f:d7:ef:0a:64:06:b2:7d:22:26:65:16:ae:42:40:9b:ce:dc:9f:9f:76:07:3e:c3:30:55:87:19:b9:4f:94:0e:5a:94:1f:55:56:b4:c2:02:2a:af:d0:98:ee:0b:40:d7:c4:d0:3b:72:c8:14:9e:ef:90:b1:11:a9:ae:d2:c8:b8:43:3a:d9:0b:0b:d5:d5:95:f5:40:af:c8:1d:ed:4d:9c:5f:57:b7:86:50:68:99:f5:8a:da:d2:c7:05:1f:a8:97:c9:dc:a4:b1:82:84:2d:c6:ad:a5:9c:c7:19:82:a6:85:0f:5e:44:58:2a:37:8f:fd:35:f1:0b:08:27:32:5a:f5:bb:8b:9e:a4:bd:51:d0:27:e2:dd:3b:42:33:a3:05:28:c4:bb:28:cc:9a:ac:2b:23:0d:78:c6:7b:e6:5e:71:b7:4a:3e:08:fb:81:b7:16:16:a1:9d:23:12:4d:e5:d7:92:08:ac:75:a4:9c:ba:cd:17:b2:1e:44:35:65:7f:53:25:39:d1:1c:0a:9a:63:1b:19:92:74:68:0a:37:c2:c2:52:48:cb:39:5a:a2:b6:e1:5d:c1:dd:a0:20:b8:21:a2:93:26:6f:14:4a:21:41:c7:ed:6d:9b:f2:48:2f:f3:03:f5:a2:68:92:53:2f:5e:e3",
        "Exponent": 65537
      }
    },
    "X509v3Extensions": {
      "KeyUsage": [
        "CRL Sign",
        "Cert Sign",
        "Digital Signature"
      ],
      "ExtendedKeyUsage": [
        "Server Auth",
        "Client Auth"
      ],
      "BasicConstraints": {
        "CA": true,
        "MaxPathLength": 0
      },
      "SubjectKeyIdentifier": "51:68:ff:90:af:02:07:75:3c:cc:d9:65:64:62:a2:12:b8:59:72:3b",
      "AuthorityKeyIdentifier": "b1:3e:c3:69:03:f8:bf:47:01:d4:98:26:1a:08:02:ef:63:64:2b:c3"
    },
    "AuthorityInformation": {
      "OCSP": [
        "http://ocsp.digicert.com"
      ],
      "CAIssuers": null
    },
    "SubjectAlternativeNames": [],
    "CertificatePolicies": [
      "2.5.29.32.0"
    ],
    "CRLDistributionPoints": [
      "http://crl4.digicert.com/DigiCertHighAssuranceEVRootCA.crl"
    ],
    "SignatureAlgorithm": "SHA256-RSA",
    "Signature": "18:8a:95:89:03:e6:6d:df:5c:fc:1d:68:ea:4a:8f:83:d6:51:2f:8d:6b:44:16:9e:ac:63:f5:d2:6e:6c:84:99:8b:aa:81:71:84:5b:ed:34:4e:b0:b7:79:92:29:cc:2d:80:6a:f0:8e:20:e1:79:a4:fe:03:47:13:ea:f5:86:ca:59:71:7d:f4:04:96:6b:d3:59:58:3d:fe:d3:31:25:5c:18:38:84:a3:e6:9f:82:fd:8c:5b:98:31:4e:cd:78:9e:1a:fd:85:cb:49:aa:f2:27:8b:99:72:fc:3e:aa:d5:41:0b:da:d5:36:a1:bf:1c:6e:47:49:7f:5e:d9:48:7c:03:d9:fd:8b:49:a0:98:26:42:40:eb:d6:92:11:a4:64:0a:57:54:c4:f5:1d:d6:02:5e:6b:ac:ee:c4:80:9a:12:72:fa:56:93:d7:ff:bf:30:85:06:30:bf:0b:7f:4e:ff:57:05:9d:24:ed:85:c3:2b:fb:a6:75:a8:ac:2d:16:ef:7d:79:27:b2:eb:c2:9d:0b:07:ea:aa:85:d3:01:a3:20:28:41:59:43:28:d2:81:e3:aa:f6:ec:7b:3b:77:b6:40:62:80:05:41:45:01:ef:17:06:3e:de:c0:33:9b:67:d3:61:2e:72:87:e4:69:fc:12:00:57:40:1e:70:f5:1e:c9:b4"
  }
]
```

## Library

The CLI tools are wrappers around a public API.
You can read the reference docs on [pkg.go.dev](https://pkg.go.dev/github.com/brcrwilliams/tlstools).
