package cache

import "time"

type CacheSystem map[string]*CacheNode

func (c CacheSystem) Get(key string) interface{} {
	// If the node does not exist
	if c[key] == nil {
		return nil
	}

	// If the node has expired
	if c[key].HasExpired() {
		delete(c, key)
		return nil
	}

	c[key].CountAccess++

	return c[key].Data
}

func (c CacheSystem) Set(key string, data interface{}, expirationSeconds time.Duration) {
	var cn CacheNode
	cn.Data = data
	cn.CreatedAt = time.Now()
	cn.ExpireSeconds = expirationSeconds

	c[key] = &cn
}
