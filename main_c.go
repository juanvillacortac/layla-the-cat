//go:build nintendosdk || microsoftgdk
// +build nintendosdk microsoftgdk

package main

import "C"

//export LaylaMain
func LaylaMain() {
	main()
}
