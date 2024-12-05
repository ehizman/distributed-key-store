package handlers

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"github.com/ehizman/key-value-store/models"
	"github.com/gorilla/mux"
)

var store = models. LockableMap {
	M: make(map[string] string),
}
var ErrorNoSuchKey = errors.New("no such key")

func PutHandlerFunc(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	value, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	err = put(key, string(value))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func GetHandlerFunc(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	value, err := get(key)

	if errors.Is(err, ErrorNoSuchKey) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, value)
}

func DeleteFuncHandler(w http.ResponseWriter,  r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	deleteFunc(key)
	fmt.Fprintln(w, fmt.Sprintf("%v has been successfully deleted", key))
	w.WriteHeader(http.StatusOK)

}

func put(key, value string) error {
	store.Lock()
	defer store.Unlock()

	store.M[key] = value
	return nil
}

func get(key string) (string, error) {
	store.RLock()
	defer store.RUnlock()

	value, ok := store.M[key]

	if !ok {
		return "", ErrorNoSuchKey
	}

	return value, nil
}

func deleteFunc (key string)  {
	store.Lock()
	defer store.Unlock()

	delete(store.M, key)	
}