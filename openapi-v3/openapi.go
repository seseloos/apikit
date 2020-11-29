package openapi3

import (
	"errors"
	"fmt"
)

type OpenAPI struct {
	OpenAPI      string        `json:"openapi" yaml:"openapi"`
	Info         Info          `json:"info" yaml:"info"`
	Servers      []interface{} `json:"servers,omitempty" yaml:"servers,omitempty"`
	Paths        interface{}   `json:"paths" yaml:"paths"`
	Components   interface{}   `json:"components,omitempty" yaml:"components,omitempty"`
	Security     []interface{} `json:"security,omitempty" yaml:"security,omitempty"`
	Tags         []interface{} `json:"tags,omitempty" yaml:"tags,omitempty"`
	ExternalDocs interface{}   `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
}

func (api *OpenAPI) Validate() error {

	if api.OpenAPI == "" {
		return errors.New("openapi.OpenAPI is required and must be a non-empty string")
	}

	if err := api.Info.Validate(); err != nil {
		return fmt.Errorf("invalid openapi.Info: %w", err)
	}

	if api.Paths == nil {
		return errors.New("openapi.Path is required")
	}

	return nil
}
