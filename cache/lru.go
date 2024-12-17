package cache

import (
	"container/list"
	"sync"
)

type Cache struct {
	capacity  int
	mutex     sync.Mutex
	items     map[string]*list.Element
	evictList *list.List
}

type entry struct {
	key   string
	value string
}

func NewCache(capacity int) *Cache {
	return &Cache{
		capacity:  capacity,
		items:     make(map[string]*list.Element),
		evictList: list.New(),
	}
}

func (c *Cache) Get(key string) (string, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if elem, ok := c.items[key]; ok {
		c.evictList.MoveToFront(elem)
		return elem.Value.(*entry).value, true
	}

	return "", false
}

func (c *Cache) Set(key, value string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if elem, ok := c.items[key]; ok {
		c.evictList.MoveToFront(elem)
		elem.Value.(*entry).value = value
		return
	}

	if c.evictList.Len() >= c.capacity {
		c.evict()
	}

	newEntry := &entry{key, value}
	elem := c.evictList.PushFront(newEntry)
	c.items[key] = elem
}

func (c *Cache) evict() {
	lastElem := c.evictList.Back()
	if lastElem != nil {
		c.evictList.Remove(lastElem)
		delete(c.items, lastElem.Value.(*entry).key)
	}
}
