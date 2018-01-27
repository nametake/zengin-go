package internal

import (
	"go/ast"
	"go/format"
	"go/token"
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
	f := &ast.File{
		Name: ast.NewIdent("internal"),
	}

	return format.Node(w, token.NewFileSet(), f)
}
