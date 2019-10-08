package cache

type Callback func(item *Item) string

type Interface interface {
	Get(key string, callback Callback)
}
