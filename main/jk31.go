package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// 极客时间31章学习,此章学习的内容是 sync.waitGroup和sync.Once的应用
//wg使用的时候,我应该先初始化wg的数量,然后在go func进行 down 防止 并发操作 add和down造成 wg的count小于0最后统一的wait
// 杜绝对同一个waitGroup值的两种操作并发执行
// once和waitGroup的比较,once用到了原子值还有锁,waitGroup只用到了原子值

func main() {
	//一个简单的例子练习sync.waitGroup
	//simpleWait()
	//simpleWait2()
	//关于Once的练习,无论我扔几个函数进去,老子只执行一个
	//once()
	jkt()
}

func simpleWait() {
	var d sync.WaitGroup
	d.Add(2)
	//切记,这里不可以像goroutine传值,否则这里是值传递,是不行的,对外层的waitGroup起不到任何的作用
	go func() {
		defer func() {
			fmt.Print("down 1")
			d.Done()
		}()
	}()
	go func() {
		defer func() {
			fmt.Print("down 2")
			d.Done()
		}()
	}()
	d.Wait()
	fmt.Println("结束等待")
}

//反面教材,不能先让wg的值小于1否则会panic
func simpleWait2() {
	var d sync.WaitGroup
	//切记,这里不可以像goroutine传值,否则这里是值传递,是不行的,对外层的waitGroup起不到任何的作用
	//go func() {
	//	//fmt.Println("add先执行")
	//	time.Sleep(time.Second)
	//	defer func() {
	//		d.Done()
	//	}()
	//}()
	go func() {
		d.Add(1)
	}()
	go func() {
		d.Wait()
	}()
	time.Sleep(time.Second)
	fmt.Println("结束等待")
}

//这里练习 sync.once的语法
// sync.once的Do方法进行了二次检查,防止并发导致多次执行,类似于java中的二次加锁的单例模式
// 这里保证了只执行一次方法有两个前提: 1.原子的uint32 2.在Do方法中加锁二次检查.
func once() {
	//结构体包含值字段的mutex 复制会导致值传递失效
	var d sync.Once
	//Do方法的两个特点
	// 1.多个goroutine如果调用同一个once.do,那么如果被执行的方法时间很长的话,所有的goroutine都会阻塞!
	// 2.函数执行以后,内部的down的值一定是1,无论你的函数执行成功或者失败,都会defer
	d.Do(func() {
		fmt.Println("1")
	})
	d.Do(func() {
		fmt.Println("2")
	})
}

//极客时间特殊练习
// 这个例子有两种结果一种是打印recover的值,另一种是直接Do task 5的逻辑,不过两个go func 都是会走的
// 并且wg.wait也而已成功的结束阻塞
func jkt() {
	// 示例3。
	once := sync.Once{}
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		defer func() {
			if p := recover(); p != nil {
				fmt.Printf("fatal error: %v\n", p)
			}
		}()
		once.Do(func() {
			fmt.Println("Do task. [4]")
			panic(errors.New("something wrong"))
		})
	}()
	go func() {
		defer wg.Done()
		once.Do(func() {
			fmt.Println("Do task. [5]")
		})
		fmt.Println("Done. [5]")
	}()
	wg.Wait()
}
