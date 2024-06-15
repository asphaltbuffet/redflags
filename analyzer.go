package redflags

import (
	"go/ast"
	"go/token"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var flagMappings = map[string]string{
	"verbose": "v",
	"v":       "verbose",
	"version": "V",
	"V":       "version",
	"help":    "h",
	"h":       "help",
}

type Options struct {
	ShortToLong bool // Enforce short flags determine name of long flags
	LongToShort bool // Enforce long flags determine name of short flags
}

// New creates a new redflags analyzer with the given options.
func New(opts *Options) *analysis.Analyzer {
	if opts == nil {
		opts = &Options{
			ShortToLong: true,
			LongToShort: true,
		}
	}

	return &analysis.Analyzer{
		Name: "redflags",
		Doc:  "ensure consistent long and short flag names in cli applications",
		Run: func(pass *analysis.Pass) (any, error) {
			run(pass)
			return nil, nil
		},
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	}
}

func run(pass *analysis.Pass) {
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	filter := []ast.Node{(*ast.CallExpr)(nil)}

	inspector.Preorder(filter, func(node ast.Node) {
		visit(pass, node)
	})
}

func visit(pass *analysis.Pass, node ast.Node) {
	call := node.(*ast.CallExpr)

	fn, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return
	}

	if !strings.HasSuffix(fn.Sel.Name, "P") {
		return
	}

	var idx int
	if strings.HasSuffix(fn.Sel.Name, "VarP") {
		idx = 1
	}

	if len(call.Args) > 1 {
		long, isLit := call.Args[idx].(*ast.BasicLit)
		if !isLit || long.Kind != token.STRING {
			return
		}

		abbrev, isLit := call.Args[idx+1].(*ast.BasicLit)
		if !isLit || abbrev.Kind != token.STRING {
			return
		}

		flagName := strings.Trim(long.Value, "\"")
		abbrevName := strings.Trim(abbrev.Value, "\"")

		expected, exists := flagMappings[flagName]
		if exists && expected != abbrevName {
			pass.Reportf(node.Pos(), "flag %q should use %q instead of %q", flagName, expected, abbrevName)
			return
		}

		expected, exists = flagMappings[abbrevName]
		if exists && expected != flagName {
			pass.Reportf(node.Pos(), "short flag %q should be for %q instead of %q", abbrevName, expected, flagName)
			return
		}

		return
	}
}
