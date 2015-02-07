package verbs

type VerbType int

const (
	EnterRoomType VerbType = iota
	LeaveRoomType
	ExitType
	HelpType
	MsgType
	StatusType
	TalkType
)
