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

	c[key].IncrementHits(1)

	return c[key].Data()
}

func (c CacheSystem) GetNode(key string) *CacheNode {
	return c[key]
}

func (c CacheSystem) Set(key string, data interface{}, expirationSeconds time.Duration) *CacheNode {
	var cn CacheNode
	cn.SetData(data)
	cn.SetCreatedAt(time.Now())
	cn.SetExpireSeconds(expirationSeconds)
	cn.SetExpireHits(0)

	c[key] = &cn
	return c[key]
}

func (c CacheSystem) SetExpirationByHits(key string, nHits uint) {
	node := c[key]
	if node == nil {
		return
	}

	node.SetExpireHits(nHits)
}

func (c CacheSystem) Drop(key string) {
	c[key] = nil
}
