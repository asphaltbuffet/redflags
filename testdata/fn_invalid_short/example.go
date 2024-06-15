package fn_invalid_short

import (
	"github.com/spf13/cobra"
)

func f() {
	g := &cobra.Command{}

	g.Flags().BoolP("verbose", "V", false, "")      // want `flag "verbose" should use "v" instead of "V"`
	g.Flags().BoolVarP(nil, "help", "H", false, "") // want `flag "help" should use "h" instead of "H"`
}
