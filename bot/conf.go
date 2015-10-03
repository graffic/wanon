package bot

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// ConfService is the configuration service
type ConfService struct {
	bytes []byte
}

// LoadConf loads the configuration from a file
func LoadConf(fileName string) (*ConfService, error) {
	conf := new(ConfService)

	file, err := os.Open(fileName)
	if err != nil {
		return conf, err
	}

	conf.bytes, err = ioutil.ReadAll(file)
	if err != nil {
		return conf, err
	}

	return conf, nil
}

// Get a configuraiton area
// It unmashals the structure you give from the loaded configuration.
func (service *ConfService) Get(in interface{}) error {
	return yaml.Unmarshal(service.bytes, in)
}
