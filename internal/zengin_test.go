package internal

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

func TestOutput(t *testing.T) {
	t.Run("output test", func(t *testing.T) {

		expected, err := os.Open("./banks_test.go")
		if err != nil {
			t.Log("foo")
			t.Error(err)
		}
		defer expected.Close()

		actual, err := os.Create("banks.go")
		if err != nil {
			t.Error(err)
		}
		defer actual.Close()

		if err := Output(actual, Banks); err != nil {
			t.Error(err)
		}

		if _, err := actual.Seek(0, 0); err != nil {
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
			t.Errorf("not equal: output file:\nexpected:\n%s\nactual:\n%s", string(e), string(a))
		}
	})

	if err := os.Remove("banks.go"); err != nil {
		t.Error(err)
	}
}
