/*
Package kvrepo - the simple map based backend
*/
package kvrepo

import (
	"encoding/gob"
	"os"
)

var pouchFilePath = ".pouchfile.db"

type kvrepo struct {
	kvStore map[string]string
}

func NewKVStore() kvrepo {
	_, err := os.Stat(pouchFilePath)
	if os.IsNotExist(err) {
		return kvrepo{kvStore: make(map[string]string)}
	} else {
		return readFromFile()
	}
}

func (kvs kvrepo) Put(key, value string) {
	kvs.kvStore[key] = value
}

func (kvs kvrepo) Get(key string) string {
	return kvs.kvStore[key]
}

func (kvs kvrepo) HasEntry(key string) bool {
	if _, ok := kvs.kvStore[key]; ok {
		return ok
	}
	return false
}

func (kvs kvrepo) Pop(key string) {
	delete(kvs.kvStore, key)
}

func (kvs kvrepo) writeToFile() {
	pouchfile, err := os.Create(pouchFilePath)
	if err != nil {
		panic("can't write to file")
	}
	dataEncoder := gob.NewEncoder(pouchfile)
	dataEncoder.Encode(kvs)
	pouchfile.Close()
}

func readFromFile() kvrepo {
	kvs := kvrepo{kvStore: make(map[string]string)}
	pouchfile, err := os.Open(pouchFilePath)
	if err != nil {
		panic("cannot read from file")
	}

	dataDecoder := gob.NewDecoder(pouchfile)
	dataDecoder.Decode(&kvs)
	pouchfile.Close()
	return kvs
}
