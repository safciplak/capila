package cli

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

// Test_registerBoiler
func Test_registerBoiler(t *testing.T) {
	setup()

	capila := CapilaCLI{
		cli: cli.NewApp(),
	}

	capila.registerBoiler()

	boilerCommand := capila.cli.Command("boiler")

	// Test 1 parent command
	assert.Equal(t, 1, len(capila.cli.Commands))
	assert.NotNil(t, boilerCommand)
	assert.Greater(t, len(boilerCommand.Subcommands), 0)

	for _, cmd := range boilerCommand.Subcommands {
		// @TODO: IP-72: Mock cmd commands
		// This codes creates files and directories, only run in manual tests
		// err := capila.cli.Run([]string{"capila", "boiler", cmd.Name, "contact"})
		// assert.Nil(t, err)
		//
		// err = capila.cli.Run([]string{"capila", "boiler", cmd.Name, "company", "companies"})
		// assert.Nil(t, err)
		assert.NotNil(t, cmd.Name)
	}

	teardown()
}

// Test_getGitRepository tests the getGitRepository func
// @TODO: IP-72: Mock cmd commands
func Test_getGitRepository(t *testing.T) {
	setup()

	repo := getGitRepository()
	assert.Equal(t, "capila", repo)

	teardown()
}
