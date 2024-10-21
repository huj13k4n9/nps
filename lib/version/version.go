package version

import "strings"

const VERSION = "0.27.0"

func CheckVersionRequirement(ver string) bool {
	currentVersion := VERSION
	versionPrefixOfInput := ver[:strings.LastIndexByte(ver, '.')]
	versionPrefix := currentVersion[:strings.LastIndexByte(currentVersion, '.')]

	return versionPrefix == versionPrefixOfInput
}

func MinVersion() string {
	currentVersion := VERSION
	versionPrefix := currentVersion[:strings.LastIndexByte(currentVersion, '.')]
	return versionPrefix + ".0"
}
