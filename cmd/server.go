package main

import (
	"github.com/codegangsta/cli"
	"github.com/koblas/go-angular-es6/conf"
	"github.com/koblas/go-angular-es6/service"
	"log"
    "os"
)

func main() {
	app := cli.NewApp()
	app.Name = "testapp"
	app.Usage = "Basic test service"
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		cli.StringFlag{"config, c", "config.yaml", "config file to use", "APP_CONFIG"},
	}

	app.Commands = []cli.Command{
		{
			Name:  "server",
			Usage: "Run the http server",
			Action: func(c *cli.Context) {
				err := conf.LoadConfig(c)
				if err != nil {
					log.Fatal(err)
					return
				}

				svc := service.LikeMarkService{}

				if err = svc.Run(&conf.Config); err != nil {
					log.Fatal(err)
				}
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

				svc := service.LikeMarkService{}

				if err = svc.Migrate(&conf.Config); err != nil {
					log.Fatal(err)
				}
			},
		},
	}
	app.Run(os.Args)
}
