package main

import (
	"github.com/codegangsta/cli"
	"github.com/koblas/go-angular-es6/app"
	"github.com/koblas/go-angular-es6/conf"
	"log"
    "os"
)

func main() {
	cliapp := cli.NewApp()
	cliapp.Name = "testapp"
	cliapp.Usage = "Basic test service"
	cliapp.Version = "0.0.1"

	cliapp.Flags = []cli.Flag{
		cli.StringFlag{"config, c", "config.yaml", "config file to use", "APP_CONFIG"},
	}

	cliapp.Commands = []cli.Command{
		{
			Name:  "server",
			Usage: "Run the http server",
			Action: func(c *cli.Context) {
				err := conf.LoadConfig(c)
				if err != nil {
					log.Fatal(err)
					return
				}

                a := app.NewApplication(&conf.Config)

				a.Run()
			},
		},
		{
			Name:  "migratedb",
			Usage: "Perform database migrations",
			Action: func(c *cli.Context) {
				err := conf.LoadConfig(c)
				if err != nil {
					log.Fatal(err)
					return
				}

                a := app.NewApplication(&conf.Config)

				if err = a.Migrate(); err != nil {
					log.Fatal(err)
				}
			},
		},
	}

	cliapp.Run(os.Args)
}
