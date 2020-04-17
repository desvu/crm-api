package game

type Platform uint8

const (
	PlatformUndefined Platform = 0
	PlatformLinux     Platform = 1
	PlatformMacOS     Platform = 2
	PlatformWindows   Platform = 3
	PlatformWeb       Platform = 4
)

func NewPlatform(v uint8) Platform {
	var source Platform

	switch v := Platform(v); v {
	case PlatformLinux,
		PlatformMacOS,
		PlatformWindows,
		PlatformWeb:
		source = v
	}

	return source
}

func NewPlatformByString(v string) Platform {
	switch v {
	case "linux":
		return PlatformLinux
	case "macOS":
		return PlatformMacOS
	case "windows":
		return PlatformWindows
	case "web":
		return PlatformWeb
	default:
		return PlatformUndefined
	}
}

func (p Platform) Value() uint8 {
	return uint8(p)
}

func (p Platform) String() string {
	switch p {
	case PlatformLinux:
		return "linux"
	case PlatformMacOS:
		return "macOS"
	case PlatformWindows:
		return "windows"
	case PlatformWeb:
		return "web"
	default:
		return "undefined"
	}
}
