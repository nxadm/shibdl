package main

import (
    "fmt"
    "github.com/docopt/docopt-go"
    "github.com/howeyc/gopass"
    "golang.org/x/crypto/ssh/terminal"
    "os"
    "runtime"
    "strings"
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
  shibdl <URL> -u user [-p password | -P ] [-f file | -d directory]
  shibdl -h | --help
  shibdl -v | --version

Options:
  <URL>                              URL to download
  -u <user>, --user <user>           Username
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
    fmt.Print("Enter Password: ")
    var password string
    var bytePassword []byte
    var err error
    for {
        switch {
        case runtime.GOOS == "windows":
            bytePassword, err = gopass.GetPasswd()
        default:
            bytePassword, err = terminal.ReadPassword(0)
        }
        if err == nil {
            password = string(bytePassword)
            fmt.Println()
            break
        }
    }
    return strings.TrimSpace(password)
}
