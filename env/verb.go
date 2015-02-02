package env

type Verb struct {
    Names []String
    Types []VerbType
}

func (verb *Verb) String() string {
    return verb.Names[0]
}

type VerbHandler func(ActionQueue chan<- Action, target *Interactible, verb Verb, tool *Interactable, origin *Interactable, msgwrite chan<- string)
