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
	"github.com/bmmcginty/go-openal/openal"
	_ "github.com/bmmcginty/gumble/opus"
)

func show_devs(name string, args []string) {
if args == nil {
fmt.Printf("no items for %s\n",name)
}
fmt.Printf("%s\n",name)
for i:=0; i<len(args); i++ {
fmt.Printf("%s\n",args[i])
}
}

func do_list_devices() {
odevs := openal.GetStrings(openal.AllDevicesSpecifier)
if odevs!=nil && len(odevs)>0 {
show_devs("All outputs:",odevs)
} else {
odevs=openal.GetStrings(openal.DeviceSpecifier)
show_devs("All outputs:",odevs)
}
idevs := openal.GetStrings(openal.CaptureDeviceSpecifier)
show_devs("Inputs:",idevs)
}

func main() {
	// Command line flags
	server := flag.String("server", "localhost:64738", "the server to connect to")
	username := flag.String("username", "", "the username of the client")
	password := flag.String("password", "", "the password of the server")
	insecure := flag.Bool("insecure", false, "skip server certificate verification")
	certificate := flag.String("certificate", "", "PEM encoded certificate and private key")
	cfgfn := flag.String("config", "~/.barnard.yaml", "Path to YAML formatted configuration file")
	list_devices := flag.Bool("list_devices", false, "do not connect; instead, list available audio devices and exit")
	serverSet := false
	usernameSet := false

	flag.Parse()
	userConfig := config.NewConfig(cfgfn)

	flag.CommandLine.Visit(func (theFlag *flag.Flag) {
		switch theFlag.Name {
		case "server":
			serverSet = true
		case "username":
			usernameSet = true
		}
	})

	if ! serverSet {
		server = userConfig.GetDefaultServer()
	}
	if !usernameSet {
		username = userConfig.GetUsername()
	}

	if os.Getenv("ALSOFT_LOGLEVEL") == "" {
		os.Setenv("ALSOFT_LOGLEVEL", "0")
	}

if (*list_devices) {
do_list_devices()
os.Exit(0)
}

if !strings.Contains(*server,":") {
*server=(*server+":64738")
}

	// Initialize
	b := Barnard{
		Config:     gumble.NewConfig(),
		UserConfig: userConfig,
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
