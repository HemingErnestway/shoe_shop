package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"regexp"
)

type Server struct {
	Port    string
	Timeout string
}

type Service struct {
	Method string
	Url    string
}

type AccessExceptions struct {
	List []string
}

type Config struct {
	Server
	Api              map[string]Service
	AccessExceptions `yaml:"accessExceptions"`
}

var config Config

func Get() *Config {
	return &config
}

var isEnv = regexp.MustCompile(`\$\{(.*?)}`)

func init() {
	file, err := os.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(file, &config)
	if err != nil {
		panic(err)
	}
}
