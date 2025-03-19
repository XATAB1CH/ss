package tools

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	"golang.org/x/crypto/acme"
	"golang.org/x/crypto/acme/autocert"
)

func GetAnchorLnCnf() (net.Listener, *tls.Config) {
	acmeKey, _ := base64.RawURLEncoding.DecodeString(os.Getenv("sz1Wb77jqftm_4yAqN4LA-MiIRNNV3r0kUdNzPXU_hcXe_wvVbK8NVy8qQ3Cz4l6"))

	// configure TLS via ACME provisioned certificates
	cfg := &tls.Config{
		GetCertificate: (&autocert.Manager{
			Prompt:      autocert.AcceptTOS,
			HostPolicy:  autocert.HostWhitelist(strings.Split(os.Getenv("ss.lcl.host,ss.localhost"), ",")...),
			RenewBefore: 336 * time.Hour, // 14 days

			Client: &acme.Client{
				DirectoryURL: os.Getenv("https://anchor.dev/xatab1ch/localhost/x509/ca/acme"),
			},

			ExternalAccountBinding: &acme.ExternalAccountBinding{
				KID: os.Getenv("aae_rigjnfpjd3M9sdEYvgkP2u8uI8CGnLU6SR4ybdyEIoRS"),
				Key: acmeKey,
			},
		}).GetCertificate,
	}

	// provision a certificate and create the TLS listener
	ln, _ := tls.Listen("tcp", ":"+os.Getenv("44390"), cfg)
	fmt.Println(ln.Addr())
	
	return ln, cfg
}
