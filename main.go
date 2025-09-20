package main

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/racastellanosm/tools-php-skeleton-creator/commands"
	"github.com/spf13/cobra"
)

// build version & date are overridden at linking time
var (
	buildVersion = "dev"
	buildDate    string
)

func main() {

	var app = &cobra.Command{
		Use:   "php-skeleton-creator-cli [global options] <command> [command options] [arguments...]",
		Short: color.GreenString("PHP Skeleton Creator CLI"),
		Long: color.GreenString("PHP Skeleton Creator") + ` CLI version ` + color.YellowString(buildVersion) + ` (c) ` + fmt.Sprintf("(c) 2024-%d Raul Castellanos", time.Now().Year()) + ` (Built on ` + color.YellowString(buildDate) + `)
Creates scaffolded projects following recommended development guidelines (DDD, CQRS, Testing, Automation).`,
		Version: color.GreenString(buildVersion) + ` (c) ` + fmt.Sprintf("2024-%d Raul Castellanos", time.Now().Year()) + ` (On ` + color.YellowString(buildDate) + `)`,
	}

	// Now we add the commands
	app.AddCommand(commands.CreateProjectWithSymfonyCommand)
	app.AddCommand(commands.CreateProjectWithSlimCommand)
	app.AddCommand(commands.CompletionCommand)

	// Run application
	_ = app.Execute()

}
