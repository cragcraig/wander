package verbs

var EnterRoom Verb = Verb{
	Speakable{"entered", "enter", "entering"},
	Speakable{"entered", "enters", "entering"},
	[]string{},
	[]VerbType{EnterRoomType},
	"Enter a room.",
	false,
	nil}

var LeaveRoom Verb = Verb{
	Speakable{"left", "leave", "leaving"},
	Speakable{"left", "leaves", "leaving"},
	[]string{},
	[]VerbType{LeaveRoomType},
	"Leave a room.",
	false,
	nil}

var Exit Verb = Verb{
	Speakable{"exited", "exits", "exiting"},
	Speakable{"exited", "exits", "exiting"},
	[]string{"exit", "logout", "quit"},
	[]VerbType{ExitType},
	"Disconnect from the game.",
	true,
	nil}

var Help Verb = Verb{
	Speakable{"help", "help", "help"},
	Speakable{"help", "help", "help"},
	[]string{"help", "halp"},
	[]VerbType{HelpType},
	"List all available actions.",
	true,
	StringParse}

var Status Verb = Verb{
	Speakable{"status", "status", "status"},
	Speakable{"status", "status", "status"},
	[]string{"status"},
	[]VerbType{StatusType},
	"Get your current status.",
	true,
	nil}

var Msg Verb = Verb{
	Speakable{"messaged", "message", "messaging"},
	Speakable{"messaged", "messages", "messaging"},
	[]string{},
	[]VerbType{MsgType},
	"Send a message to everyone nearby.",
	false,
	StringParse}

var Talk Verb = Verb{
	Speakable{"said", "say", "saying"},
	Speakable{"said", "says", "saying"},
	[]string{"say", "speak", "talk"},
	[]VerbType{TalkType},
	"Speak to everyone nearby.",
	false,
	StringParse}
