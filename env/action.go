package env

import (
	"github.com/gnarlyskier/wander/core"
)

type Action struct {
	Verb       Verb
	Tool       *Interactable
	TargetHint *string
	User       *core.ActiveUser
	Args       []string
}
