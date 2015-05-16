package main

import (
    "gopkg.in/yaml.v2"
	"github.com/codegangsta/cli"
	"github.com/koblas/likemark/conf"
	"github.com/koblas/likemark/service"
	"log"
    "os"
    "errors"
    "io/ioutil"
)

func getConfig(c *cli.Context) (conf.Config, error) {
	yamlPath := c.GlobalString("config")
	config := conf.Config{}

	if _, err := os.Stat(yamlPath); err != nil {
		return config, errors.New("config path not valid")
	}

	ymlData, err := ioutil.ReadFile(yamlPath)
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal([]byte(ymlData), &config)
	return config, err
}

func main() {
	app := cli.NewApp()
	app.Name = "likemark"
	app.Usage = "Basic likemark service"
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		cli.StringFlag{"config, c", "config.yaml", "config file to use", "APP_CONFIG"},
	}

	app.Commands = []cli.Command{
		{
			Name:  "server",
			Usage: "Run the http server",
			Action: func(c *cli.Context) {
				cfg, err := getConfig(c)
				if err != nil {
					log.Fatal(err)
					return
				}

				svc := service.LikeMarkService{}

				if err = svc.Run(cfg); err != nil {
					log.Fatal(err)
				}
			},
		},
		{
			Name:  "migratedb",
			Usage: "Perform database migrations",
			Action: func(c *cli.Context) {
				cfg, err := getConfig(c)
				if err != nil {
					log.Fatal(err)
					return
				}

				svc := service.LikeMarkService{}

				if err = svc.Migrate(cfg); err != nil {
					log.Fatal(err)
				}
			},
		},
	}
	app.Run(os.Args)
}
