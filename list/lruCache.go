package list

import "container/list"

type LRUCache struct {
	cacheList *list.List
	hash      map[int]*list.Element
	capacity  int
}

type Item struct {
	Key   int
	Value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		hash:      make(map[int]*list.Element),
		cacheList: list.New(),
		capacity:  capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if cached, ok := this.hash[key]; ok {
		this.cacheList.MoveToFront(cached)
		return cached.Value.(Item).Value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	// if not exist, then push front to list

	if cached, ok := this.hash[key]; !ok {
		if this.cacheList.Len() == this.capacity {
			pop := this.cacheList.Back()
			this.cacheList.Remove(pop)
			delete(this.hash, pop.Value.(Item).Key)
		}
		this.hash[key] = this.cacheList.PushFront(Item{key, value})
	} else {
		cached.Value = Item{key, value}
		this.cacheList.MoveToFront(cached)
	}
}
