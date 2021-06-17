package config

import (
	"io/ioutil"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type (
	ConfigMap struct {
		HTTPData HTTPConfig
	}

	HTTPConfig struct {
		Host           string        `yaml:"host"`
		Port           string        `yaml:"port"`
		ReadTimeout    time.Duration `yaml:"readTimeout"`
		WriteTimeout   time.Duration `yaml:"writeTimeout"`
		MaxHeaderBytes int           `yaml:"maxHeaderBytes"`
	}
)

//getConfig function opens yaml config file, then gets settings into Config struct
func GetConfig(path string) (*ConfigMap, error) {
	config := &ConfigMap{}

	//Opening a config file
	ymlfile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	buffer, _ := ioutil.ReadAll(ymlfile)
	defer ymlfile.Close()

	err = yaml.Unmarshal(buffer, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
