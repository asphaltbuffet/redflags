package redflags_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/asphaltbuffet/redflags"
)

func TestOptionsDefaults(t *testing.T) {
	t.Parallel()

	analyzer := redflags.New(nil)
	if analyzer == nil {
		t.Fatal("New(nil) returned nil")
	}

	opts := &redflags.Options{
		UseDefaults: true,
		Mappings:    []redflags.FlagPair{{Long: "output", Short: "o"}},
	}
	analyzer = redflags.New(opts)
	if analyzer == nil {
		t.Fatal("New(opts) returned nil")
	}
}

func TestFuncLinting(t *testing.T) {
	t.Parallel()

	testdata := analysistest.TestData()

	type args struct {
		opts    *redflags.Options
		pattern string
	}

	tests := []struct {
		name string
		args args
	}{
		{"valid", args{nil, "testdata/fn_valid"}},
		{"bad short", args{nil, "testdata/fn_invalid_short"}},
		{"version removed", args{nil, "testdata/version_removed"}},
		{"output default", args{nil, "testdata/output_default"}},
		{"custom override", args{&redflags.Options{
			UseDefaults: true,
			Mappings:    []redflags.FlagPair{{Long: "verbose", Short: "V"}},
		}, "testdata/custom_override"}},
		{"custom only", args{&redflags.Options{
			UseDefaults: false,
			Mappings:    []redflags.FlagPair{{Long: "deploy", Short: "d"}},
		}, "testdata/custom_only"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			analyzer := redflags.New(tt.args.opts)

			analysistest.Run(t, testdata, analyzer, tt.args.pattern)
		})
	}
}
