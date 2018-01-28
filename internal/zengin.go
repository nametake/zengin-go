package internal

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/format"
	"go/token"
	"io"
	"os"
	"path"
)

const (
	importPath string = `"github.com/nametake/zengin-go"`

	banksFile   string = "banks.json"
	branchesDir string = "branches"
)

func Read(datapath string) (map[string]Bank, error) {
	f, err := os.Open(path.Join(datapath, banksFile))
	if err != nil {
		return nil, err
	}

	var banks map[string]Bank

	if err := json.NewDecoder(f).Decode(&banks); err != nil {
		return nil, err
	}

	for k, b := range banks {
		n := fmt.Sprintf("%s.json", k)
		bf, err := os.Open(path.Join(datapath, branchesDir, n))
		if err != nil {
			return nil, err
		}
		if err := json.NewDecoder(bf).Decode(&b.Branches); err != nil {
			return nil, err
		}

	}

	return banks, nil
}

func Output(w io.Writer, pkgname string, banks map[string]Bank) error {
	f := &ast.File{
		// packange name
		Name: ast.NewIdent(pkgname),
		Decls: []ast.Decl{
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
										X: &ast.Ident{
											Name: "Bank",
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

func bankElts(banks map[string]Bank) []ast.Expr {
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
					Type: &ast.Ident{
						Name: "Bank",
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

func genBranches(branches map[string]Branch) *ast.CompositeLit {
	lit := &ast.CompositeLit{
		Type: &ast.MapType{
			Key: &ast.Ident{
				Name: "string",
			},
			Value: &ast.StarExpr{
				X: &ast.Ident{
					Name: "Branch",
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
					Type: &ast.Ident{
						Name: "Branch",
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
