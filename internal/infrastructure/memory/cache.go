package memory

import (
	"sync"

	"github.com/rayfiyo/random16type/internal/domain/typecode"
)

// Cache は日付ごとにコードを保存するインメモリ実装。
type Cache struct {
	mu    sync.RWMutex
	store map[string]typecode.Code
}

func NewCache() *Cache {
	return &Cache{
		store: make(map[string]typecode.Code),
	}
}

func (c *Cache) Get(key string) (typecode.Code, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	code, ok := c.store[key]
	return code, ok
}

func (c *Cache) Set(key string, code typecode.Code) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.store[key] = code
}
