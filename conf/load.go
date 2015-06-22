package conf

import (
	"github.com/codegangsta/cli"
    "gopkg.in/yaml.v2"
    "os"
    "io/ioutil"
)

// Save the configuration - it's load once
var (
    configCache *ConfigData
)

func LoadConfig(c *cli.Context) (error) {
    if configCache != nil {
        return nil
    }

	yamlPath := c.GlobalString("config")

	if _, err := os.Stat(yamlPath); err != nil {
		return nil
	}

	ymlData, err := ioutil.ReadFile(yamlPath)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal([]byte(ymlData), &Config)
    if err != nil {
        return err
    }

    configCache = &Config

	return nil
}
