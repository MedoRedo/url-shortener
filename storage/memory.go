package storage

import "sync"

var urlMap = map[string]string{}
var mu sync.RWMutex

func SaveURL(code, original string) {
	mu.Lock()
	defer mu.Unlock()
	urlMap[code] = original
}

func GetURL(code string) (string, bool) {
	mu.RLock()
	defer mu.RUnlock()
	url, found := urlMap[code]
	return url, found
}
