package noctor

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
)

const doc = "noctor finds unnecessary constructor like functions"

var Analyzer = &analysis.Analyzer{
	Name: "noctor",
	Doc:  doc,
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		for _, decl := range file.Decls {
			if isUnnecessaryConstructor(pass, decl) {
				pass.Reportf(decl.Pos(), "unnecessary constructor like function")
			}
		}
	}
	return nil, nil
}

func isUnnecessaryConstructor(pass *analysis.Pass, decl ast.Decl) bool {
	fundecl, _ := decl.(*ast.FuncDecl)
	if fundecl == nil || fundecl.Recv != nil {
		return false
	}

	sig, _ := pass.TypesInfo.TypeOf(fundecl.Name).(*types.Signature)
	return sig != nil && sig.Results().Len() == 1 &&
		isStructOrPtrStruct(sig.Results().At(0).Type()) &&
		onlyReturn(fundecl.Body)
}

func isStructOrPtrStruct(typ types.Type) bool {
	switch typ := typ.Underlying().(type) {
	case *types.Struct:
		return true
	case *types.Pointer:
		// no recursive: pointer of pointer become false
		_, isStruct := typ.Elem().Underlying().(*types.Struct)
		return isStruct
	}
	return false
}

func onlyReturn(body *ast.BlockStmt) bool {
	if len(body.List) != 1 {
		return false
	}
	_, isRet := body.List[0].(*ast.ReturnStmt)
	return isRet
}
