package verbs

var EnterRoom Verb = Verb{
	Speakable{"entered", "enters", "entering"},
	[]string{},
	[]VerbType{EnterRoomType},
	false}

var LeaveRoom Verb = Verb{
	Speakable{"left", "leaves", "leaving"},
	[]string{},
	[]VerbType{LeaveRoomType},
	false}

var Exit Verb = Verb{
	Speakable{"exited", "exits", "exiting"},
	[]string{"exit", "logout", "quit"},
	[]VerbType{ExitType},
	true}

var Msg Verb = Verb{
	Speakable{"messaged", "message", "messaging"},
	[]string{},
	[]VerbType{MsgType},
	false}

var Status Verb = Verb{
	Speakable{"status", "status", "status"},
	[]string{"status"},
	[]VerbType{StatusType},
	true}

var Talk Verb = Verb{
	Speakable{"said", "says", "saying"},
	[]string{"say", "speak", "talk"},
	[]VerbType{TalkType},
	false}
