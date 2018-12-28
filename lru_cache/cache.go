package lru_cache

import (
	"container/list"
	"fmt"
	"sync"
)

type value struct {
	data   []byte
	lruPos *list.Element
}

type lruCache struct {
	maxSize int
	data    map[interface{}]*value
	lck     *sync.Mutex
	lru     *list.List
}

func NewLruCache(maxLength int) *lruCache {
	return &lruCache{
		maxSize: maxLength,
		data:    make(map[interface{}]*value),
		lck:     new(sync.Mutex),
		lru:     list.New(),
	}
}

func (c *lruCache) Set(key interface{}, data []byte) {
	c.lck.Lock()
	defer c.lck.Unlock()
	if val, found := c.data[key]; found {
		c.deleteLruItem(val.lruPos)  //删除原先在list中的位置
		val.lruPos = c.updateNewItem(key) // 追加到list末尾，更新位置标示
		c.data[key] = val
	} else {
		var pos *list.Element
		if len(c.data) < c.maxSize {
			pos = c.updateNewItem(key) //直接追加到list末尾
		} else {
			c.deleteLruItem(c.lru.Front()) //删除最久未使用的
			pos = c.updateNewItem(key)     //将本次key更新到list的末尾
		}
		val := &value{
			data:   data,
			lruPos: pos,
		}
		c.data[key] = val
	}
}

func (c *lruCache) Get(key interface{}) (exist bool, data []byte) {
	c.lck.Lock()
	defer c.lck.Unlock()
	if val, found := c.data[key]; found {
		data = val.data
		c.deleteLruItem(val.lruPos)  //删除原先在list中的位置
		val.lruPos = c.updateNewItem(key) // 追加到list末尾，更新位置标示
		c.data[key] = val
		exist = true
	}
	return
}

func (c *lruCache) deleteLruItem(pos *list.Element) (item *list.Element) {
	c.lru.Remove(pos)
	delete(c.data, pos.Value)
	return
}

func (c *lruCache) updateNewItem(key interface{}) (item *list.Element) {
	item = c.lru.PushBack(key)
	return
}

func (c *lruCache) Length () int {
	c.lck.Lock()
	defer c.lck.Unlock()
	return len(c.data)
}

func (c *lruCache)DebugShowMapData () {
	c.lck.Lock()
	defer c.lck.Unlock()
	fmt.Println("=== map ===")
	for k, v := range c.data {
		fmt.Println(k, v)
	}
	fmt.Println("=== map over ===")
}
func (c *lruCache)DebugShowLruList () {
	c.lck.Lock()
	defer c.lck.Unlock()
	fmt.Println("=== list ===")
	for v := c.lru.Front(); v != nil; v = v.Next() {
		fmt.Print(v.Value, " ")
	}
	fmt.Println("\n=== list over ===")
}

func (c *lruCache) Delete (key interface{}) {
	c.lck.Lock()
	defer c.lck.Unlock()
	if val, found := c.data[key]; found {
		c.deleteLruItem(val.lruPos)
	}
}