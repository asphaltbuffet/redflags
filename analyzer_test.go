package redflags_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/asphaltbuffet/redflags"
)

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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			analyzer := redflags.New(tt.args.opts)

			analysistest.Run(t, testdata, analyzer, tt.args.pattern)
		})
	}
}
