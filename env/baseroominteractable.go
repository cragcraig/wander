package env

import (
	"errors"
	"fmt"
)

var EnterRoom Verb = Verb{"<enter room>", []VerbType{enterRoom}}
var LeaveRoom Verb = Verb{"<leave room>", []VerbType{leaveRoom}}

type BaseRoomInteractable struct{}

func (inter BaseRoomInteractable) GetName() string {
	return "<base room>"
}

func (inter *BaseRoomInteractable) GetHandler(verb *Verb) VerbHandler {
	switch {
	case verb.HasType(enterRoom):
		// TODO(craig): Don't inline function.
		return func(room *Room, action *Action, target Interactable) error {
			room.Users[action.User.Id] = action.User
			return nil
		}
	case verb.HasType(leaveRoom):
		// TODO(craig): Don't inline function.
		return func(room *Room, action *Action, target Interactable) error {
			delete(room.Users, action.User.Id)
			return nil
		}
	case verb.HasType(Talk):
		// TODO(craig): Don't inline function.
		return func(room *Room, action *Action, target Interactable) error {
			if len(action.Args) != 1 {
				return errors.New("talk requires exactly one argument")
			} else if action.User == nil {
				return errors.New("talk requires an originating user")
			}
			msg := fmt.Sprintf("%v: %v", action.User, action.Args[0])
			for _, v := range room.Users {
				v.Conn.Write <- msg
			}
			return nil
		}
	}
	return nil
}

func (inter *BaseRoomInteractable) DoesMatchHint(hint string) bool {
	return false
}
