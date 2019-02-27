package main

import (
	"fmt"
	"sync/atomic"
	//"time"
)

//极客时间第17讲 还是goroutine部分
// 其实吧 任何的 比如 java中的 while(true) golang中的 for{}  一定要加入 time.sleep  防止cpu飙升爆表
func main() {
	//gochan()
	orderGo()

}

//这个函数我用空结构体chan来控制 go 主/协程之间的运行

func gochan() {
	ch := make(chan struct{}, 10)
	for i := 0; i < 10; i++ {
		go func(j int) {
			fmt.Println(j)
			//这是空结构体的表示方法
			ch <- struct{}{}
		}(i)
	}
	for is := 0; is < 10; is++ {
		<-ch
	}
}

//这里写个复杂的控制程序顺序输出的函数,参照极客时间17章的demo,同时也可以作为其他并发编程的参考
func orderGo() {
	var count uint32
	trigger := func(i uint32, fn func()) {
		for {
			if n := atomic.LoadUint32(&count); n == i {
				fn()
				atomic.AddUint32(&count, 1)
				break
			}
			//如果不休眠的话,下次的并发trigger将会卡死在这里,一直循环!!
			//time.Sleep(time.Nanosecond)
		}
	}

	//此 i(形参) --(实际的值传递) 就是启用 goroutine时,那个当次迭代的序号

	for i := uint32(0); i < 10; i++ {
		go func(k uint32) {
			fn := func() {
				fmt.Println(k)
			}
			trigger(k, fn)
		}(i)
	}
	//这种控制方式是最low的
	//time.Sleep(time.Second * 10)
	//这种控制方式 在到10以后退出,也会竞争的
	// 不一定是 我要开始竞争了哦 然后 0,1,2,3,4,5,6,7,8,9 要用并行的态度去看待这个问题
	fmt.Println("我要开始竞争了哦")
	trigger(10, func() {})
}
