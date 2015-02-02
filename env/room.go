package env

import "fmt"

type Action struct {
    Verb Verb
    Tool *Interactable
    Origin *Interactable
    TargetHint *string
    MsgWrite chan<- string
    ActionHandledSignal core.Signal
}

type Room struct {
    Room Interactable
    Interactables []Interactable
    ActivePlayers int
}

func (room *Room) IsActive() bool {
    return room.activePlayers > 0
}

func (room *Room) GetHandler(verb Verb) VerbHandler {
    // Handle players entering and leaving room
}

func (room *Room) StartRoom() chan<- Action {
    actions := make(chan Action)
    go room.runActionHandler(actions)
    return actions
}

func (room *Room) runActionHandler(actions chan Action) {
    for action := range actions {
        var targets []Interactables
        // Check all interactables in the room as possible targets
        for i := range room.Interactables {
            inter := room.Interactables[i]
            if action.TargetHint == nil || inter.DoesMatchHint(*action.TargetHint) {
                if inter.GetHandler(action.Verb) != nil {
                    append(targets, inter)
                }
            }
        }
        // Check room as possible target
        if action.TargetHint == nil && room.GetHandler(action.Verb) != nil {
            append(targets, inter)
        }
        // Dispatch event
        switch len(targets) {
        case 0:
            // Found no targets (error)
            if action.MsgWrite != nil {
                action.MsgWrite <- fmt.Sprintf("No targets to %v\n", action.Verb)
            }
        case 1:
            // Found a single target (sucess)
            GetHandler(action.Verb)(actions, targets[0], action.Verb, action.Tool, action.Origin, action.MsgWrite)
        default:
            // Found multiple targets (prompt)
            if action.MsgWrite != nil {
                action.MsgWrite <- "Multiple targets:\n"
                for i := range targets {
                    action.MsgWrite <- inter.GetName() + "\n"
                }
            }
        }
        if action.FinishedSignal != nil {
            action.FinishedSignal.Signal()
        }
    }
}
