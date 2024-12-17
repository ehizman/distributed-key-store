package models

type Event struct {
	Sequence uint64
	EventType EventType
	Key string
	Value string
}