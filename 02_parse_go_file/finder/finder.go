package finder

import (
	"go/ast"
	"go/parser"
	"go/token"
)

// FindGo is a function looking throu the go-file, specified by fileName, and a function name (funcName)
// It's aim is to find a function -funcName- and to count all the "go" statements inside of it.
func FindInGo(fileName, funcName string) (int, error) {
	fset := token.NewFileSet()
	fileAST, err := parser.ParseFile(fset, fileName, nil, 0)
	if err != nil {
		return 0, err
	}

	counterGo := 0
	ast.Inspect(fileAST, func(x ast.Node) bool {
		funcDecl, ok := x.(*ast.FuncDecl)
		if !ok {
			return true
		}

		// look for the function with given name -funcName-
		if funcDecl.Name.String() == funcName {
			// look for the GO statements
			ast.Inspect(funcDecl, func(x ast.Node) bool {
				_, ok := x.(*ast.GoStmt)
				if !ok {
					return true
				}
				counterGo += 1
				return true
			})
			return false
		}
		return true
	})

	return counterGo, nil
}
