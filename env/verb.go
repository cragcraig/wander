package env

import (
	"github.com/gnarlyskier/wander/core"
)

type Verb struct {
	Name  string
	Types []VerbType
}

func (verb Verb) CreateUserAction(user core.ActiveUser, tool *Interactable, targetHint *string, args []string) Action {
	return Action{verb, tool, targetHint, &user, args}
}

func (verb Verb) String() string {
	return verb.Name
}

func (verb Verb) HasType(t VerbType) bool {
	for i := range verb.Types {
		if verb.Types[i] == t {
			return true
		}
	}
	return false
}

type VerbHandler func(room Room, action Action, target Interactable) error