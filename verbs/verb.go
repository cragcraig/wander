package verbs

type SpeakTarget int
type SpeakTime int

const (
	Past SpeakTime = iota
	Present
	Continuous
)

const (
	Self SpeakTarget = iota
	Other
)

type Speakable struct {
	Past, Present, Continuous string
}

type Verb struct {
	self           Speakable
	other          Speakable
	CommandAliases []string
	Types          []VerbType
	Targeted       bool
}

func (verb *Verb) Speak(target SpeakTarget, time SpeakTime) string {
	switch target {
	case Self:
		switch time {
		case Past:
			return verb.self.Past
		case Present:
			return verb.self.Present
		case Continuous:
			return verb.self.Continuous
		}
	case Other:
		switch time {
		case Past:
			return verb.other.Past
		case Present:
			return verb.other.Present
		case Continuous:
			return verb.other.Continuous
		}
	}
	return "<verb naming error>"
}

func (verb *Verb) String() string {
	return verb.other.Present
}

func (verb *Verb) HasType(t VerbType) bool {
	for i := range verb.Types {
		if verb.Types[i] == t {
			return true
		}
	}
	return false
}
