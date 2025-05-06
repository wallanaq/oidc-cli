package version

var (
	current  string = "vX.Y.Z-build"
	template string = "oidc-cli/{{.Version}}\n"
)

func Get() string {
	return current
}

func Template() string {
	return template
}
