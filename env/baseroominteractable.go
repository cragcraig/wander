package env

import (
	"fmt"
	"github.com/gnarlyskier/wander/verbs"
)

type BaseRoomInteractable struct{}

func (inter BaseRoomInteractable) GetName() string {
	return "<base room>"
}

func (inter *BaseRoomInteractable) GetHandler(verb verbs.Verb) VerbHandler {
	switch {
	case verb.HasType(verbs.EnterRoomType):
		// TODO(craig): Don't inline function.
		return func(room *Room, action *Action, target Interactable) string {
			room.Players[action.Player.Id] = *action.Player
			go func() {
				room.Actions <- action.Player.CreateAction(
					verbs.Msg, nil, nil, []string{fmt.Sprintf("%v entered", action.Player)})
			}()
			return ""
		}
	case verb.HasType(verbs.LeaveRoomType):
		// TODO(craig): Don't inline function.
		return func(room *Room, action *Action, target Interactable) string {
			delete(room.Players, action.Player.Id)
			go func() {
				room.Actions <- action.Player.CreateAction(
					verbs.Msg, nil, nil, []string{fmt.Sprintf("%v left", action.Player)})
			}()
			return ""
		}
	}
	return nil
}

func (inter *BaseRoomInteractable) DoesMatchHint(hint string) bool {
	return false
}

func (inter *BaseRoomInteractable) WhatCanThisDo() []verbs.Verb {
	return []verbs.Verb{}
}
