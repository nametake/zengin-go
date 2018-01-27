package internal

import (
	"io"

	"github.com/nametake/zengin-go"
)

const (
	banksFile   string = "banks.json"
	branchesDir string = "data"
)

func Read(fpath string) ([]*zengin.Bank, error) {
	panic("not implemented")
}

func Output(w io.Writer, banks map[string]*zengin.Bank) error {
	return nil
}
