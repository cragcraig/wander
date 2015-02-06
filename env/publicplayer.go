package env

import (
	"fmt"
	"github.com/gnarlyskier/wander/verbs"
)

type publicPlayer struct {
	player *Player
}

func (public *publicPlayer) talkHandler(room *Room, action *Action, target Interactable) string {
    if action.Player == nil {
        return "Speaking requires an originating user"
    }
    if msg, ok := action.Arg.(string); ok && msg != "" {
        public.player.Conn.Write <- fmt.Sprintf(
            "%v %v: %v",
            action.GetSpeakNick(public.player, true),
            action.Verb.Speak(action.GetSpeakTarget(public.player), verbs.Present),
            msg)
        return ""
    }
    return "You didn't say anything..."
}

func (public *publicPlayer) msgHandler(room *Room, action *Action, target Interactable) string {
    if msg, ok := action.Arg.(string); ok {
	    public.player.Conn.Write <- msg
        return ""
    }
	return "A non-string message was sent???"
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
	return hint == public.player.GetTargetHint()
}
