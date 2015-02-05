package env

type Interactable interface {
	GetName() string
	GetHandler(verb *Verb) VerbHandler
	WhatCanThisDo() []*Verb
	DoesMatchHint(hint string) bool
}
