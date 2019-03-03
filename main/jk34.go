package main

import (
	"fmt"
	"sync"
)

//极客时间34章练习,这节课讲的是sync.map

//笔记: 并发安全的map的键值不能是 函数,字典,切片, 因为无法产生hash映射
func main() {
	//simpleSync()
	exampleMap()
}

// sync.map是结构体值类型,不可以用map,但是普通的map是可以的
func simpleSync() {
	//可以的
	//a := make(map[int]int)
	//只可以用这样的方式声明
	c := sync.Map{}
	//var c sync.Map
	c.Store("1", "2")
	value3, _ := c.Load("1")
	fmt.Print(value3)
}

// 此例子摘选自极客时间
//让一个sync.map 增删改查特殊的值,我选择使用一个结构体包含它 然后实现其安全
func exampleMap() {
	var h IntStrMap
	//这里会报错,因为我外层自己的方法已经做了限制
	// h.Store("1","2")
	//这里才对,调用我自己的存储方法
	h.Store(1, "2")
	h.Store(1, "3")
	// 加载或存储,如果key不存在返回的是false,key存在返回的是true并且是其原来的值
	fmt.Println(h.LoadOrStore(1, "32222"))
	fmt.Println(h.LoadOrStore(2, "4"))
}

type IntStrMap struct {
	m sync.Map
}

//包装了一层删除方法
func (iMap *IntStrMap) Delete(key int) {
	iMap.m.Delete(key)
}

//包装了一层加载方法
func (iMap *IntStrMap) Load(key int) (value string, ok bool) {
	v, ok := iMap.m.Load(key)
	if v != nil {
		value = v.(string)
	}
	return
}

//包装了一层加载或者存储的方法,其实也就是修改
func (iMap *IntStrMap) LoadOrStore(key int, value string) (actual string, loaded bool) {
	a, loaded := iMap.m.LoadOrStore(key, value)
	actual = a.(string)
	return
}

//遍历方法
func (iMap *IntStrMap) Range(f func(key int, value string) bool) {
	f1 := func(key, value interface{}) bool {
		return f(key.(int), value.(string))
	}
	iMap.m.Range(f1)
}

func (iMap *IntStrMap) Store(key int, value string) {
	iMap.m.Store(key, value)
}
