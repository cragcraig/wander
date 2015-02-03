package env

type VerbType int

const (
	enterRoom VerbType = iota
	leaveRoom
	Talk
)
