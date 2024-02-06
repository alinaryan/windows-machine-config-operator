package windows

type ServerVersion string

const (
	// Server2019 represent Windows Server 2019
	Server2019 ServerVersion = "2019"
	// Server2022 represent Windows Server 2022
	Server2022 ServerVersion = "2022"
)

// SupportedVersions are the Windows Server versions supported by the e2e test.
// "" implies the default which is Server2022
var SupportedVersions = []ServerVersion{Server2019, Server2022, ""}

// BuildNumber returns the build for a given server version as defined by Microsoft
var BuildNumber = map[ServerVersion]string{
	Server2019: "10.0.17763",
	Server2022: "10.0.20348",
}

// IsSupported checks if the given version is supported
func IsSupported(version ServerVersion) bool {
	for _, v := range SupportedVersions {
		if v == version {
			return true
		}
	}
	return false
}
