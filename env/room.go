package env

import (
	"fmt"
	"github.com/gnarlyskier/wander/core"
)

type Room struct {
	Actions       chan<- Action
	Users         map[int]core.ActiveUser
	interactables []Interactable
}

func (room *Room) IsActive() bool {
	return len(room.Users) > 0
}

func CreateRoom() Room {
	actions := make(chan Action)
	room := Room{actions, map[int]core.ActiveUser{}, []Interactable{BaseRoomInteractable{}}}
	go room.handleActions(actions)
	return room
}

func (room Room) handleActions(actions chan Action) {
	for action := range actions {
		var targets []Interactable
		// Check all interactables in the room as possible targets
		for i := range room.interactables {
			inter := room.interactables[i]
			if action.TargetHint == nil || inter.DoesMatchHint(*action.TargetHint) {
				if inter.GetHandler(action.Verb) != nil {
					targets = append(targets, inter)
				}
			}
		}
		// Dispatch event
		switch len(targets) {
		case 0:
			// Found no targets (error)
			if action.User != nil {
				action.User.Conn.Write <- fmt.Sprintf("No targets to %v\n", action.Verb)
			}
		case 1:
			// Found a single target (sucess)
			if err := targets[0].GetHandler(action.Verb)(room, action, targets[0]); err != nil {
				if action.User != nil {
					action.User.Conn.Write <- err.Error() + "\n"
				}
			}
		default:
			// Found multiple targets (prompt)
			if action.User != nil {
				action.User.Conn.Write <- "Multiple targets:\n"
				for i := range targets {
					action.User.Conn.Write <- targets[i].GetName() + "\n"
				}
			}
		}
		if action.User != nil {
			action.User.Conn.Prompt <- true
		}
	}
}
