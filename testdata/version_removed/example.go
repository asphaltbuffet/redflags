package version_removed

import "github.com/spf13/cobra"

func f() {
	g := &cobra.Command{}

	// "version"/"V" was a built-in default before; it should no longer trigger.
	g.Flags().BoolP("version", "V", false, "")
}
