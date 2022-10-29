package configs

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var ConfigFile *ini.File

//Config
type Config struct {
	ConfigFile *ini.File
	ServerPort int
	ServerHost string
}

func NewCfg() *Config {
	return &Config{
		ConfigFile: ini.Empty(),
	}
}
func (cfg *Config) readServerSettings(iniFile *ini.File) error {
	server := iniFile.Section("server")
	cfg.ServerPort = server.Key("port").MustInt(8080)
	cfg.ServerHost = server.Key("host").MustString("localhost")
	return nil
}

func (cfg *Config) Load() error {
	file, err := ini.Load("default.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		return err
	}
	cfg.ConfigFile = file
	ConfigFile = file
	cfg.readServerSettings(file)
	return nil
}

func GetConfigs() *ini.File {
	return ConfigFile
}
