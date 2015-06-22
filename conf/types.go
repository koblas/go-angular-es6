package conf

type AppConfig struct {
    HttpPort    int     `yaml:"http_port"`
    Database    string  `yaml:"database"`
}

type ConfigData struct {
    App       AppConfig  `yaml:"APP"`
}

var (
    Config  ConfigData = ConfigData{ App: AppConfig { HttpPort:   3000 } }
)
