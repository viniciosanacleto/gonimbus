package cache

import "time"

type CacheNode struct {
	data          interface{}
	createdAt     time.Time
	expireSeconds time.Duration
	expireHits    uint
	countHits     uint
}

func (c *CacheNode) Data() interface{} {
	return c.data
}

func (c *CacheNode) SetData(data interface{}) {
	c.data = data
}

func (c *CacheNode) CreatedAt() time.Time {
	return c.createdAt
}

func (c *CacheNode) SetCreatedAt(createdAt time.Time) {
	c.createdAt = createdAt
}

func (c *CacheNode) ExpireSeconds() time.Duration {
	return c.expireSeconds
}

func (c *CacheNode) SetExpireSeconds(expireSeconds time.Duration) {
	c.expireSeconds = expireSeconds
}

func (c *CacheNode) ExpireHits() uint {
	return c.expireHits
}

func (c *CacheNode) SetExpireHits(expireHits uint) {
	c.expireHits = expireHits
}

func (c *CacheNode) SetCountHits(countHits uint) {
	c.countHits = countHits
}

func (c *CacheNode) IncrementHits(n uint) {
	c.countHits += n
}

func (c *CacheNode) HasExpired() bool {

	// Verify time expiration
	if c.expireSeconds > 0 {
		now := time.Now()
		expTime := c.createdAt.Add(c.expireSeconds * time.Second)

		if expTime.Before(now) {
			return true
		}
	}

	// Verify hits expiration
	if c.expireHits > 0 {
		if c.countHits >= c.expireHits {
			return true
		}
	}

	return false
}
