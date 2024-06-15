package fn_valid

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var help bool

func c() {
	cmd := &cobra.Command{}

	cmd.Flags().BoolP("verbose", "v", false, "")
	cmd.Flags().BoolVarP(nil, "help", "h", false, "")
	cmd.Flags().StringP("file", "f", "a_file.txt", "")
}

func p() {
	pflag.BoolP("verbose", "v", false, "")
	pflag.BoolVarP(&help, "help", "h", false, "")
	pflag.StringP("file", "f", "a_file.txt", "")
}
