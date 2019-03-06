package env

import (
	"github.com/cragcraig/wander/verbs"
)

type Action struct {
	Verb       verbs.Verb
	Tool       Interactable // optional
	TargetHint *string      // optional
	Player     *Player      // optional
	Arg       interface{}
}

func (action *Action) GetSpeakTarget(target *Player) verbs.SpeakTarget {
	if action.Player == nil || action.Player.Id != target.Id {
		return verbs.Other
	}
	return verbs.Self
}

func (action *Action) GetSpeakNick(target *Player, capitalize bool) string {
	if action.GetSpeakTarget(target) == verbs.Self {
		if capitalize {
			return "You"
		}
		return "you"
	}
	return target.Nick
}
