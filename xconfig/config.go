package xconfig

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func Load(config interface{}, files ...string) (err error) {
	data, err := ioutil.ReadFile(files[0])
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, config)
	return
}
