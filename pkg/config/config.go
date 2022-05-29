package config

import (
	"io/ioutil"

	"github.com/kanztu/goblog/pkg/logger"
	"gopkg.in/yaml.v2"
)

type GlobalConfig struct {
	DB   SqliteConfig `yaml:"db"`
	Stor Storage      `yaml:"storage"`
}

type SqliteConfig struct {
	File string `yaml:"file"`
}

type Storage struct {
	Blog string `yaml:"blog"`
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
