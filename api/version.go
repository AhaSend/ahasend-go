package api

import (
	"runtime/debug"
	"strings"
)

// modulePath is this SDK's module path, used to find its own version in the
// build information of the program that imports it.
const modulePath = "github.com/AhaSend/ahasend-go"

// Version is the version of this SDK.
//
// A Go library cannot be stamped at build time the way a binary can: consumers
// compile it from source, so linker flags set during a release never reach
// them. This constant is therefore the source of truth, and the release
// workflow refuses to publish a tag that does not match it.
const Version = "0.1.0"

// UserAgent is the User-Agent header sent when the caller has not set one.
var UserAgent = "AhaSend-Go-SDK/" + resolveVersion()

// resolveVersion reports the version to advertise.
//
// It prefers the module version the Go toolchain recorded in the consuming
// binary, which is exact — it reflects the version actually resolved, even if
// that is a pseudo-version or a version whose constant was never bumped. It
// falls back to Version when there is no build information to read: `go test`
// inside this module, a build with -buildvcs=false, or a consumer using a
// replace directive.
func resolveVersion() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return Version
	}
	if info.Main.Path == modulePath {
		if version := releaseVersion(info.Main.Version); version != "" {
			return version
		}
	}
	for _, dep := range info.Deps {
		if dep == nil || dep.Path != modulePath {
			continue
		}
		// A replaced module reports the replacement's version.
		if dep.Replace != nil {
			dep = dep.Replace
		}
		if version := releaseVersion(dep.Version); version != "" {
			return version
		}
	}
	return Version
}

// releaseVersion normalizes a module version for display, and returns an empty
// string for the placeholders the toolchain uses when there is no real one.
func releaseVersion(version string) string {
	switch version {
	case "", "(devel)", "devel":
		return ""
	}
	return strings.TrimPrefix(version, "v")
}
