package platform

type PlatformType string

const (
	Desktop PlatformType = "desktop"
	Wasm    PlatformType = "wasm"
	Mobile  PlatformType = "mobile"
)
