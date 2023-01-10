package config

import "layla/pkg/platform"

type Config struct {
	BaseWidth  int
	BaseHeight int

	Width  int
	Height int

	Scale float64

	Touch     bool
	CrtShader bool
}

var C *Config

func init() {
	C = &Config{
		BaseWidth:  640 / 2,
		BaseHeight: 360 / 2,

		Width:  640 / 2,
		Height: 360 / 2,

		Scale: 2,

		Touch:     platform.Platform() == platform.Mobile,
		CrtShader: platform.Platform() != platform.Mobile,
	}
}
