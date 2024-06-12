package analyzer

import (
	"go/ast"
	"go/token"
	"strings"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "redflags",
	Doc:  "Checks for consistent flag names in cli applications",
	Run:  run,
}

var flagMappings = map[string]string{
	"verbose": "v",
	"v":       "verbose",
	"version": "V",
	"V":       "version",
	"help":    "h",
	"h":       "help",
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := func(node ast.Node) bool {
		x, ok := node.(*ast.CallExpr)
		if !ok {
			return true
		}

		fun, ok := x.Fun.(*ast.SelectorExpr)
		if !ok {
			return true
		}

		// only considering {Bool,Int,String,...}P flags for now
		if !strings.HasSuffix(fun.Sel.Name, "P") {
			return true
		}
		var idx int
		if strings.HasSuffix(fun.Sel.Name, "VarP") {
			idx = 1
		}

		if len(x.Args) > 1 {
			long, isLit := x.Args[idx].(*ast.BasicLit)
			if !isLit || long.Kind != token.STRING {
				return true
			}

			abbrev, isLit := x.Args[idx+1].(*ast.BasicLit)
			if !isLit || abbrev.Kind != token.STRING {
				return true
			}

			flagName := strings.Trim(long.Value, "\"")
			abbrevName := strings.Trim(abbrev.Value, "\"")

			expected, exists := flagMappings[flagName]
			if exists && expected != abbrevName {
				pass.Reportf(node.Pos(), "flag %q should use %q instead of %q", flagName, expected, abbrevName)
				return true
			}

			expected, exists = flagMappings[abbrevName]
			if exists && expected != flagName {
				pass.Reportf(node.Pos(), "short flag %q should be for %q instead of %q", abbrevName, expected, flagName)
				return true
			}

			return true
		}

		return true
	}

	for _, f := range pass.Files {
		ast.Inspect(f, inspect)
	}

	return nil, nil //nolint: nilnil // optional return
}
