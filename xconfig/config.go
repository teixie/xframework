package xconfig

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func Load(config interface{}, file string) (err error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, config)
	return
}
