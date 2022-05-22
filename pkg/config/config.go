package config

import (
	"io/ioutil"

	"github.com/kanztu/goblog/pkg/logger"
	"gopkg.in/yaml.v2"
)

type GlobalConfig struct {
	DB SqliteConfig `yaml:"db"`
}

type SqliteConfig struct {
	File string `yaml:"file"`
}

var (
	CfgGlobal GlobalConfig
)

func LoadGlobalConfig(path string) {
	log := logger.InitLogger(logger.LEVEL_INFO, "yaml")
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(yamlFile, &CfgGlobal)
	if err != nil {
		log.Fatal(err)
	}
}
