package env

import (
	"github.com/gnarlyskier/wander/core"
	"github.com/gnarlyskier/wander/verbs"
)

type Player struct {
	core.ActiveUser
	Public  []Interactable
	Private []Interactable
}

func CreatePlayer(user *core.ActiveUser) *Player {
	player := Player{*user, []Interactable{}, []Interactable{}}
	player.Public = append(player.Public, &publicPlayer{&player})
	player.Private = append(player.Private, &privatePlayer{&player})
	return &player
}

func (player *Player) CreateAction(verb verbs.Verb, tool Interactable, targetHint *string, args []string) *Action {
	return &Action{verb, tool, targetHint, player, args}
}
