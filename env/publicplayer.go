package env

import (
	"fmt"
	"github.com/gnarlyskier/wander/verbs"
)

type publicPlayer struct {
	player *Player
}

func (public *publicPlayer) talkHandler(room *Room, action *Action, target Interactable) string {
	if len(action.Args) != 1 {
		return "You didn't say anything..."
	} else if action.Player == nil {
		return "Speaking requires an originating user"
	}
	public.player.Conn.Write <- fmt.Sprintf(
		"%v %v: %v",
		action.GetSpeakNick(public.player, true),
		action.Verb.Speak(action.GetSpeakTarget(public.player), verbs.Present),
		action.Args[0])
	return ""
}

func (public *publicPlayer) msgHandler(room *Room, action *Action, target Interactable) string {
	if len(action.Args) != 1 {
		return "msg requires exactly one argument"
	}
	public.player.Conn.Write <- action.Args[0]
	return ""
}

func (public *publicPlayer) GetName() string {
	return public.player.String()
}

func (public *publicPlayer) GetHandler(verb verbs.Verb) VerbHandler {
	switch {
	case verb.HasType(verbs.TalkType):
		return public.talkHandler
	case verb.HasType(verbs.MsgType):
		return public.msgHandler
	}
	return nil
}

func (public *publicPlayer) WhatCanThisDo() []verbs.Verb {
	return []verbs.Verb{}
}

func (public *publicPlayer) DoesMatchHint(hint string) bool {
	return hint == public.player.Nick
}
