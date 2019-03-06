package env

import (
	"fmt"
	"github.com/cragcraig/wander/verbs"
)

type privatePlayer struct {
	player *Player
}

func (private *privatePlayer) statusHandler(room *Room, action *Action, target Interactable) string {
	return fmt.Sprintf("Your nickname: %v", action.Player)
}

func (private *privatePlayer) exitHandler(room *Room, action *Action, target Interactable) string {
	private.player.Conn.Close()
	return ""
}

func (private *privatePlayer) helpHandler(room *Room, action *Action, target Interactable) string {
	doables := private.WhatCanThisDo()
	private.player.Conn.Write <- "Available actions:"
	for i := range doables {
		if len(doables[i].CommandAliases) != 0 {
			private.player.Conn.Write <- fmt.Sprintf(
				" %-14v - %v",
				doables[i].CommandAliases[0],
				doables[i].Help)
		}
	}
	return ""
}

func (private *privatePlayer) GetName() string {
	return "you"
}

func (private *privatePlayer) GetHandler(verb verbs.Verb) VerbHandler {
	switch {
	case verb.HasType(verbs.StatusType):
		return private.statusHandler
	case verb.HasType(verbs.ExitType):
		return private.exitHandler
	case verb.HasType(verbs.HelpType):
		return private.helpHandler
	}
	return nil
}

func (private *privatePlayer) WhatCanThisDo() []verbs.Verb {
	return []verbs.Verb{verbs.Exit, verbs.Status, verbs.Help, verbs.Talk}
}

func (private *privatePlayer) DoesMatchHint(hint string) bool {
	return hint == "me" || hint == "myself"
}
