package lru

import (
	"container/list"
)

type Cache struct {
	maxBytes  int64                         //允許使用的最大內存
	nbytes    int64                         //當前使用的內存
	ll        *list.List                    //快取雙向鍊表
	dict      map[string]*list.Element      //快取映射dict
	OnEvicted func(key string, value Value) //紀錄被刪除後的callback
}

//雙向鍊表節點數據類型
type entry struct {
	key   string
	value Value
}

//用以返回entry內 Value 所佔的內存大小
type Value interface {
	Len() int
}

//實例化
func New(maxBytes int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		ll:        list.New(),
		dict:      make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

//查找快取功能 1. 從dict找到對應鍊表的節點 2. 將該節點移至鍊尾
func (c *Cache) Get(key string) (value Value, ok bool) {
	if el, ok := c.dict[key]; ok {
		c.ll.MoveToFront(el)
		kv := el.Value.(*entry)
		return kv.value, true
	}

	return
}

//緩存淘汰功能 1. 移除最近最少訪問的節點 2. 將dict內的key值移除
func (c *Cache) RemoveOldest() {
	el := c.ll.Back() //取得鍊首節點
	if el != nil {
		c.ll.Remove(el)
		kv := el.Value.(*entry)
		delete(c.dict, kv.key)
		c.nbytes -= int64(len(kv.key)) + int64(kv.value.Len())
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)
		}
	}
}

//新增與修改緩存功能
func (c *Cache) Add(key string, value Value) {
	if el, ok := c.dict[key]; ok {
		c.ll.MoveToFront(el)
		kv := el.Value.(*entry)
		c.nbytes += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
	} else {
		el := c.ll.PushFront(&entry{key: key, value: value})
		c.dict[key] = el
		c.nbytes += int64(len(key)) + int64(value.Len())
	}

	for c.maxBytes != 0 && c.maxBytes < c.nbytes {
		c.RemoveOldest()
	}
}

//測試用
func (c *Cache) Len() int {
	return c.ll.Len()
}
