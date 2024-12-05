package main

import (
	"errors"
	"fmt"
)

var  store = make(map[string] string)
var ErrorNoSuchKey = errors.New("no such key")

func Put(key, value string) error {
	store [key] = value
	return nil
}

func Get(key string) (string, error) {
	value, error := store[key]

	if !error {
		return "", ErrorNoSuchKey
	}
	return value, nil
}

func main(){
	Put("five", "5")

	fmt.Println(store)
}