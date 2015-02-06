package verbs

type VerbType int

const (
	EnterRoomType VerbType = iota
	LeaveRoomType
	ExitType
	MsgType
	StatusType
	TalkType
)
