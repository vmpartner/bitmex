package config

import (
	"os"
	"encoding/json"
	"gitlab.com/vitams/bitmex/tools"
)

type Config struct {
	Host   string
	Key    string
	Secret string
}

type MasterConfig struct {
	IsDev  bool
	Master Config
	Dev    Config
}

func LoadConfig(path string) Config {
	config := LoadMasterConfig(path)
	if config.IsDev {
		return config.Dev
	}

	return config.Master
}

func LoadMasterConfig(path string) MasterConfig {
	file, err := os.Open(path)
	tools.CheckErr(err)
	defer file.Close()
	decoder := json.NewDecoder(file)
	config := MasterConfig{}
	err = decoder.Decode(&config)
	tools.CheckErr(err)

	return config
}