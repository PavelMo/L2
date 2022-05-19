package cfg

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Config struct {
	Port string `yaml:"port"`
}

func NewCFG() *Config {
	return &Config{}
}
func (c *Config) GetCfg() (string, error) {
	yml, err := ioutil.ReadFile("/home/paul/GolandProjects/L2Git/develop/dev11/cfg/config.yml")
	if err != nil {
		return "", err
	}
	err = yaml.Unmarshal(yml, c)
	if err != nil {
		return "", err
	}
	return c.Port, nil
}
