package main

import (
	"fmt"
	"os"
	"github.com/briandowns/spinner"
	"time"
)

// Use SURF_DEBUG_HEADERS=1 environment variable to print debug headers.


/* Application defaults */
type Defaults struct {
	Author, Repo, Version string
	MaxForms              int
}

const author = "Claudio Ramirez <pub.claudio@gmail.com>"
const repo = "https://github.com/nxadm/shib-download-file"
const version = "v0.4.1"
const maxForms = 5

var defaults = Defaults{
	Author:  author,
	Repo:    repo,
	Version: version,
	MaxForms: maxForms,
}

//
func main() {

	/* Command line interface */
	params := getParams(defaults)
	fmt.Println("Downloading file...")
	spinner := spinner.New(spinner.CharSets[35], 500*time.Millisecond)  // Build our new spinner
	spinner.Start()                                                    // Start the spinner
	file, err := download(params)
	spinner.Stop()
	fmt.Println("")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error downloading the file: %s\n", err)
		os.Exit(2)
	}
	fmt.Printf("File downloaded as %s.\n", file)
}
