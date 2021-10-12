package configs

import (
	"fmt"
	"os"
	"server_explorer/pkg/models"

	"gopkg.in/ini.v1"
)

//Config
type Config struct {
	ConfigFile *ini.File
	Server     *models.Server
}

func Init() {
	cfg := &Config{}
	file, err := ini.Load("default.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	cfg.ConfigFile = file
	cfg.Server.Port = cfg.ConfigFile.Section("server").Key("port").MustInt(8080)
	fmt.Println("this is port - ", cfg.Server.Port)
}

func (cfg *Config) GetConfigs() *Config {
	return cfg
}
