package config

import "github.com/package-based-logger/util"

var (
	LogConf LogConfig
)

type LogConfig struct {
	Level   string         `yaml:"level"`
	Package []PackageLevel `yaml:"package"`
}

type PackageLevel struct {
	Name  string `yaml:"name"`
	Level string `yaml:"level"`
}

func InitLogConfig() {
	err := util.YamlReader(`./config/log_config.yaml`, &LogConf)
	if err != nil {
		return
	}
}
