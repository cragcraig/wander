package env

import (
	"github.com/gnarlyskier/wander/verbs"
)

type Action struct {
	Verb       verbs.Verb
	Tool       Interactable // optional
	TargetHint *string      // optional
	Player     *Player      // optional
	Args       []string
}
