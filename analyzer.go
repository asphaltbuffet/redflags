// Package redflags provides a linter that enforces consistent long and short flag name pairings
// in CLI applications using cobra or pflag.
package redflags

import (
	"go/ast"
	"go/token"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

//nolint:gochecknoglobals // package-level constant; immutable after init
var defaultFlagPairs = []FlagPair{
	{Long: "help", Short: "h"},
	{Long: "verbose", Short: "v"},
	{Long: "output", Short: "o"},
}

// FlagPair is a long/short flag name pairing.
type FlagPair struct {
	Long  string
	Short string
}

// Options configures the redflags analyzer.
type Options struct {
	UseDefaults bool
	Mappings    []FlagPair
}

// New creates a new redflags analyzer with the given options.
func New(opts *Options) *analysis.Analyzer {
	if opts == nil {
		opts = &Options{UseDefaults: true, Mappings: nil}
	}

	mappings := buildMappings(opts)

	return &analysis.Analyzer{
		Name: "redflags",
		Doc:  "ensure consistent long and short flag names in cli applications",
		Run: func(pass *analysis.Pass) (any, error) {
			run(pass, mappings)
			return nil, nil
		},
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	}
}

func buildMappings(opts *Options) map[string]string {
	m := make(map[string]string)

	if opts.UseDefaults {
		for _, p := range defaultFlagPairs {
			m[p.Long] = p.Short
			m[p.Short] = p.Long
		}
	}

	for _, p := range opts.Mappings {
		m[p.Long] = p.Short
		m[p.Short] = p.Long
	}

	return m
}

func run(pass *analysis.Pass, flagMappings map[string]string) {
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector) //nolint:errcheck // inspect.Analyzer always returns *inspector.Inspector; checked by framework
	filter := []ast.Node{(*ast.CallExpr)(nil)}

	inspector.Preorder(filter, func(node ast.Node) {
		visit(pass, node, flagMappings)
	})
}

func visit(pass *analysis.Pass, node ast.Node, flagMappings map[string]string) {
	call := node.(*ast.CallExpr) //nolint:errcheck // filter guarantees *ast.CallExpr nodes only

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
