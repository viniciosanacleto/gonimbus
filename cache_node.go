package cache

import "time"

type CacheNode struct {
	Data          interface{}
	CreatedAt     time.Time
	ExpireSeconds time.Duration
	CountAccess   uint
}

func (c *CacheNode) HasExpired() bool {
	now := time.Now()
	expTime := c.CreatedAt.Add(c.ExpireSeconds * time.Second)

	if expTime.Before(now) {
		return true
	}

	return false
}
