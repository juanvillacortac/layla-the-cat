package components

import (
	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
)

var Object = donburi.NewComponentType[resolv.Object]()

func SetObject(entry *donburi.Entry, obj *resolv.Object) {
	Object.Set(entry, obj)
}

func GetObject(entry *donburi.Entry) *resolv.Object {
	return Object.Get(entry)
}
