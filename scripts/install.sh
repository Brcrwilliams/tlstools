#!/bin/sh

cd "$(dirname "$0")/.." || exit 1

go build -o /usr/local/bin/readpem ./cmd/readpem
go build -o /usr/local/bin/x509meta ./cmd/x509meta
