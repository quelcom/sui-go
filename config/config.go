package config

import (
	"fmt"
	"io/ioutil"
	"time"

	toml "github.com/pelletier/go-toml"
)

type Config struct {
	Name      string     `toml:"name"`
	Port      int        `toml:"port"`
	Apps      []App      `toml:"app"`
	Groups    []Group    `toml:"group"`
	Providers []Provider `toml:"provider"`
	Modules   []Module   `toml:"module"`
}

type App struct {
	Name string `toml:"name"`
	URL  string `toml:"url"`
	Icon string `toml:"icon"`
}

type Group struct {
	Name  string `toml:"name"`
	Links []Link `toml:"links"`
}

type Provider struct {
	Name   string `toml:"name"`
	URL    string `toml:"url"`
	Prefix string `toml:"prefix"`
}

type Module struct {
	Name           string            `toml:"name"`
	UpdateInterval time.Duration     `toml:"update_interval"`
	Data           map[string]string `toml:"data"`
	Output         string
}

type Link struct {
	Name string `toml:"name"`
	URL  string `toml:"url"`
}

func ParseFromFile(configFile string) (Config, error) {
	conf, err := ioutil.ReadFile(configFile)
	if err != nil {
		return Config{}, err
	}

	return setConfig(conf)
}

func setConfig(confContents []byte) (conf Config, err error) {
	err = toml.Unmarshal(confContents, &conf)
	if err != nil {
		return Config{}, fmt.Errorf("Could not parse config contents. Error: %v\n", err)
	}

	if conf.Port == 0 {
		return Config{}, fmt.Errorf("Could not read port in configuration file %s\n", confContents)
	}

	return
}
