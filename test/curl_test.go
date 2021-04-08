package test

import (
	"testing"

	"github.com/itscharlieliu/golang-curl/pkg"
)

func TestCurlNoArgs(t *testing.T) {

	commandOutputString := pkg.RunCmd("go", []string{"run", "../cmd/curl.go"})

	realOutputString := pkg.RunCmd("curl", []string{})

	if commandOutputString != realOutputString {
		t.Errorf("Expected %s, got %s", realOutputString, commandOutputString)
	}
}
