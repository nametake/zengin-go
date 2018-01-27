package internal

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"github.com/nametake/zengin-go/internal/testdata"
)

func TestOutput(t *testing.T) {
	expected, err := os.Open("./testdata/zengin.go")
	if err != nil {
		t.Error(err)
	}

	actual, err := os.Create("banks.go")
	if err != nil {
		t.Error(err)
	}

	if err := Output(actual, testdata.Banks); err != nil {
		t.Error(err)
	}

	e, err := ioutil.ReadAll(expected)
	if err != nil {
		t.Error(err)
	}

	a, err := ioutil.ReadAll(actual)
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(e, a) {
		t.Error("not equal")
	}
}
