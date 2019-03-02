package main

import (
	"fmt"
	"sync/atomic"
)

//极客时间29章的教程,这章讲的是原子操作
// 问题: sync/atomic提供了几种原子操作,可操作的类型又有哪些?
// 回答: sync/atomic包中的函数可以做的原子操作有 add加法,CAS,load加载,store存储,swap交换
//      数据类型 int32 int63 uint32 uint64 uintptr unsafe.Pointer(未提供原子加法操作的函数)...一个特殊的atomic.Value,被用来存储任意的类型
// 原子操作一定要穿值的指针!!!!
func main() {
	//var at atomic.Value
	//fmtI32()
	addf()
}

//以一个int32的例子,来练习一下原子操作....其实这破单线程无法模拟呀
func fmtI32() {
	var i int32 = 1
	//原子操作,返回新值
	atomic.AddInt32(&i, 5)
	atomic.LoadInt32(&i)
	fmt.Println("我原子增加以后的值", i)
	fmt.Println("我原子加载的值", i)
	atomic.StoreInt32(&i, 99)
	fmt.Println("我原子存储以后的值", i)
	atomic.SwapInt32(&i, 100)
	fmt.Println("我原子交换以后的值", i)
	//CAS的测试
	atomic.CompareAndSwapInt32(&i, 99, 101)
	fmt.Println("模拟原子交换失败的值", i)
	atomic.CompareAndSwapInt32(&i, 100, 101)
	fmt.Println("模拟原子交换成功的值", i)

}

//两种方法原子ADD增加一个负值,两种方法得到的补码都是一致的.
func addf() {
	var i int32 = 1
	atomic.AddInt32(&i, int32(-3))
	fmt.Println("减法int32--atomic", i)
	var u uint32 = 99
	//临时变量int32绕过编译器
	i2 := int32(-3)
	//这边来进行uint32的转换,才可以成功减去
	atomic.AddUint32(&u, uint32(i2))
	fmt.Println("减法uint32--atomic", u)

	//第二种uint32的减法
	h := -3
	// 差量的绝对值减1,转换成uint32,此unt32补码和第一种方法差值的补码相同,计算机视为同一数字
	i3 := uint32(^(-h - 1))
	var u2 uint32 = 99
	atomic.AddUint32(&u2, i3)
	fmt.Println("减法uint32--atomic", u2)
}
