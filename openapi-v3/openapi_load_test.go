package openapi3_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	openapi3 "github.com/ExperienceOne/apikit/openapi-v3"
)

func TestSpecFromFile(t *testing.T) {

	testCases := []struct {
		name    string
		path    string
		wantErr bool
	}{
		{
			name:    "YAML",
			path:    "../testdata/test.openapi.yaml",
			wantErr: false,
		},
		{
			name:    "JSON",
			path:    "../testdata/test.openapi.json",
			wantErr: false,
		},
		{
			name:    "file not existing",
			path:    "../testdata/not-existing-file",
			wantErr: true,
		},
		{
			name:    "invalid spec file format",
			path:    "../testdata/test.openapi.txt",
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			api, err := openapi3.SpecFromFile(tc.path)
			if tc.wantErr {
				require.Error(t, err)
				require.Nil(t, api)
			} else {
				require.NoError(t, err)
				require.NotNil(t, api)
			}
		})
	}
}
