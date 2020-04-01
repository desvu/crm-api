package game

type Platform uint8

const (
	PlatformLinux   Platform = 1
	PlatformMacOS   Platform = 2
	PlatformWindows Platform = 3
	PlatformWeb     Platform = 4
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

func (p Platform) Value() uint8 {
	return uint8(p)
}

func (p Platform) String() string {
	var str string

	switch p {
	case PlatformLinux:
		str = "linux"
	case PlatformMacOS:
		str = "mac_os"
	case PlatformWindows:
		str = "windows"
	case PlatformWeb:
		str = "web"
	}

	return str
}
