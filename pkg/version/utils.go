package version

import "strings"

func TrimVersionPrefix(version string) string {
	return strings.TrimPrefix(version, "v")
}
