package main

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
)

//极客时间35讲,讲述还是sync.map,第二种指定类型的sync.map的实现,我需要在生成结构体的时候传入指定的type
// 笔记: 讲一下sync.Map的内部实现
// 有一个atomic.value的只读字典,还有一个原生集合的脏字典,脏字典的修改需要用到锁
// 脏字典在查询足够多的情况下会转换成只读字典,在这之后,一旦有新的键值对存入,就会依据只读字典建立脏字典
//脏字典和只读字典不是实时同步的,我做到读和写操作 都会先去只读字典查找有效的键
// 在读操作多 写操作少的情况下,性能还是可观的
// 删除情况下 只读字典有k的情况下,会逻辑删除
// 存储情况下, 有限存储到只读字典,条件是 已经有k,并且不是已删除. 否则的话存储到脏字典,删除只读字典的已删除标记

func main() {
	// type.方法名字
	a, _ := NewConMap(reflect.TypeOf("2"), reflect.TypeOf(1))
	e := a.Put("2", 3)
	if e == nil {
		fmt.Println(a.Load("2"))
	}
}

// 这里我写sync.map的另一种实现
type MyhashMap struct {
	m sync.Map
	k reflect.Type
	v reflect.Type
}

//你要搞清楚这是方法还是函数,我要生成一个新的结构体,那么我就不是方法了,我是一个函数衍生出来的
// 居然将所学的知识混淆了,小兄台仍需努力呀
func NewConMap(k1 reflect.Type, v1 reflect.Type) (*MyhashMap, error) {
	// 这里要检查key和value 的类型 还有 key必须是可比较的
	if k1 == nil {
		return nil, errors.New("nil key type")
	}
	if !k1.Comparable() {
		return nil, fmt.Errorf("incomparable key type: %s", k1)
	}
	if v1 == nil {
		return nil, errors.New("nil value type")
	}
	cMap := &MyhashMap{
		k: k1,
		v: v1,
	}
	return cMap, nil
}

//我自己的put方法
func (ma *MyhashMap) Put(k interface{}, v interface{}) error {
	if reflect.TypeOf(k) != ma.k {
		// k的类型不对
		return errors.New(" k type is   invalid")
	}
	if reflect.TypeOf(v) != ma.v {
		// k的类型不对
		return errors.New(" v type is   invalid")
	}
	ma.m.Store(k, v)
	return nil
}

func (ma *MyhashMap) Load(k interface{}) (v interface{}, b bool) {
	if reflect.TypeOf(k) != ma.k {
		// k的类型不对
		return nil, false
	}
	return ma.m.Load(k)
}

func (ma *MyhashMap) LoadOrStore(k interface{}, v interface{}) (v2 interface{}, b bool) {
	if reflect.TypeOf(k) != ma.k {
		// k的类型不对
		return nil, false
	}
	if reflect.TypeOf(v) != ma.v {
		// k的类型不对
		return nil, false
	}
	return ma.m.LoadOrStore(k, v)
}
func (ma *MyhashMap) Delete(key interface{}) {
	if reflect.TypeOf(key) != ma.k {
		return
	}
	ma.m.Delete(key)
}
func (cMap *MyhashMap) Range(f func(key, value interface{}) bool) {
	cMap.m.Range(f)
}
