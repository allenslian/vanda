package cmd

import (
	"fmt"

	"github.com/urfave/cli"
)

var (
	export bool
)

func getConfigCommand() cli.Command {
	var command = cli.Command{
		Name:    "config",
		Aliases: []string{"c"},
		Usage:   fmt.Sprintf("print the configuration information of %s.", appName),
		Action:  printConfigInfo,
	}

	command.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "export, e",
			Usage:       "export the configuration information with default value.",
			Destination: &export,
		},
	}

	return command
}

func printConfigInfo(c *cli.Context) error {
	if export {
		return exportConfigInfo()
	}

	fmt.Println("print all the configuration file.")
	return nil
}

func exportConfigInfo() error {
	return nil
}
