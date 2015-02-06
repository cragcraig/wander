package env

import (
	"fmt"
)

type Room struct {
	Actions       chan<- *Action
	Players       map[int]Player
	interactables []Interactable
}

func (room *Room) IsActive() bool {
	return len(room.Players) > 0
}

func CreateRoom() *Room {
	actions := make(chan *Action)
	room := Room{actions, map[int]Player{}, []Interactable{&BaseRoomInteractable{}}}
	go room.handleActions(actions)
	return &room
}

func getTargets(interactables []Interactable, action *Action) []Interactable {
	var targets []Interactable
	// Check all interactables as possible targets
	for i := range interactables {
		inter := interactables[i]
		if action.TargetHint == nil || inter.DoesMatchHint(*action.TargetHint) {
			if inter.GetHandler(action.Verb) != nil {
				targets = append(targets, inter)
			}
		}
	}
	return targets
}

func (room *Room) handleActions(actions chan *Action) {
	for action := range actions {
		// Get all targets
		targets := getTargets(room.interactables, action)
		for i := range room.Players {
			targets = append(targets, getTargets(room.Players[i].Public, action)...)
		}
		if action.Player != nil {
			targets = append(targets, getTargets(action.Player.Private, action)...)
		}
		// Dispatch event
		switch {
		case len(targets) == 0:
			// Found no targets (error)
			if action.Player != nil {
				action.Player.Conn.Write <- fmt.Sprintf("No valid targets for %v", action.Verb)
			}
		case len(targets) == 1 || !action.Verb.Targeted:
			// Success
			for i := range targets {
				if msg := targets[i].GetHandler(action.Verb)(room, action, targets[i]); msg != "" && action.Player != nil {
					action.Player.Conn.Write <- msg
				}
			}
		default:
			// Found multiple targets (prompt)
			if action.Player != nil {
				action.Player.Conn.Write <- "Multiple targets:"
				for i := range targets {
					action.Player.Conn.Write <- targets[i].GetName()
				}
			}
		}
		if action.Player != nil {
			action.Player.Conn.Prompt <- true
		}
	}
}
