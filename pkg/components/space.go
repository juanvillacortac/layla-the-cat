package components

import (
	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
)

var Space = donburi.NewComponentType[resolv.Space]()

func AddToSpace(space *donburi.Entry, objects ...*donburi.Entry) {
	for _, obj := range objects {
		Space.Get(space).Add(GetObject(obj))
	}
}
