package cmd

import (
	"fmt"

	"github.com/urfave/cli"
)

func getVersionCommand() cli.Command {
	return cli.Command{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "print version",
		Action: func(c *cli.Context) error {
			fmt.Printf("%s v%d.%d.%s\n", appName, major, minor, patch)
			return nil
		},
	}
}
