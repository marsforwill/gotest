package main

import "container/list"

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

type LRUCache struct {
	list *list.List
	m    map[int]list.Element
	cap  int
}

func Constructor(capacity int) LRUCache {
	l := list.New()
	m := make(map[int]list.Element, capacity)
	return LRUCache{
		list: l,
		m:    m,
		cap:  capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	k, ok := this.m[key]
	if !ok {
		return -1
	}
	this.list.Remove(&k)
	this.list.InsertBefore(k)
	return k
}

func (this *LRUCache) Put(key int, value int) {
	this.m[key] = value
	head := this.list.Front()
	this.list.InsertBefore(value, head)
	if this.list.Len() > this.cap {
		tail := this.list.Back()
		this.list.Remove(tail)
		delete(this.m, tail.Value.(int))
	}
}

func main() {

}
