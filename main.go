package main

import (
	"fmt"
	"github.com/Equation-Labs-I-O/eqlabs-tools-php-skeleton-creator/commands"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"time"
)

// build version & date are overridden at linking time
var (
	buildVersion = "dev"
	buildDate    string
)

func main() {

	var app = &cobra.Command{
		Use:   "equationlabs-cli [global options] <command> [command options] [arguments...]",
		Short: color.GreenString("EquationLabs CLI"),
		Long: color.GreenString("EquationLabs") + ` CLI version ` + color.YellowString(buildVersion) + ` (c) ` + fmt.Sprintf("(c) 2024-%d Raul Castellanos", time.Now().Year()) + ` (Built on ` + color.YellowString(buildDate) + `)
Helps developers manage php projects scaffolding from local to production environments and following equationlabs development guidelines (DDD, CQRS).`,
		Version: fmt.Sprintf("%s (Compiled on %s)", buildVersion, buildDate),
	}

	// Now we add the commands
	app.AddCommand(commands.CreateProjectCommand)
	app.AddCommand(commands.NewSlimProjectCommand)

	// Run application
	app.Execute()

}
