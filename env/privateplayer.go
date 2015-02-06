package env

import (
	"fmt"
	"github.com/gnarlyskier/wander/verbs"
)

type privatePlayer struct {
	player *Player
}

func (public *privatePlayer) statusHandler(room *Room, action *Action, target Interactable) string {
	return fmt.Sprintf("Your nickname: %v", action.Player)
}

func (public *privatePlayer) exitHandler(room *Room, action *Action, target Interactable) string {
	public.player.Conn.Close()
	return ""
}

func (public *privatePlayer) GetName() string {
	return "you"
}

func (public *privatePlayer) GetHandler(verb verbs.Verb) VerbHandler {
	switch {
	case verb.HasType(verbs.StatusType):
		return public.statusHandler
	case verb.HasType(verbs.ExitType):
		return public.exitHandler
	}
	return nil
}

func (public *privatePlayer) WhatCanThisDo() []verbs.Verb {
	return []verbs.Verb{verbs.Exit, verbs.Status, verbs.Talk}
}

func (public *privatePlayer) DoesMatchHint(hint string) bool {
	return hint == "me" || hint == "myself"
}
