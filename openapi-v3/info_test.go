package openapi3_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	openapi3 "github.com/ExperienceOne/apikit/openapi-v3"
)

func TestInfo_Validate(t *testing.T) {

	testCases := []struct {
		name string
		info openapi3.Info
		want string
	}{
		{
			name: "invalid version",
			info: openapi3.Info{
				Title:          "title",
				Description:    nil,
				TermsOfService: nil,
				Contact:        nil,
				License:        nil,
				Version:        "",
			},
			want: "info.version is required and must be a non-empty string",
		},
		{
			name: "invalid license",
			info: openapi3.Info{
				Title:          "title",
				Description:    nil,
				TermsOfService: nil,
				Contact:        nil,
				License: &openapi3.License{
					Name: "",
					URL:  "",
				},
				Version: "version",
			},
			want: "invalid info.license: license.name is required and must be a non-empty string",
		},
		{
			name: "valid info",
			info: openapi3.Info{
				Title:          "title",
				Description:    nil,
				TermsOfService: nil,
				Contact:        nil,
				License: &openapi3.License{
					Name: "name",
					URL:  "",
				},
				Version: "version",
			},
			want: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			err := tc.info.Validate()
			if tc.want != "" {
				require.EqualError(t, err, tc.want)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
