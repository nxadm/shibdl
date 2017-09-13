package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"testing"
)

func TestDownload(t *testing.T) {
	testParams := Params{}
	err := testParams.getTestParams()
	if err != nil {
		t.Error(err)
	}

	dir, _ := os.Getwd()
	testParams.Dir = dir
	_, err = download(testParams)
	if err != nil {
		t.Error(err)
	}
}

func (p *Params) getTestParams() error {
	// See sampleTestConnectionParams.json for a test configuration file
	// Load it as an environment variable, e.g.:
	// export SHIBDL_CONNECTION_PARAMS=/var/tmp/params.json"
	config := os.Getenv("SHIBDL_CONNECTION_PARAMS")
	if config == "" {
		return nil
	}

	raw, err := ioutil.ReadFile(config)
	if err != nil {
		return errors.New("Can not read the json configuration provided ($SHIBDL_CONNECTION_PARAMS)")
	}

	err = json.Unmarshal(raw, p)
	if err != nil {
		return errors.New("Can not parse the json configuration provided ($SHIBDL_CONNECTION_PARAMS)")
	}

	return nil
}
