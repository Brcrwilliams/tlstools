package cli

import (
	"os"
	"strings"

	"github.com/brcrwilliams/tlstools"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use: "readpem [host:port]",
		Long: `readpem is a tool to retrieve all the peer certificate PEMs from a remote address.
If no port is given, it will default to port 443.
It will output the certificates to stdout, and can be piped as needed.
Ex: readpem example.com:443 > chain.pem

Use x509meta if you want to see the x509 metadata.
Ex: x509meta --pem chain.pem`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			addr := args[0]
			if !strings.Contains(addr, ":") {
				addr = addr + ":443"
			}
			return getPEMs(addr)
		},
	}
}

func getPEMs(addr string) error {
	certs, err := tlstools.Dial(addr)
	if err != nil {
		return err
	}

	for _, cert := range certs {
		tlstools.WritePEM(os.Stdout, cert)
	}
	return nil
}
