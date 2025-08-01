package commands

import "github.com/spf13/cobra"

var DummyShowcaseCommand = &cobra.Command{
	Use:   "showcase",
	Short: "Dummy showcase command for demo purposes",
	Long:  "To demonstrate how to scaffold several commands using cobra",
	Run: func(cmd *cobra.Command, args []string) {
		println("The dummy showcase command...!!!")
	},
}
