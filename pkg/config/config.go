package config

// import "layla/pkg/platform"

type CrtQuality int

const (
	CrtQualityOff CrtQuality = iota
	CrtQualityLow
	CrtQualityHigh
)

type Config struct {
	InputScale float64

	Touch      bool
	CrtQuality CrtQuality
}

func (c *Config) ToggleCrtQuality() {
	c.CrtQuality += 1
	if c.CrtQuality > CrtQualityHigh {
		c.CrtQuality = CrtQualityOff
	}
}

var (
	BaseWidth  = 640 / 3
	BaseHeight = 360 / 3
	Width      = BaseWidth
	Height     = BaseHeight
	Scale      = 3.0

	DataDir = ""
)
var C *Config

func init() {
	// crt := CrtQualityHigh
	// switch platform.Platform() {
	// case platform.Wasm:
	// 	crt = CrtQualityLow
	// case platform.Mobile:
	// 	crt = CrtQualityOff
	// default:
	// }

	C = &Config{
		InputScale: 1,

		// Touch:      platform.Platform() == platform.Mobile,
		Touch: true,
		// CrtQuality: crt,
	}
}
