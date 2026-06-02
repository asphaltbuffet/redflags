// Package gclplugin registers redflags as a golangci-lint module plugin.
package gclplugin

import (
	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"

	"github.com/asphaltbuffet/redflags"
)

func init() { //nolint:gochecknoinits // required by golangci-lint module plugin registration
	register.Plugin("redflags", New)
}

type pluginSettings struct {
	UseDefaults bool       `json:"use-defaults"`
	Mappings    []flagPair `json:"mappings"`
}

type flagPair struct {
	Long  string `json:"long"`
	Short string `json:"short"`
}

// New is the constructor called by golangci-lint with the decoded settings block.
func New(settings any) (register.LinterPlugin, error) {
	s, err := register.DecodeSettings[pluginSettings](settings)
	if err != nil {
		return nil, err
	}

	pairs := make([]redflags.FlagPair, len(s.Mappings))
	for i, m := range s.Mappings {
		pairs[i] = redflags.FlagPair{Long: m.Long, Short: m.Short}
	}

	return &redflagsPlugin{
		opts: &redflags.Options{
			UseDefaults: s.UseDefaults,
			Mappings:    pairs,
		},
	}, nil
}

type redflagsPlugin struct {
	opts *redflags.Options
}

func (p *redflagsPlugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{redflags.New(p.opts)}, nil
}

func (p *redflagsPlugin) GetLoadMode() string {
	return register.LoadModeSyntax
}
