/*
Package kvrepo - the simple map based backend
*/
package kvrepo

type kvrepo struct {
	kvStore map[string]string
}

func NewKVStore() kvrepo {
	return kvrepo{kvStore: make(map[string]string)}
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

func (kvs kvrepo) WriteToFile() {
}

func (kvs kvrepo) ReadFromFile() {
}
