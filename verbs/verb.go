package verbs

type Speakable struct {
	Past, Present, Continuous string
}

type Verb struct {
	Speakable
	CommandAliases []string
	Types          []VerbType
	Targeted       bool
}

func (verb *Verb) String() string {
	return verb.Present
}

func (verb *Verb) HasType(t VerbType) bool {
	for i := range verb.Types {
		if verb.Types[i] == t {
			return true
		}
	}
	return false
}
