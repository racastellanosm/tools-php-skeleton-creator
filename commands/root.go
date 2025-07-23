package commands

import (
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// BuildVersion and BuildDate are set during the build process using ldflags.
var (
	BuildVersion = "dev"
	BuildDate    = "unknown"
)

var rootCmd = &cobra.Command{
	Use:   "equationlabs-cli [global options] <command> [command options] [arguments...]",
	Short: color.GreenString("EquationLabs CLI"),
	Long: color.GreenString("EquationLabs") + ` CLI version ` + color.YellowString(BuildVersion) + ` (c) ` + fmt.Sprintf("(c) 2024-%d Raul Castellanos", time.Now().Year()) + ` (Built on ` + color.YellowString(BuildDate) + `)
Helps developers manage php projects scaffolding from local to production environments and provides a set of tools to automate common tasks.`,
	Version: fmt.Sprintf("%s (Compiled on %s)", BuildVersion, BuildDate),
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func init() {
	// Set global flags available to all commands
	rootCmd.PersistentFlags().BoolP("help", "h", false, "Show help")
	rootCmd.PersistentFlags().Bool("no-interaction", false, "Disable all interactions")
	rootCmd.PersistentFlags().BoolP("version", "V", false, "Print the version")
	rootCmd.PersistentFlags().Bool("no-git", false, "Do not initialize Git")
	// Add commands to the root command
	rootCmd.AddCommand(newCommand)
	// rootCmd.AddCommand(composerCmd) // Añadir el comando 'composer'
}
