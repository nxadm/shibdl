package main

import (
	"encoding/base64"
	"errors"
	"github.com/headzoo/surf"
	"github.com/headzoo/surf/browser"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func download(params Params) (string, error) {

	/* 1. Prepare the local file */
	// Create a name for the local filename so we exit now instead after
	// the download of a big file
	var filename string
	if params.File == "" {
		uri, err := url.Parse(params.Url)
		if err != nil {
			return "", errors.New("URL can not be parsed")
		}
		parts := strings.Split(uri.Path, "/")
		filename = filepath.FromSlash(params.Dir + "/" + parts[len(parts)-1])
	} else {
		filename = params.File
	}
	// Create the local file
	file, err := os.Create(filename)
	if err != nil {
		return "", errors.New("Can not create local file")
	}


	/* 2. Prepare the request */
	bow := surf.NewBrowser()
	bow.SetUserAgent("shibdl/" + defaults.Version)

	/* 3. Open the url with the file to be downloaded */
	err = bow.Open(params.Url)
	if err != nil {
		return "", errors.New("Connection failed")
	}

	/* 4. Add credentials for basic authentication */
	// ref https://wiki.shibboleth.net/confluence/display/IDP30/PasswordAuthnConfiguration#PasswordAuthnConfiguration-UserInterface
	// "The first user interface layer of the flow is actually HTTP Basic authentication;
	// if a header with credentials is supplied, the credentials are tested immediately with no prompting."
	bow.AddRequestHeader("Authorization", "Basic "+basicAuth(params.User, params.Password))

	/* 5. Keep submitting login forms */
	err = sendForms(bow, defaults.MaxForms)
	if err != nil {
		return "", err
	}

	/* 6. Clean-up secrets */
	bow.DelRequestHeader("Authorization")

	/* 7. Download the file */
	_, err = bow.Download(file)
	if err != nil {
		os.Remove(filename)
		return "", errors.New("Can not download the file")
	}

	return filename, nil
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func sendForms(browser *browser.Browser, maxForms int) error {
	formsCounter := 0
	for {
		formsCounter += 1
		if formsCounter >= maxForms {
			return errors.New("Failed to login (maximal login forms reached)")
			break
		}

		if len(browser.Forms()) >= 1 {
			form := browser.Forms()[1]
			if form != nil && form.Submit() != nil {
				return errors.New("Failed to send the Shibboleth login form")
			}
		} else {
			break
		}
	}
	return nil
}
