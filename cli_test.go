package main

import (
	"fmt"
	"os"
	"testing"
)

func TestGetParams(t *testing.T) {
	os.Args = []string{"shibdl", "https://www.example.com", "-u", "foo", "-p", "bar"}
	params := getParams(defaults)
	fmt.Printf("PARAMS: %+v\n", params)
}
