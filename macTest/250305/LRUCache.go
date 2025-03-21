package main

//import "container/list"
//
//type LRUEntry struct {
//	key int
//	val int
//}
//
//type LRUCache struct {
//	list       *list.List
//	mapKeyNode map[int]*list.Element
//	capacity   int
//}
//
//func Constructor(capacity int) LRUCache {
//	return LRUCache{list.New(), map[int]*list.Element{}, capacity}
//}
//
//func (this *LRUCache) Get(key int) int {
//	v := this.mapKeyNode[key]
//	if v == nil {
//		return -1
//	}
//	this.list.MoveToFront(v)
//	return v.Value.(LRUEntry).val
//}
//
//func (this *LRUCache) Put(key int, value int) {
//	v := this.mapKeyNode[key]
//	if v != nil {
//		v.Value = LRUEntry{key, value}
//	} else {
//		this.mapKeyNode[key] = this.list.PushFront(LRUEntry{key, value})
//		if len(this.mapKeyNode) > this.capacity {
//			delete(this.mapKeyNode, this.list.Remove(this.list.Back()).Value.(LRUEntry).key)
//			this.list.Remove(this.list.Back())
//		}
//	}
//}
//
///**
// * Your LRUCache object will be instantiated and called as such:
// * obj := Constructor(capacity);
// * param_1 := obj.Get(key);
// * obj.Put(key,value);
// */
