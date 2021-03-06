package cli

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/brcrwilliams/tlstools"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	var addr string
	var pem string
	var der string
	var chain bool
	cmd := &cobra.Command{
		Use: "x509meta [--remote host:port | --pem filepath | --der filepath] [--chain]",
		Long: `x509meta is a tool for inspecting x509 certificate attributes.
It will output a JSON document in a similar format to ` + "`openssl x509 -text -noout`." + `
It can read certificates from a remote address, or from a file.
If --chain is given, then it will return the full certificate chain from the remote address.
Otherwise, it will only give the server certificate. When given a filepath, it will always
print all certificates in the file. If you wish to read from stdin, give "-" as the value
to either --pem or --der.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			count := 0
			if addr != "" {
				count++
			}
			if pem != "" {
				count++
			}
			if der != "" {
				count++
			}
			if count == 0 {
				return fmt.Errorf("At least one of --remote, --pem, or --der is required")
			}
			if count > 1 {
				return fmt.Errorf("Only one of --remote, --pem, or --der can be used at a time")
			}
			if addr != "" {
				if !strings.Contains(addr, ":") {
					addr = addr + ":443"
				}
				return getRemote(addr, chain)
			}
			if pem != "" {
				return getPEM(pem)
			}
			if der != "" {
				return getDER(der)
			}
			return fmt.Errorf("I dunno how you got here, but this is a bug.")
		},
	}

	flags := cmd.Flags()
	flags.StringVar(&addr, "remote", "", "The host:port of the remote address to read certificates from: example.com:443 (defaults to port 443)")
	flags.StringVar(&pem, "pem", "", "The path to a PEM file to read certificates from.")
	flags.StringVar(&der, "der", "", "The path to a DER file to read certficates from.")
	flags.BoolVar(&chain, "chain", false, "Used with --remote - If given, will print the full chain of certs.")
	return cmd
}

func getRemote(addr string, chain bool) error {
	certs, err := tlstools.Dial(addr)
	if err != nil {
		return err
	}

	if !chain {
		return tlstools.WriteX509Meta(os.Stdout, certs[0])
	}

	return tlstools.WriteX509Metas(os.Stdout, certs)
}

func getReader(input string) (io.ReadCloser, error) {
	if input == "-" {
		return os.Stdin, nil
	}

	f, err := os.Open(input)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}

	return f, nil
}

func getPEM(path string) error {
	reader, err := getReader(path)
	if err != nil {
		return err
	}
	defer reader.Close()

	certs, err := tlstools.ReadPEM(reader)
	if err != nil {
		return err
	}

	return tlstools.WriteX509Metas(os.Stdout, certs)
}

func getDER(path string) error {
	reader, err := getReader(path)
	if err != nil {
		return err
	}
	defer reader.Close()

	cert, err := tlstools.ReadDER(reader)
	if err != nil {
		return err
	}

	return tlstools.WriteX509Meta(os.Stdout, cert)
}
