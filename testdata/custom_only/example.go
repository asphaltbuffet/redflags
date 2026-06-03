package custom_only

import "github.com/spf13/cobra"

func f() {
	g := &cobra.Command{}

	// "verbose"/"v" are NOT in the custom mapping, so this should be fine.
	g.Flags().BoolP("verbose", "x", false, "")
	// "deploy"/"d" IS in the custom mapping and correct — no diagnostic.
	g.Flags().BoolP("deploy", "d", false, "")
	// "deploy" with wrong short — should trigger.
	g.Flags().BoolP("deploy", "z", false, "") // want `flag "deploy" should use "d" instead of "z"`
}
