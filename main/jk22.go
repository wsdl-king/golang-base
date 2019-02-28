package main

import (
	"fmt"
)

//极客时间22讲, panic defer recover等等练习
func main() {
	//aroundP()
	//moreDefer1()
	generate()
}

//包含panic
func aroundP() {
	fmt.Println("进入panic")
	//被延迟执行的是defer函数 不是defer语句,必须要把defer函数放在最前面
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("defer函数处理结果")
		}
	}()
	//这里接受的是interface{} 类型的,所有的类型都实现了interface{}
	// panic 以后panic以后的语句都是不会打印的!
	//下文我们要提到defer的用法了,defer代码块是用来执行延迟代码的,函数结束的那一刻,我的defer再运行,无论defer的原因是什么
	// 注意 被延迟执行的是defer函数 不是defer语句
	panic("包含?")
	fmt.Println("退出panic")
}

//多个defer 其实会加入队列的,这个队列是先进后出的
func moreDefer1() {
	defer func() {
		fmt.Println("开始defer")
	}()
	for i := 0; i < 10; i++ {
		defer func(j int) {
			fmt.Println("这是我第几次打印defer语句", j)
		}(i)
	}
	defer func() {
		fmt.Println("defer语句结束")
	}()
}

//多个defer 其实会加入队列的,这个队列是先进后出的
func moreDefer2() {
	defer func() {
		fmt.Println("开始defer")
	}()
	for i := 0; i < 10; i++ {
		defer func() {
			fmt.Println("这是我第几次打印defer语句", i)
		}() //比较1和2的不同你会发现,这里传值和不传值打印的值是不一样的,这里需要特别注意
		//有点像 go那个孙子的陷阱 其实你理解好这个执行的顺序就好了 defer会初始化 加入先进后出队列,如果有外部值传入
		//会一起外部初始化的,如果没有的话 最后取值的话,会拿最后的i的值副本
	}
	defer func() {
		fmt.Println("defer语句结束")
	}()
}

//defer引发panic
// 这里也是可以使用panic生成的,所以 增加一个经验 对程序错误recover你一定要用defer写在第一位
// 如果是 goroutine里面的panic defer一定要写在里面,  多个go routine 也是蛮有意思的
func generate() {
	i := make(chan struct{})
	go func() {
		i <- struct{}{}
		defer func() {
			fmt.Println(recover())
		}()
		panic("我生成1")
	}()
	go func() {
		panic("我生成2")
	}()
	<-i
}
