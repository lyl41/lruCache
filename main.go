package main

import (
	"cache/lru_cache"
	"container/list"
	"fmt"
	"os"
)

func testList() {
	l := list.New()
	m := make(map[int]int)
	key := 2
	m[key] = 3
	now := new(list.Element)
	now.Value = key
	l.PushBack(key)
	fmt.Println(now)
	fmt.Println(l.Front().Value)

	if _, found := m[l.Front().Value.(int)]; found {
		fmt.Println("found")
	} else {
		fmt.Println("not found")
	}

	if _, found := m[now.Value.(int)]; found {
		fmt.Println("found")
	} else {
		fmt.Println("not found")
	}

	fmt.Println(l.Front().Value)
	fmt.Println(l.Front().Value)
	os.Exit(0)
}

func main() {
	//testList()
	cache := lru_cache.NewLruCache(3)
	cache.Set(1, []byte("2"))
	//cache.DebugShowMapData()

	cache.Set(2, []byte("2"))
	//cache.DebugShowMapData()

	cache.Set(3, []byte("2"))
	//cache.DebugShowMapData()

	cache.Set(4, []byte("2")) // 2 3 4
	//cache.DebugShowMapData()
	//cache.DebugShowLruList()
	//fmt.Println(cache.Length())

	cache.Set(5, []byte("3")) // 3 4 5
	//cache.DebugShowMapData()
	//cache.DebugShowLruList()

	cache.Set(2, []byte("4")) // 4 5 2
	//cache.DebugShowLruList()
	//cache.DebugShowMapData()

	//fmt.Println(cache.Get(4))
	//cache.DebugShowLruList()
	//cache.DebugShowMapData()

	//fmt.Println(cache.Length())

	cache.Delete(1)
	//cache.DebugShowLruList()
	//cache.DebugShowMapData()

	cache.Get(2) // 4 5 2
	//cache.DebugShowLruList()
	//cache.DebugShowMapData()

	cache.Set(4, []byte("3")) // 5 2 4
	cache.DebugShowLruList()
	cache.DebugShowMapData()

}