package api

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestVersionConstant(t *testing.T) {
	// The release workflow compares this constant against the tag, so it has to
	// stay in the shape a tag can hold.
	assert.Regexp(t, regexp.MustCompile(`^[0-9]+\.[0-9]+\.[0-9]+$`), Version)
}

func TestUserAgentCarriesTheVersion(t *testing.T) {
	require.NotEmpty(t, UserAgent)
	assert.Regexp(t, regexp.MustCompile(`^AhaSend-Go-SDK/[0-9]+\.[0-9]+\.[0-9]+`), UserAgent)

	// Tests run inside this module, where the toolchain records no release
	// version, so the constant is what gets advertised here.
	assert.Equal(t, "AhaSend-Go-SDK/"+Version, UserAgent)
}

func TestReleaseVersionIgnoresPlaceholders(t *testing.T) {
	testCases := []struct {
		name    string
		version string
		want    string
	}{
		{name: "tagged release", version: "v1.2.3", want: "1.2.3"},
		{name: "pseudo-version", version: "v0.0.0-20260101000000-abcdefabcdef", want: "0.0.0-20260101000000-abcdefabcdef"},
		{name: "development build", version: "(devel)", want: ""},
		{name: "unset", version: "", want: ""},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, releaseVersion(tc.version))
		})
	}
}

func TestDefaultUserAgentIsTheResolvedOne(t *testing.T) {
	assert.Equal(t, UserAgent, GetDefaults().DefaultUserAgent)
}
