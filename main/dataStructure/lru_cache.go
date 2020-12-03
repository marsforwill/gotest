package main

import (
	"container/list"
	"fmt"
)

//设计和构建一个“最近最少使用”缓存，该缓存会删除最近最少使用的项目。缓存应该从键映射到值(允许你插入和检索特定键对应的值)，并在初始化时指定最大容量。当缓存
//被填满时，它应该删除最近最少使用的项目。
//
// 它应该支持以下操作： 获取数据 get 和 写入数据 put 。
//
// 获取数据 get(key) - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。
//写入数据 put(key, value) - 如果密钥不存在，则写入其数据值。当缓存容量达到上限时，它应该在写入新数据之前删除最近最少使用的数据值，从而为新
//的数据值留出空间。
//
// 示例:
//
// LRUCache cache = new LRUCache( 2 /* 缓存容量 */ );
//
//cache.put(1, 1);
//cache.put(2, 2);
//cache.get(1);       // 返回  1
//cache.put(3, 3);    // 该操作会使得密钥 2 作废
//cache.get(2);       // 返回 -1 (未找到)
//cache.put(4, 4);    // 该操作会使得密钥 1 作废
//cache.get(1);       // 返回 -1 (未找到)
//cache.get(3);       // 返回  3
//cache.get(4);       // 返回  4

//leetcode submit region begin(Prohibit modification and deletion)
// 自定义需要的node信息
type Node struct {
	key   int
	value int
}

type LRUCache struct {
	list *list.List
	m    map[int]*list.Element
	cap  int
}

func Constructor(capacity int) LRUCache {
	// 链表维护在缓存中的最近节点访问顺序
	l := list.New()
	// map维护在缓存中的key -> list node
	m := make(map[int]*list.Element, capacity)
	return LRUCache{
		list: l,
		m:    m,
		cap:  capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	value, ok := this.m[key]
	if !ok {
		return -1
	}
	this.list.MoveToFront(value)
	// list元素的结构莫名的存了两层
	return value.Value.(list.Element).Value.(*Node).value
}

func (this *LRUCache) Put(key int, value int) {
	v, ok := this.m[key]
	if ok {
		v.Value.(list.Element).Value.(*Node).value = value
		this.list.MoveToFront(v)
		return
	}
	e := list.Element{
		Value: &Node{
			key:   key,
			value: value,
		},
	}
	this.list.PushFront(e)
	this.m[key] = this.list.Front()
	if this.list.Len() > this.cap {
		tail := this.list.Back()
		delete(this.m, tail.Value.(list.Element).Value.(*Node).key)
		this.list.Remove(tail)
	}
}

func main() {
	cache := Constructor(1 /* 缓存容量 */)

	cache.Put(2, 1)
	//cache.Put(2, 2)
	cache.Get(2)    // 返回  1
	cache.Put(3, 2) // 该操作会使得密钥 2 作废
	a := cache.Get(2)
	fmt.Println(a)  // 返回 -1 (未找到)
	cache.Put(4, 4) // 该操作会使得密钥 1 作废
	cache.Get(1)    // 返回 -1 (未找到)
	cache.Get(3)    // 返回  3
	cache.Get(4)    // 返回  4
}
