package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

func GetConfig() (*Config, error) {
	var conf Config

	if _, err := toml.DecodeFile("config.toml", &conf); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &conf, nil

}