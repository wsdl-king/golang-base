package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

//极客时间27讲 条件变量
// 笔记:  条件变量的初始化离不开互斥锁,并且它的方法有的也是基于互斥所
// 条件变量提供的方法有三个: 等待通知(wait),单发通知(signal),广播通知(broadcast)
// 条件变量可以协调那些想访问共享资源的线程,当共享资源发生变化的时候,它可以被用来通知被互斥锁阻塞的线程

//sync.NewCond 方法接受的是一个Locker接口类型的,而这个接口类型的实现是 读写锁的指针
func main() {
	//打印结果是2对2对出现的,必定是 写-写/读-读
	//signalcondiT()
	broadcastcondiT()
}

// 我们在利用条件变量 等待通知 的时候,需要在它基于的那个互斥锁保护下进行
// 而进行单发通知或者广播通知的时候,是相反的, 需要在对应的互斥锁解锁以后再做操作
//单项广播函数
func signalcondiT() {
	// mailbox 代表信箱。
	// 0代表信箱是空的，1代表信箱是满的。
	var mailbox uint8
	// lock 代表信箱上的锁。
	//条件锁的初始化离不开互斥锁
	var lock sync.RWMutex
	//写锁的条件控制
	// sendCond 代表专用于发信的条件变量。
	sendCond := sync.NewCond(&lock)
	//读锁的条件控制
	// recvCond 代表专用于收信的条件变量。
	recvCond := sync.NewCond(lock.RLocker())

	// sign 用于传递演示完成的信号。
	sign := make(chan struct{}, 3)
	max := 5
	go func(max int) { // 用于发信。
		defer func() {
			sign <- struct{}{}
		}()
		for i := 1; i <= max; i++ {
			time.Sleep(time.Millisecond * 500)
			//增加写锁
			lock.Lock()
			for mailbox == 1 {
				// 这里是条件变量写锁,我想等待通知,那我就要在写锁的保护下进行
				sendCond.Wait()
			}
			log.Printf("sender [%d]: the mailbox is empty.", i)
			mailbox = 1
			log.Printf("sender [%d]: the letter has been sent.", i)
			//解除写锁
			lock.Unlock()
			//老子写完了,信箱里只有一个信了,通知你个孙子可以取了!
			//我想接受(读锁),那我必须在你解除写锁以后我在操作
			recvCond.Signal()
		}
	}(max)
	go func(max int) { // 用于收信。
		defer func() {
			sign <- struct{}{}
		}()
		for j := 1; j <= max; j++ {
			time.Sleep(time.Millisecond * 100)
			//老子做的安全判断,看信箱里有没有信,也就是读, 所以呢 要在读锁的保护下
			lock.RLock()
			for mailbox == 0 {
				//我曹 还真没有信,那老子就等待(读锁等待)
				// 这里我fmt一个数字 表明:
				//这里的4个字只有在阻塞前会输出,而阻塞要看信箱里是否有信,如果有信则不会输出,所以这里的输出是没法确定的
				//除非我显示的等待时间大于send函数,我才会不会输出.,若是小于send函数的时间 我每次都会被输出
				fmt.Println("我先进入")
				recvCond.Wait()
			}
			log.Printf("receiver [%d]: the mailbox is full.", j)
			mailbox = 0
			log.Printf("receiver [%d]: the letter has been received.", j)
			//我曹 我被不知名的唤醒了,那我的安全判断也要结束了
			lock.RUnlock()
			//老子读完了,爷爷给你发个信号,你可以继续写了
			sendCond.Signal()
		}
	}(max)

	<-sign
	<-sign
}

//参照就极客时间demo62.go
//多项广播函数
// 注意注意 这里我用的是互斥锁 并不是读写锁 不信你看, 我在声明条件变量的时候,我用的都是&lock 指针变量
func broadcastcondiT() {
	// mailbox 代表信箱。
	// 0代表信箱是空的，1代表信箱是满的。
	var mailbox uint8
	// lock 代表信箱上的锁。 --注意 这里是全体锁不是读写锁
	var lock sync.Mutex
	// sendCond 代表专用于发信的条件变量。
	sendCond := sync.NewCond(&lock)
	// recvCond 代表专用于收信的条件变量。
	recvCond := sync.NewCond(&lock)

	// send 代表用于发信的函数。
	send := func(id, index int) {
		lock.Lock()
		for mailbox == 1 {
			sendCond.Wait()
		}
		log.Printf("sender [%d-%d]: the mailbox is empty.",
			id, index)
		mailbox = 1
		log.Printf("sender [%d-%d]: the letter has been sent.",
			id, index)
		lock.Unlock()
		recvCond.Broadcast()
	}

	// recv 代表用于收信的函数。
	recv := func(id, index int) {
		lock.Lock()
		for mailbox == 0 {
			recvCond.Wait()
		}
		log.Printf("receiver [%d-%d]: the mailbox is full.",
			id, index)
		mailbox = 0
		log.Printf("receiver [%d-%d]: the letter has been received.",
			id, index)
		lock.Unlock()
		sendCond.Signal() // 确定只会有一个发信的goroutine。
	}

	// sign 用于传递演示完成的信号。
	sign := make(chan struct{}, 3)
	max := 6
	go func(id, max int) { // 用于发信。
		defer func() {
			sign <- struct{}{}
		}()
		for i := 1; i <= max; i++ {
			time.Sleep(time.Millisecond * 500)
			send(id, i)
		}
	}(0, max)

	//下面的两个收信会全部被唤醒 然后开始竞争的,竞争失败的都是弟弟,可能一个发,接受的是 1 或者是2!
	go func(id, max int) { // 用于收信。
		defer func() {
			sign <- struct{}{}
		}()
		for j := 1; j <= max; j++ {
			time.Sleep(time.Millisecond * 200)
			recv(id, j)
		}
	}(1, max/2)
	go func(id, max int) { // 用于收信。
		defer func() {
			sign <- struct{}{}
		}()
		for k := 1; k <= max; k++ {
			time.Sleep(time.Millisecond * 200)
			recv(id, k)
		}
	}(2, max/2)

	<-sign
	<-sign
	<-sign
}
