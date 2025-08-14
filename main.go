package main

import (
	"embed"
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

	//go:embed templates/*
	templates embed.FS
)

func main() {

	var app = &cobra.Command{
		Use:   "equationlabs-php-cli [global options] <command> [command options] [arguments...]",
		Short: color.GreenString("EquationLabs CLI"),
		Long: color.GreenString("EquationLabs") + ` CLI version ` + color.YellowString(buildVersion) + ` (c) ` + fmt.Sprintf("(c) 2024-%d Raul Castellanos", time.Now().Year()) + ` (Built on ` + color.YellowString(buildDate) + `)
Creates scaffolded projects following equationlabs development guidelines (DDD, CQRS, Testing, Automation).`,
		Version: color.GreenString(buildVersion) + ` (c) ` + fmt.Sprintf("2024-%d Raul Castellanos", time.Now().Year()) + ` (On ` + color.YellowString(buildDate) + `)`,
	}

	// Now we add the commands
	app.AddCommand(commands.CreateProjectWithSymfonyCommand)
	app.AddCommand(commands.CreateProjectWithSlimCommand)
	app.AddCommand(commands.CompletionCommand)

	// Run application
	_ = app.Execute()

}
