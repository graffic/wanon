package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Configuration struct {
	Token string
}

func Load(fileName string) (*Configuration, error) {
	conf := new(Configuration)

	file, err := os.Open(fileName)
	if err != nil {
		return conf, err
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return conf, err
	}

	err = yaml.Unmarshal(bytes, conf)
	if err != nil {
		return conf, err
	}

	return conf, nil
}
