package openapi3

import (
	"errors"
	"fmt"
)

type Info struct {
	Title          string   `json:"title" yaml:"title"`
	Description    *string  `json:"description,omitempty" yaml:"description,omitempty"`
	TermsOfService *string  `json:"termsOfService,omitempty" yaml:"termsOfService,omitempty"`
	Contact        *Contact `json:"contact,omitempty" yaml:"contact,omitempty"`
	License        *License `json:"license,omitempty" yaml:"license,omitempty"`
	Version        string   `json:"version" yaml:"version"`
}

func (i *Info) Validate() error {

	if i.Title == "" {
		return errors.New("info.title is required and must be a non-empty string")
	}

	if i.License != nil {
		if err := i.License.Validate(); err != nil {
			return fmt.Errorf("invalid info.license: %w", err)
		}
	}

	if i.Version == "" {
		return errors.New("info.version is required and must be a non-empty string")
	}

	return nil
}

type Contact struct {
	Name  string `json:"name,omitempty" yaml:"name,omitempty"`
	URL   string
	Email string
}

type License struct {
	Name string `json:"name" yaml:"name"`
	URL  string `json:"url,omitempty" yaml:"url,omitempty"`
}

func (l *License) Validate() error {

	if l.Name == "" {
		return errors.New("license.name is required and must be a non-empty string")
	}

	return nil
}
