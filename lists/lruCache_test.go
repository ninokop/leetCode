package list

import "testing"
import "log"

func TestLRUCache(t *testing.T) {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(1, 1)
	log.Println(cache.Get(1))
	cache.Put(3, 3)
	log.Println(cache.Get(2))
	cache.Put(4, 4)
	log.Println(cache.Get(1))
	log.Println(cache.Get(3))
	log.Println(cache.Get(4))
}
