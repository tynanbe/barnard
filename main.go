package main

import (
"fmt"
"os"
"strings"
"github.com/bmmcginty/barnard/config"
	"crypto/tls"
	"flag"

	"github.com/bmmcginty/barnard/uiterm"
	"github.com/bmmcginty/gumble/gumble"
	_ "github.com/bmmcginty/gumble/opus"
)

func main() {
	// Command line flags
	server := flag.String("server", "localhost:64738", "the server to connect to")
	username := flag.String("username", "", "the username of the client")
	password := flag.String("password", "", "the password of the server")
	insecure := flag.Bool("insecure", false, "skip server certificate verification")
	certificate := flag.String("certificate", "", "PEM encoded certificate and private key")

	flag.Parse()

if !strings.Contains(*server,":") {
*server=(*server+":64738")
}

	// Initialize
	b := Barnard{
		Config:     gumble.NewConfig(),
		UserConfig: config.NewConfig(),
		Address:    *server,
	}

b.Hotkeys=b.UserConfig.GetHotkeys()
b.UserConfig.SaveConfig()
	b.Config.Username = *username
	b.Config.Password = *password

	if *insecure {
		b.TLSConfig.InsecureSkipVerify = true
	}
	if *certificate != "" {
		cert, err := tls.LoadX509KeyPair(*certificate, *certificate)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
		b.TLSConfig.Certificates = append(b.TLSConfig.Certificates, cert)
	}

	b.Ui = uiterm.New(&b)
	b.Ui.Run()
if b.exitMessage!="" {
fmt.Fprintf(os.Stderr,"%s\n",b.exitMessage)
}
os.Exit(b.exitStatus)
}
