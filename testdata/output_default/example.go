package output_default

import "github.com/spf13/cobra"

func f() {
	g := &cobra.Command{}

	// correct pairing — no diagnostic
	g.Flags().StringP("output", "o", "", "")
	// wrong short for "output" — should trigger
	g.Flags().StringP("output", "x", "", "") // want `flag "output" should use "o" instead of "x"`
}
