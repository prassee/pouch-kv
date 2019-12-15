package main
import "C"
import (
	"simpleGoKVStore/kvrepo"
	"sync"
)

var mut sync.RWMutex
var kv = kvrepo.NewKVStore()

func main() {}

//export Put
func Put(key, value string) {
	mut.Lock()
	defer mut.Unlock()
	kv.Put(key, value)
}

//export Delete
func Delete(key string) {
	mut.Lock()
	defer mut.Unlock()
	if kv.HasEntry(key) {
		kv.Pop(key)
	}
}

//export Get
func Get(key string) string {
	mut.RLock()
	defer mut.RUnlock()
	if kv.HasEntry(key) {
		return kv.Get(key)
	}
	return ""
}
