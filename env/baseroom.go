package env

import (
	"fmt"
	"github.com/cragcraig/wander/verbs"
)

type BaseRoom struct{}

func (inter BaseRoom) GetName() string {
	return "<base room>"
}

func (inter *BaseRoom) GetHandler(verb verbs.Verb) VerbHandler {
	switch {
	case verb.HasType(verbs.EnterRoomType):
		// TODO(craig): Don't inline function.
		return func(room *Room, action *Action, target Interactable) string {
			room.Players[action.Player.Id] = *action.Player
			go func() {
				room.Actions <- action.Player.CreateAction(
					verbs.Msg, nil, nil, fmt.Sprintf("%v entered", action.Player))
			}()
			return ""
		}
	case verb.HasType(verbs.LeaveRoomType):
		// TODO(craig): Don't inline function.
		return func(room *Room, action *Action, target Interactable) string {
			delete(room.Players, action.Player.Id)
			go func() {
				room.Actions <- action.Player.CreateAction(
					verbs.Msg, nil, nil, fmt.Sprintf("%v left", action.Player))
			}()
			return ""
		}
	}
	return nil
}

func (inter *BaseRoom) DoesMatchHint(hint string) bool {
	return false
}

func (inter *BaseRoom) WhatCanThisDo() []verbs.Verb {
	return []verbs.Verb{}
}
