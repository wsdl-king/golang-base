package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

//极客时间29章的教程,这章讲的是原子操作
// 问题: sync/atomic提供了几种原子操作,可操作的类型又有哪些?
// 回答: sync/atomic包中的函数可以做的原子操作有 add加法,CAS,load加载,store存储,swap交换
//      数据类型 int32 int63 uint32 uint64 uintptr unsafe.Pointer(未提供原子加法操作的函数)...一个特殊的atomic.Value,被用来存储任意的类型
// 原子操作一定要穿值的指针!!!!
func main() {
	//var at atomic.Value
	//fmtI32()
	//addf()
	//forAndCAS1()
	forAndCAS2()
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

//极客时间29讲的git-示例代码
// forAndCAS1 用于展示简易的自旋锁。
// 答案可能是 2,4,6,8,10,0   也可能是2,4,6,8...这里已经加到10了但是goroutine进行了切换...0,10
func forAndCAS1() {
	sign := make(chan struct{}, 2)
	num := int32(0)
	fmt.Printf("The number: %d\n", num)
	go func() { // 定时增加num的值。
		defer func() {
			sign <- struct{}{}
		}()
		for {
			time.Sleep(time.Millisecond * 500)
			newNum := atomic.AddInt32(&num, 2)
			fmt.Printf("The number: %d\n", newNum)
			if newNum == 10 {
				break
			}
		}
	}()
	go func() { // 定时检查num的值，如果等于10就将其归零。
		defer func() {
			sign <- struct{}{}
		}()
		for {
			if atomic.CompareAndSwapInt32(&num, 10, 0) {
				fmt.Println("The number has gone to zero.")
				break
			}
			time.Sleep(time.Millisecond * 500)
		}
	}()
	<-sign
	<-sign
}

//极客时间git-对照代码的demo
// forAndCAS2 用于展示一种简易的（且更加宽松的）互斥锁的模拟
// 利用CAS两个线程进行竞争比较并且交换值
func forAndCAS2() {
	sign := make(chan struct{}, 2)
	num := int32(0)
	fmt.Printf("The number: %d\n", num)
	max := int32(20)
	go func(id int, max int32) { // 定时增加num的值。
		defer func() {
			sign <- struct{}{}
		}()
		for i := 0; ; i++ {
			currNum := atomic.LoadInt32(&num)
			if currNum >= max {
				break
			}
			newNum := currNum + 2
			time.Sleep(time.Millisecond * 200)
			if atomic.CompareAndSwapInt32(&num, currNum, newNum) {
				fmt.Printf("The number: %d [%d-%d]\n", newNum, id, i)
			} else {
				fmt.Printf("The CAS operation failed. [%d-%d]\n", id, i)
			}
		}
	}(1, max)
	go func(id int, max int32) { // 定时增加num的值。
		defer func() {
			sign <- struct{}{}
		}()
		for j := 0; ; j++ {
			currNum := atomic.LoadInt32(&num)
			if currNum >= max {
				break
			}
			newNum := currNum + 2
			time.Sleep(time.Millisecond * 200)
			if atomic.CompareAndSwapInt32(&num, currNum, newNum) {
				fmt.Printf("The number: %d [%d-%d]\n", newNum, id, j)
			} else {
				fmt.Printf("The CAS operation failed. [%d-%d]\n", id, j)
			}
		}
	}(2, max)
	<-sign
	<-sign
}
