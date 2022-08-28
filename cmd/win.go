package cmd

import (
	"fmt"
	"os"

	"github.com/aliaqa256/bash_translator/jsonreader"
	"github.com/aliaqa256/bash_translator/exec"
	"github.com/spf13/cobra"
)

var winCmd = &cobra.Command{
	Use:   "win",
	Short: "Translate windows commands to bash",
	Run: func(cmd *cobra.Command, args []string) {
		var translated string
		// get user input
		input := args[0]
		translations, err := jsonreader.ToStruct()
		must(err)

		// get array of translations
		arrayOfTranslations := translations.Translations

		// loop through array of translations
		for _, translation := range arrayOfTranslations {
			// if input matches translation
			if input == translation.Windows {
				translated = translation.Linux
				
			}
		}

		must(exec.ExecuteLinux(translated))

	},
}

func init() {
	RootCmd.AddCommand(winCmd)
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
