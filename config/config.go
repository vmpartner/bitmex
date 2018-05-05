package config

import (
	"encoding/json"
	"github.com/vmpartner/bitmex/tools"
	"os"
)

type Config struct {
	Host    string
	Key     string
	Secret  string
	Timeout int64
	DB      struct {
		Host     string
		Login    string
		Password string
		Name     string
	}
	Neural struct {
		Iterations int
		Predict    float64
	}
	Strategy struct {
		Profit   float64
		StopLose float64
		Quantity float32
	}
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
