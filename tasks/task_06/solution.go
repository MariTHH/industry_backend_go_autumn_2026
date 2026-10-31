package main

import "container/list"

type entry[K comparable, V any] struct {
	key   K
	value V
}
type LRUCache[K comparable, V any] struct {
	capacity int
	ll       list.List
	items    map[K]*list.Element
}

func NewLRUCache[K comparable, V any](capacity int) *LRUCache[K, V] {
	return &LRUCache[K, V]{capacity: capacity, items: make(map[K]*list.Element)}
}

func (c *LRUCache[K, V]) Get(key K) (value V, ok bool) {

	if e, exists := c.items[key]; exists {
		c.ll.MoveToFront(e)

		return e.Value.(entry[K, V]).value, true
	}

	return value, false
}

func (c *LRUCache[K, V]) Set(key K, value V) {
	if c.capacity <= 0 {
		return
	}
	if e, exists := c.items[key]; exists {
		e.Value = entry[K, V]{key, value}

		return
	}
	if len(c.items) >= c.capacity {
		old := c.ll.Back()
		delete(c.items, old.Value.(entry[K, V]).key)
		c.ll.Remove(old)
	}

	c.items[key] = c.ll.PushFront(entry[K, V]{key, value})
}

type LRU[K comparable, V any] interface {
	Get(K) (V, bool)
	Set(K, V)
}
