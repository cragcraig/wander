package env

import (
	"github.com/cragcraig/wander/verbs"
)

type Interactable interface {
	GetName() string
	GetHandler(verb verbs.Verb) VerbHandler
	WhatCanThisDo() []verbs.Verb
	DoesMatchHint(hint string) bool
}

type VerbHandler func(room *Room, action *Action, target Interactable) string
