//go:build (android || ios || (darwin && arm) || (darwin && arm64)) && !js

package platform

func Platform() PlatformType {
	return Mobile
}
