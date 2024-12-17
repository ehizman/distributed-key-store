package models

import (
	"fmt"
	"os"
)

type FileTransactionLogger struct {
	events chan <- Event //write only channel for Events
	errors <- chan error // Read only channel for go error
	lastSequence uint64 // The last used event sequence number
	file *os.File // the location of the transaction log
}

func NewFileTransactionLogger(filename string) (TransactionLogger,  error) {
	// open a file in READ, APPEND and WRITE mode
	file, err := os.OpenFile(filename, os.O_RDWR | os.O_APPEND | os.O_CREATE, 0755)

	// return error if you cannot find the file
	if err != nil {
		return nil, fmt.Errorf("cannot open the transaction log file: %w", err)
	}
	// return a new FileTransactionLogger with a reference to the file
	return &FileTransactionLogger{file: file}, nil
}
func (l *FileTransactionLogger) WritePut (key, value string) {
	l.events <- Event{EventType: EventPut, Key: key, Value: value}
}

func (l *FileTransactionLogger) WriteDelete (key string) {
	l.events <- Event{EventType: EventDelete, Key: key}
}

func (l *FileTransactionLogger) Err() <- chan error {
	return l.errors
}

func (l *FileTransactionLogger) Run() {
	// create a standard buffered channel with a capacity of 16 user-defined `Event` model
	events := make(chan Event, 16)

	// create an error channel with a capacity of just 1 error
	errors:= make(chan error, 1)

	// start a go routine
	go func() {
		for e:= range events { // retrieve the next Event
			l.lastSequence++ // increment the sequence number
			_, err := fmt.Fprintf(
				l.file,
				"%d\t%d\t%s\t%s\n",
				l.lastSequence, e.EventType, e.Key, e.Value)

			if err != nil {
				// write to error channel
				errors <- err
				return
			}
		}
	}()
}