package config

import (
	"github.com/teixie/xframework/fs"
	"gopkg.in/yaml.v2"
)

func LoadFromYamlFile(cfg interface{}, filename string) (err error) {
	data, err := fs.GetFileContents(filename)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(data, cfg)
	return
}
