package models

type EventType byte

const (
	_  = iota //here iota == 0; ignore the zero value
	EventDelete EventType = iota //iota == 1
	EventPut //iota ==2 ; implicit repeat
)

