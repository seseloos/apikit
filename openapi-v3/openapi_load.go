package openapi3

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func SpecFromFile(path string) (*OpenAPI, error) {

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	api := OpenAPI{}
	if err := yaml.Unmarshal(data, &api); err != nil {
		return nil, err
	}

	return &api, nil
}
