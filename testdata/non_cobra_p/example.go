package non_cobra_p

// FakeFlags mimics cobra's FlagSet API but is not from cobra/pflag.
type FakeFlags struct{}

func (f *FakeFlags) BoolP(name, shorthand string, value bool, usage string) *bool { return nil }

type FakeCommand struct{}

func (c *FakeCommand) Flags() *FakeFlags { return &FakeFlags{} }

func h() {
	cmd := &FakeCommand{}

	// These should NOT be flagged — FakeFlags is not cobra/pflag.
	cmd.Flags().BoolP("verbose", "V", false, "")
	cmd.Flags().BoolP("help", "H", false, "")
}
