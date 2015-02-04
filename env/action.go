package env

import (
	"github.com/gnarlyskier/wander/core"
)

type Action struct {
	Verb       *Verb
	Tool       Interactable // optional
	TargetHint *string // optional
	User       *core.ActiveUser // optional
	Args       []string
}
