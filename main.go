package main

import "C"
import (
	"simpleGoKVStore/kvrepo"
	"sync"
)

var mut sync.RWMutex
var kv = kvrepo.NewKVStore()

func main() {}

//Put - adds a key value entry in to store
func Put(key, value string) {
	mut.Lock()
	defer mut.Unlock()
	kv.Put(key, value)
}

//Delete - removes key from store
func Delete(key string) {
	mut.Lock()
	defer mut.Unlock()
	if kv.HasEntry(key) {
		kv.Pop(key)
	}
}

//Get - finds the value of given key
func Get(key string) string {
	mut.RLock()
	defer mut.RUnlock()
	if kv.HasEntry(key) {
		return kv.Get(key)
	}
	return ""
}
