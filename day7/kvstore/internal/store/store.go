package store

var kv = make(map[string]string)

func Set(key, value string) {
	kv[key] = value
}

func Get(key string) (string, bool) {
	value, ok := kv[key]
	return value, ok
}
