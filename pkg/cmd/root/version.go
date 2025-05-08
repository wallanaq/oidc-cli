package root

var (
	current  string = "vX.Y.Z-build"
	template string = "oidc-cli/{{.Version}}\n"
)

func GetVersion() string {
	return current
}

func VersionTemplate() string {
	return template
}
