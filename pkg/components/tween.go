package components

import (
	"github.com/tanema/gween"
	"github.com/yohamta/donburi"
)

var Tween = donburi.NewComponentType[gween.Tween]()
var TweenSeq = donburi.NewComponentType[gween.Sequence]()
