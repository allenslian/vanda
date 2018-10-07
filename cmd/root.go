package cmd

import (
	"fmt"

	"github.com/urfave/cli"

	"github.com/allenslian/vanda/infrastructure/config"
)

var (
	appName = "vanda"
	major   = 1
	minor   = 0
	patch   = "0-alpha"

	configFile string
)

//Execute will run the application.
func Execute(args []string) error {
	app := cli.NewApp()
	app.Name = appName
	app.Version = fmt.Sprintf("%d.%d.%s", major, minor, patch)
	app.Usage = "Run vanda application."
	app.Action = runAction

	// flags
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "config, c",
			Usage:       "specify configuration file path",
			Destination: &configFile,
		},
	}

	// sub commands
	app.Commands = []cli.Command{
		getVersionCommand(),
		getConfigCommand(),
	}

	return app.Run(args)
}

func runAction(c *cli.Context) error {
	config, err := config.LoadConfigFile(appName)
	if err != nil {
		return err
	}
	return setup(config)
}
