package main

import (
	"fmt"

	"github.com/pelletier/go-toml"
)

type Tag struct {
	Action string
	Urn    string
}

type VCenter struct {
	Password string
	server   string
	user     string
}

type Config struct {
	Tag     Tag
	VCenter VCenter
}

func main() {
	cfg := loadTomlCfg()

	fmt.Println(cfg.Tag.Action)
	fmt.Println(cfg.VCenter.Password)
}

func loadTomlCfg() *Config {
	config, err := toml.LoadFile("vcconfig")

	if err != nil {
		fmt.Println(err)
	}

	cfg := Config{}

	config.Unmarshal(&cfg)
	config.Unmarshal(&cfg)

	return &cfg
}
