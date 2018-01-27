package internal

import (
	"go/ast"
	"go/format"
	"go/token"
	"io"

	"github.com/nametake/zengin-go"
)

const (
	importPath  string = `"github.com/nametake/zengin-go"`
	banksFile   string = "banks.json"
	branchesDir string = "data"
)

func Read(fpath string) ([]*zengin.Bank, error) {
	panic("not implemented")
}

func Output(w io.Writer, banks map[string]*zengin.Bank) error {
	f := &ast.File{
		// packange name
		Name: ast.NewIdent("internal"),
		Decls: []ast.Decl{
			// import zengin-go
			&ast.GenDecl{
				Tok: token.IMPORT,
				Specs: []ast.Spec{
					&ast.ImportSpec{
						Name: &ast.Ident{
							Name: "zengin",
						},
						Path: &ast.BasicLit{
							Kind:  token.STRING,
							Value: importPath,
						},
					},
				},
			},
			// banks
			&ast.GenDecl{
				Tok: token.VAR,
				Specs: []ast.Spec{
					&ast.ValueSpec{
						Names: []*ast.Ident{
							&ast.Ident{
								Name: "Banks",
							},
						},
						Values: []ast.Expr{
							&ast.CompositeLit{
								Type: &ast.MapType{
									Key: &ast.Ident{
										Name: "string",
									},
									Value: &ast.StarExpr{
										X: &ast.SelectorExpr{
											X: &ast.Ident{
												Name: "zengin",
											},
											Sel: &ast.Ident{
												Name: "Bank",
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	return format.Node(w, token.NewFileSet(), f)
}
