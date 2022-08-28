package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "translate",
	Short: "Translate windows commands to bash",
}
