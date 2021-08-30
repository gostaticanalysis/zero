package zero

import (
	"fmt"
	"go/ast"
	"go/constant"
	"go/token"
	"go/types"
	"reflect"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "zero finds unnecessary assignment which zero value assigns to a variable"

var Analyzer = &analysis.Analyzer{
	Name: "zero",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

var (
	objFalse = types.Universe.Lookup("false")
)

func run(pass *analysis.Pass) (_ interface{}, rerr error) {
	defer func() {
		if v := recover(); v != nil {
			switch v := v.(type) {
			case error:
				rerr = v
			default:
				rerr = fmt.Errorf("panic:%v", v)
			}
		}
	}()

	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.GenDecl)(nil),
		(*ast.AssignStmt)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.GenDecl:
			checkDecl(pass, n)
		case *ast.AssignStmt:
			checkAssign(pass, n)
		}
	})

	return nil, nil
}

func checkDecl(pass *analysis.Pass, decl *ast.GenDecl) {
	if decl == nil || decl.Tok != token.VAR {
		return
	}

	for _, spec := range decl.Specs {
		spec, _ := spec.(*ast.ValueSpec)
		if spec == nil || len(spec.Names) != len(spec.Values) {
			continue
		}

		for i := range spec.Names {
			if spec.Values[i] == nil {
				continue
			}
			typ := pass.TypesInfo.TypeOf(spec.Names[i])
			if isZero(pass, typ, spec.Values[i]) {
				pass.Reportf(spec.Values[i].Pos(), "shoud not assign zero value")
			}
		}
	}
}

func checkAssign(pass *analysis.Pass, n *ast.AssignStmt) {
	if n.Tok != token.DEFINE || len(n.Lhs) != len(n.Rhs) {
		return
	}

	for i := range n.Lhs {
		if n.Rhs[i] == nil {
			continue
		}

		// conversion (including function calling)
		if _, ok := n.Rhs[i].(*ast.CallExpr); ok {
			continue
		}

		typ := pass.TypesInfo.TypeOf(n.Lhs[i])
		if isZero(pass, typ, n.Rhs[i]) {
			pass.Reportf(n.Rhs[i].Pos(), "shoud not assign zero value")
		}
	}
}

func isZero(pass *analysis.Pass, typ types.Type, v ast.Expr) bool {
	switch typ.Underlying().(type) {
	case *types.Basic:
		tv := pass.TypesInfo.Types[v]
		return tv.Value != nil && reflect.ValueOf(constant.Val(tv.Value)).IsZero()
	case *types.Pointer, *types.Slice, *types.Map, *types.Chan, *types.Signature, *types.Interface:
		return types.Identical(pass.TypesInfo.TypeOf(v).Underlying(), types.Typ[types.UntypedNil])
	case *types.Struct, *types.Array:
		clit, _ := v.(*ast.CompositeLit)
		return clit != nil && len(clit.Elts) == 0
	}
	return false
}
