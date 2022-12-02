package configs

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var ConfigFile *Config

//Config
type Config struct {
	ConfigFile *ini.File
	ServerPort int
	ServerHost string
}

func init() {
	ConfigFile = new(Config)
	ConfigFile.Load()
}
func NewCfg() *Config {
	return ConfigFile
}

func (cfg *Config) Load() error {
	file, err := ini.Load("default.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		return err
	}
	cfg.ConfigFile = file
	server := file.Section("server")
	cfg.ServerPort = server.Key("port").MustInt(8080)
	cfg.ServerHost = server.Key("host").MustString("192.168.1.108")
	return nil
}

func GetConfigs() *Config {
	return ConfigFile
}
