package main

import (
	"fmt"
	"github.com/docopt/docopt-go"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"strings"
	"syscall"
)

type Params struct {
	Url, User, Password, Dir, File string
}

func getParams(defaults Defaults) Params {
	/* Fill defaults */
	p := Params{}

	/* Parse the arguments */
	args := docoptArgs(defaults)

	/* Fill the parameters */
	// Mandatory
	p.Url = args["<URL>"].(string)
	p.User = args["--user"].(string)

	// Mandatory + OR
	switch {
	case args["--prompt"].(bool):
		p.Password = promptForPassword()
	default:
		p.Password = args["--pass"].(string)
	}

	// Optional + OR
	switch {
	case args["--file"] != nil:
		p.File = args["--file"].(string)
	case args["--dir"] != nil:
		p.Dir = args["--dir"].(string)
	default:
		dir, err := os.Getwd()
		if err != nil {
			dir = os.TempDir()
		}
		p.Dir = dir
	}

	return p
}

func docoptArgs(defaults Defaults) map[string]interface{} {
	versionMsg := "shibdl " + defaults.Version + "."
	usage := versionMsg + "\n" +
		`Download files secured by a Shibboleth IdP.
Code, bugs and feature requests: ` + defaults.Repo + `.
Author: ` + defaults.Author + `.
        _       _       _       _       _       _       _       _
     _-(_)-  _-(_)-  _-(_)-  _-(")-  _-(_)-  _-(_)-  _-(_)-  _-(_)-
   *(___)  *(___)  *(___)  *%%%%%  *(___)  *(___)  *(___)  *(___)
    // \\   // \\   // \\   // \\   // \\   // \\   // \\   // \\

Usage:
  shibdl <URL> -u user [-p password | -P ] [-f file | -d directory] [-l]
  shibdl -h | --help
  shibdl -v | --version

Options:
  <URL>                              URL to download
  -u <user>, --user <user>			 Username
  -p <password>, --pass <password>   Password
  -P, --prompt                       Prompt for password
  -d <directory>, --dir <directory>  Directory to safe file (optional)
  -f <file>, --file <file>           Full path of filename (optional)
  -h, --help                         Show this help screen
  -v, --version                      Show the version message

`
	args, _ := docopt.Parse(usage, nil, true, versionMsg, false)
	return args
}

func promptForPassword() string {
	fd := syscall.Stdin
	fmt.Print("Enter Password: ")
	var password string
	for {
		bytePassword, err := terminal.ReadPassword(fd)
		if err == nil {
			password = string(bytePassword)
			fmt.Println()
			break
		}
	}
	return strings.TrimSpace(password)
}
