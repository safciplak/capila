package cli

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"

	"github.com/urfave/cli/v2"

	helpers "github.com/safciplak/capila/src/helpers/environment"
)

// CapilaCLI struct container.
type CapilaCLI struct {
	cli *cli.App
}

// Run starts the Capilla CLI with an array of arguments.
func Run(arguments []string) {
	var capila CapilaCLI

	capila.cli = cli.NewApp()

	capila.setInfo()
	capila.registerMigrations()
	capila.registerBoiler()
	capila.registerCacheCmd()

	expectedErr := capila.cli.Run(arguments)

	handleError(expectedErr)
}

// setInfo sets the CLI information that is shown with help
func (capila CapilaCLI) setInfo() {
	capila.cli.Name = "Capila CLI"
	capila.cli.Usage = "This is a tool that is used for a layer around other CLI or custom commands"
	capila.cli.Version = "1.8.0"
}

// run runs the command and logs the output
func (CapilaCLI) run(_ *cli.Context, command *exec.Cmd) {
	var (
		out    bytes.Buffer
		stderr bytes.Buffer
	)

	command.Stdout = &out
	command.Stderr = &stderr

	// @TODO: Add a prompt for destructive actions on live environment.
	err := command.Run()

	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return
	}

	if stderr.String() != "" {
		fmt.Printf("%s\n", stderr.String())
	}

	if out.String() != "" {
		fmt.Printf("%s\n", out.String())
	}
}

// getOrFail attempts to get the environment value and fails if it cannot be found
func getOrFail(value string) string {
	environmentHelper := new(helpers.EnvironmentHelper)
	environmentString, err := environmentHelper.GetString(value)

	handleError(err)

	return environmentString
}

// handleError handles the errors
func handleError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
