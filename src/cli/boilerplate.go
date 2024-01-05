package cli

import (
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"os"
	"os/exec"
	"strings"
	"text/template"

	"github.com/urfave/cli/v2"

	"github.com/safciplak/capila/src/cli/templates"
)

const defaultBusinessPath = "src/business/"

// boilerInfoCommand returns the boiler info command
func (capila CapilaCLI) boilerInfoCommand() *cli.Command {
	return &cli.Command{
		Name:  "info",
		Usage: "Shows information about the boiler package",
		Action: func(context *cli.Context) error {
			return nil
		},
	}
}

// getGitRepository retrieves the current git repository name
func getGitRepository() string {
	var (
		repositoryURL []byte
		err           error
	)

	repositoryURL, err = exec.Command("git", "config", "--get", "remote.origin.url").CombinedOutput()
	if err != nil {
		return os.Getenv("APPLICATION_NAME")
	}

	splittedURL := strings.Split(string(repositoryURL), "/")

	return strings.Split(splittedURL[len(splittedURL)-1], ".")[0]
}

// createDirCommand returns the create dir command
func createDirCommand(name string) *exec.Cmd {
	dir := defaultBusinessPath + name

	return exec.Command("mkdir", "-p", dir)
}

// generateHandlerInterfaceBoilerplate generates the interface Handler boilerplate
func generateHandlerInterfaceBoilerplate(names map[string]string) *exec.Cmd {
	path := defaultBusinessPath + names["PrivateNamePlural"] + "/handlers/"
	file := path + "interface" + names["PublicNameSingular"] + "Handler.go"

	return generateBoilerplate(names, file, templates.HandlerInterface)
}

// generateHandlerBoilerplate generates the Handler boilerplate
func generateHandlerBoilerplate(names map[string]string) *exec.Cmd {
	path := defaultBusinessPath + names["PrivateNamePlural"] + "/handlers/"
	file := path + names["PrivateNameSingular"] + "Handler.go"

	return generateBoilerplate(names, file, templates.Handler)
}

// generateHandlerTestBoilerplate generates the Handler test boilerplate
func generateHandlerTestBoilerplate(names map[string]string) *exec.Cmd {
	path := defaultBusinessPath + names["PrivateNamePlural"] + "/handlers/"
	file := path + names["PrivateNameSingular"] + "Handler_test.go"

	return generateBoilerplate(names, file, templates.HandlerTest)
}

// generateServiceInterfaceBoilerplate generates the Service boilerplate
func generateServiceInterfaceBoilerplate(names map[string]string) *exec.Cmd {
	path := defaultBusinessPath + names["PrivateNamePlural"] + "/services/"
	file := path + "interface" + names["PublicNameSingular"] + "Service.go"

	return generateBoilerplate(names, file, templates.ServiceInterface)
}

// generateHandlerBoilerplate generates the Service boilerplate
func generateServiceBoilerplate(names map[string]string) *exec.Cmd {
	path := defaultBusinessPath + names["PrivateNamePlural"] + "/services/"
	file := path + names["PrivateNameSingular"] + "Service.go"

	return generateBoilerplate(names, file, templates.Service)
}

// generateServiceTestBoilerplate generates the Service test boilerplate
func generateServiceTestBoilerplate(names map[string]string) *exec.Cmd {
	path := defaultBusinessPath + names["PrivateNamePlural"] + "/services/"
	file := path + names["PrivateNameSingular"] + "Service_test.go"

	return generateBoilerplate(names, file, templates.ServiceTest)
}

// generateRepositoryInterfaceBoilerplate generates the Repository boilerplate
func generateRepositoryInterfaceBoilerplate(names map[string]string) *exec.Cmd {
	path := defaultBusinessPath + names["PrivateNamePlural"] + "/repositories/"
	file := path + "interface" + names["PublicNameSingular"] + "Repository.go"

	return generateBoilerplate(names, file, templates.RepositoryInterface)
}

// generateRepositoryBoilerplate generates the Repository boilerplate
func generateRepositoryBoilerplate(names map[string]string) *exec.Cmd {
	path := defaultBusinessPath + names["PrivateNamePlural"] + "/repositories/"
	file := path + names["PrivateNameSingular"] + "Repository.go"

	return generateBoilerplate(names, file, templates.Repository)
}

// generateRepositoryTestBoilerplate generates the Repository boilerplate
func generateRepositoryTestBoilerplate(names map[string]string) *exec.Cmd {
	path := defaultBusinessPath + names["PrivateNamePlural"] + "/repositories/"
	file := path + names["PrivateNameSingular"] + "Repository_test.go"

	return generateBoilerplate(names, file, templates.RepositoryTest)
}

// generateRequestModelsBoilerplate generates the RequestModels boilerplate
func generateRequestModelsBoilerplate(names map[string]string) *exec.Cmd {
	path := defaultBusinessPath + names["PrivateNamePlural"] + "/models/"
	file := path + "requestModels.go"

	return generateBoilerplate(names, file, templates.RequestModels)
}

// generateModelBoilerplate generates the RequestModels boilerplate
func generateModelBoilerplate(names map[string]string) *exec.Cmd {
	path := "src/models/"
	file := path + names["PrivateNameSingular"] + ".go"

	return generateBoilerplate(names, file, templates.Models)
}

// generateBoilerplate generates the boilerplate's
func generateBoilerplate(names map[string]string, file, boilerplate string) *exec.Cmd {
	outfile, err := os.Create(file)
	if err != nil {
		panic(err)
	}
	defer outfile.Close()

	t := template.Must(template.New("").Parse(boilerplate))
	err = t.Execute(outfile, names)

	if err != nil {
		panic(err)
	}

	return exec.Command("echo", "file generated")
}

// registerBoiler registers boiler commands
func (capila CapilaCLI) registerBoiler() {
	cmd := &cli.Command{
		Name:  "boiler",
		Usage: "Run and check for boiler CLI",
		Subcommands: []*cli.Command{
			capila.boilerInfoCommand(),
			capila.boilerCreateCommand(),
		},
	}

	capila.cli.Commands = append(capila.cli.Commands, cmd)
}

// boilerCreateCommand returns the migrate create command
//
//nolint:funlen // Skip check, registers all commands
func (capila CapilaCLI) boilerCreateCommand() *cli.Command {
	return &cli.Command{
		Name:  "create",
		Usage: "Supply a name and a new migration ( up & down ) will be created",
		Action: func(context *cli.Context) error {
			privateNameSingular := context.Args().First()
			privateNamePlural := context.Args().Get(1)

			if privateNameSingular == "" {
				fmt.Printf("%s", "Please provide a name \n")
				return nil
			}

			if privateNamePlural == "" {
				privateNamePlural = privateNameSingular + "s"
			}

			names := map[string]string{
				"ApplicationName":          getGitRepository(),
				"PrivateNameSingular":      privateNameSingular,
				"PrivateNamePlural":        privateNamePlural,
				"PublicNameSingular":       cases.Title(language.Und, cases.NoLower).String(privateNameSingular),
				"PublicNamePlural":         cases.Title(language.Und, cases.NoLower).String(privateNamePlural),
				"BaseRequestGUIDBinding":   "`uri:\"guid\" binding:\"required,uuid\"`",
				"CreateRequestGUIDBinding": "`json:\"guid\" binding:\"omitempty,uuid\"`",
				"UpdateRequestGUIDBinding": "`json:\"-\" uri:\"guid\" binding:\"required,uuid\"`",
				"LanguageBinding":          "`json:\"-\" binding:\"omitempty,min=2,max=5\"`",
			}

			// create models
			capila.run(context, createDirCommand(privateNamePlural+"/models"))
			capila.run(context, generateRequestModelsBoilerplate(names))

			capila.run(context, exec.Command("mkdir", "-p", "src/models"))
			capila.run(context, generateModelBoilerplate(names))

			// create repositories
			capila.run(context, createDirCommand(privateNamePlural+"/repositories"))
			capila.run(context, generateRepositoryInterfaceBoilerplate(names))
			capila.run(context, generateRepositoryBoilerplate(names))

			// create services
			capila.run(context, createDirCommand(privateNamePlural+"/services"))
			capila.run(context, generateServiceInterfaceBoilerplate(names))
			capila.run(context, generateServiceBoilerplate(names))

			// create handlers
			capila.run(context, createDirCommand(privateNamePlural+"/handlers"))
			capila.run(context, generateHandlerInterfaceBoilerplate(names))
			capila.run(context, generateHandlerBoilerplate(names))

			// create test files
			capila.run(context, generateRepositoryTestBoilerplate(names))
			capila.run(context, generateServiceTestBoilerplate(names))
			capila.run(context, generateHandlerTestBoilerplate(names))

			// go commands
			capila.run(context, exec.Command("gofmt", "-s", "-w", "."))

			return nil
		},
	}
}
