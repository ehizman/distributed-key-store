package models

import "sync"

type LockableMap struct {
	sync.RWMutex
	M map[string]string
}
