package openapi3_test

import (
	"testing"

	openapi3 "github.com/ExperienceOne/apikit/openapi-v3"
	"github.com/stretchr/testify/require"
)

func TestOpenAPI_Validate(t *testing.T) {

	testCases := []struct {
		name string
		api  openapi3.OpenAPI
		want string
	}{
		{
			name: "openapi required",
			api: openapi3.OpenAPI{
				OpenAPI:      "",
				Info:         openapi3.Info{},
				Servers:      nil,
				Paths:        nil,
				Components:   nil,
				Security:     nil,
				Tags:         nil,
				ExternalDocs: nil,
			},
			want: "openapi.OpenAPI is required and must be a non-empty string",
		},
		{
			name: "invalid info",
			api: openapi3.OpenAPI{
				OpenAPI: "3.0.0",
				Info: openapi3.Info{
					Title:          "",
					Description:    nil,
					TermsOfService: nil,
					Contact:        nil,
					License:        nil,
					Version:        "",
				},
				Servers:      nil,
				Paths:        nil,
				Components:   nil,
				Security:     nil,
				Tags:         nil,
				ExternalDocs: nil,
			},
			want: "invalid openapi.Info: info.title is required and must be a non-empty string",
		},
		{
			name: "paths required",
			api: openapi3.OpenAPI{
				OpenAPI: "3.0.0",
				Info: openapi3.Info{
					Title:          "title",
					Description:    nil,
					TermsOfService: nil,
					Contact:        nil,
					License:        nil,
					Version:        "version",
				},
				Servers:      nil,
				Paths:        nil,
				Components:   nil,
				Security:     nil,
				Tags:         nil,
				ExternalDocs: nil,
			},
			want: "openapi.Path is required",
		},
		{
			name: "valid spec",
			api: openapi3.OpenAPI{
				OpenAPI: "3.0.0",
				Info: openapi3.Info{
					Title:          "title",
					Description:    nil,
					TermsOfService: nil,
					Contact:        nil,
					License:        nil,
					Version:        "version",
				},
				Servers:      nil,
				Paths:        "paths",
				Components:   nil,
				Security:     nil,
				Tags:         nil,
				ExternalDocs: nil,
			},
			want: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			err := tc.api.Validate()
			if tc.want != "" {
				require.EqualError(t, err, tc.want)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
