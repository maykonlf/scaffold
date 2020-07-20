package command

import "github.com/spf13/cobra"

var templateCmd = &cobra.Command{
	Use:     "template",
	Short:   "manage project templates",
	Example: "scaffold template",
}
