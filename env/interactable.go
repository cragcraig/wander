package env

type Interactable interface {
	GetName() string
	GetHandler(verb Verb) VerbHandler
	DoesMatchHint(hint string) bool
}
