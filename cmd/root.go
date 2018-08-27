package cmd

import (
	"fmt"

	"github.com/spf13/viper"
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
	config, err := readConfigFile()
	if err != nil {
		return err
	}
	fmt.Println(config.Database.DefaultURI)
	return nil
}

func readConfigFile() (*config.Configuration, error) {
	v := viper.New()
	v.SetConfigName(appName)
	v.SetConfigType("toml")
	v.AddConfigPath(".")
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	var c config.Configuration
	if err := v.Unmarshal(&c); err != nil {
		return nil, err
	}

	return &c, nil
}
