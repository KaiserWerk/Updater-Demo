package assets

import _ "embed"

//go:embed versionFiles/launcher.txt
var launcherVersion string

//go:embed versionFiles/app.txt
var appVersion string

func GetLauncherVersion() string {
	return launcherVersion
}

func GetAppVersion() string {
	return appVersion
}