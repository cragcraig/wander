package simpleverbs

import "github.com/gnarlyskier/wander/env"

var Talk env.Verb = env.Verb{
    env.Speakable{"said", "says", "saying"},
    []string{"say", "speak", "talk"},
    []env.VerbType{env.Talk}}
