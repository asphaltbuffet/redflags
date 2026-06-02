package custom_override

import "github.com/spf13/cobra"

func f() {
	g := &cobra.Command{}

	// "verbose" default is "v", but we override it to "V" in this test.
	// Using "V" should be fine (no diagnostic).
	g.Flags().BoolP("verbose", "V", false, "")
	// Using the old default "v" for "verbose" should now trigger.
	g.Flags().BoolP("verbose", "v", false, "") // want `flag "verbose" should use "V" instead of "v"`
}
