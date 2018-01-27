package internal

import (
	"fmt"
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
				// var
				Tok: token.VAR,
				Specs: []ast.Spec{
					&ast.ValueSpec{
						// Banks
						Names: []*ast.Ident{
							&ast.Ident{
								Name: "Banks",
							},
						},
						Values: []ast.Expr{
							&ast.CompositeLit{
								// map[string]*zengin.Bank
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
								Elts: bankElts(banks),
							},
						},
					},
				},
			},
		},
	}

	return format.Node(w, token.NewFileSet(), f)
}

func bankElts(banks map[string]*zengin.Bank) []ast.Expr {
	var elts []ast.Expr

	for k, bank := range banks {
		expr := &ast.KeyValueExpr{
			Key: &ast.BasicLit{
				Kind:  token.STRING,
				Value: fmt.Sprintf("\"%s\"", k),
			},
			Value: &ast.UnaryExpr{
				Op: token.AND,
				X: &ast.CompositeLit{
					Type: &ast.SelectorExpr{
						X: &ast.Ident{
							Name: "zengin",
						},
						Sel: &ast.Ident{
							Name: "Bank",
						},
					},
					Elts: []ast.Expr{
						&ast.KeyValueExpr{
							Key: &ast.Ident{
								Name: "Code",
							},
							Value: &ast.BasicLit{
								Kind:  token.STRING,
								Value: fmt.Sprintf("\"%s\"", bank.Code),
							},
						},
						&ast.KeyValueExpr{
							Key: &ast.Ident{
								Name: "Name",
							},
							Value: &ast.BasicLit{
								Kind:  token.STRING,
								Value: fmt.Sprintf("\"%s\"", bank.Name),
							},
						},
						&ast.KeyValueExpr{
							Key: &ast.Ident{
								Name: "Kana",
							},
							Value: &ast.BasicLit{
								Kind:  token.STRING,
								Value: fmt.Sprintf("\"%s\"", bank.Kana),
							},
						},
						&ast.KeyValueExpr{
							Key: &ast.Ident{
								Name: "Hira",
							},
							Value: &ast.BasicLit{
								Kind:  token.STRING,
								Value: fmt.Sprintf("\"%s\"", bank.Hira),
							},
						},
						&ast.KeyValueExpr{
							Key: &ast.Ident{
								Name: "Roma",
							},
							Value: &ast.BasicLit{
								Kind:  token.STRING,
								Value: fmt.Sprintf("\"%s\"", bank.Roma),
							},
						},
						&ast.KeyValueExpr{
							Key: &ast.Ident{
								Name: "Branches",
							},
							Value: genBranches(bank.Branches),
						},
					},
				},
			},
		}
		elts = append(elts, expr)
	}

	return elts
}

func genBranches(branches map[string]*zengin.Branch) *ast.CompositeLit {
	lit := &ast.CompositeLit{
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
						Name: "Branch",
					},
				},
			},
		},
	}

	for k, branch := range branches {
		expr := &ast.KeyValueExpr{
			Key: &ast.BasicLit{
				Kind:  token.STRING,
				Value: fmt.Sprintf("\"%s\"", k),
			},
			Value: &ast.UnaryExpr{
				Op: token.AND,
				X: &ast.CompositeLit{
					Type: &ast.SelectorExpr{
						X: &ast.Ident{
							Name: "zengin",
						},
						Sel: &ast.Ident{
							Name: "Bank",
						},
					},
					Elts: []ast.Expr{
						&ast.KeyValueExpr{
							Key: &ast.Ident{
								Name: "Code",
							},
							Value: &ast.BasicLit{
								Kind:  token.STRING,
								Value: fmt.Sprintf("\"%s\"", branch.Code),
							},
						},
						&ast.KeyValueExpr{
							Key: &ast.Ident{
								Name: "Name",
							},
							Value: &ast.BasicLit{
								Kind:  token.STRING,
								Value: fmt.Sprintf("\"%s\"", branch.Name),
							},
						},
						&ast.KeyValueExpr{
							Key: &ast.Ident{
								Name: "Kana",
							},
							Value: &ast.BasicLit{
								Kind:  token.STRING,
								Value: fmt.Sprintf("\"%s\"", branch.Kana),
							},
						},
						&ast.KeyValueExpr{
							Key: &ast.Ident{
								Name: "Hira",
							},
							Value: &ast.BasicLit{
								Kind:  token.STRING,
								Value: fmt.Sprintf("\"%s\"", branch.Hira),
							},
						},
						&ast.KeyValueExpr{
							Key: &ast.Ident{
								Name: "Roma",
							},
							Value: &ast.BasicLit{
								Kind:  token.STRING,
								Value: fmt.Sprintf("\"%s\"", branch.Roma),
							},
						},
					},
				},
			},
		}
		lit.Elts = append(lit.Elts, expr)
	}

	return lit
}
