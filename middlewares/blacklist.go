package middlewares

import (
	"sync"
)

var (
	blacklistedTokens = make(map[string]struct{})
	blacklistMutex    sync.RWMutex
)

func AddToBlacklist(tokenString string) {
	blacklistMutex.Lock()
	defer blacklistMutex.Unlock()
	blacklistedTokens[tokenString] = struct{}{}
}

func IsTokenBlacklisted(tokenString string) bool {
	blacklistMutex.RLock()
	defer blacklistMutex.RUnlock()
	_, exists := blacklistedTokens[tokenString]
	return exists
}
