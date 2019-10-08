package cache

import "time"

type ItemInterface interface {
	SetExpire(ttl time.Duration)
}

type Item struct {
	TTL time.Duration
}

func NewItem() Item {
	return Item{}
}

func (item *Item) SetExpire(ttl time.Duration) {
	item.TTL = ttl
}
