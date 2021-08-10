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
	if sig == nil || sig.Results().Len() != 1 {
		return false
	}

	ret := sig.Results().At(0).Type()
	return !isShortFuncName(fundecl.Name, ret) && isStructOrPtrStruct(ret) && onlySimpleReturn(fundecl.Body)
}

func isShortFuncName(id *ast.Ident, typ types.Type) bool {
	named, _ := typ.(*types.Named)
	if named == nil {
		return false
	}
	return len(id.Name) < len(named.Obj().Id())
}

func isStructOrPtrStruct(typ types.Type) bool {
	switch typ := typ.Underlying().(type) {
	case *types.Struct:
		return onlyExportedFields(typ)
	case *types.Pointer:
		// no recursive: pointer of pointer become false
		return onlyExportedFields(typ.Elem())
	}
	return false
}

func onlyExportedFields(typ types.Type) bool {
	st, _ := typ.Underlying().(*types.Struct)
	if st == nil {
		return false
	}

	for i := 0; i < st.NumFields(); i++ {
		if !st.Field(i).Exported() {
			return false
		}
	}

	return true
}

func onlySimpleReturn(body *ast.BlockStmt) bool {
	if body == nil || len(body.List) != 1 {
		return false
	}

	ret, _ := body.List[0].(*ast.ReturnStmt)
	if ret == nil || len(ret.Results) != 1 {
		return false
	}

	clit := compositLit(ret.Results[0])
	if clit == nil {
		return false
	}

	for _, expr := range clit.Elts {
		if !identOrLit(expr) {
			return false
		}
	}

	return true
}

func compositLit(expr ast.Expr) *ast.CompositeLit {
	switch expr := expr.(type) {
	case *ast.CompositeLit:
		return expr
	case *ast.UnaryExpr:
		return compositLit(expr.X)
	}
	return nil
}

func identOrLit(expr ast.Expr) bool {
	switch expr := expr.(type) {
	case *ast.Ident:
		return true
	case *ast.BasicLit:
		return true
	case *ast.KeyValueExpr:
		return identOrLit(expr.Value)
	default:
		return false
	}
}
