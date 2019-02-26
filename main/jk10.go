package main

import (
	"fmt"
	"time"
)

//极客时间第10章关于通道的练习
func main() {
	strings := make(chan string, 1)
	NoCacheChannl(strings)
	//go AbstractWriteChan(strings)
	//go AbstractReadChan(strings)
	//blockRun(strings)
	//练习for-break
	//forBreak()
}

func NoCacheChannl(strings chan string) {
	//非缓冲通道,需要用到go的协程,不然会一致阻塞,因为程序执行是一条线执行的
	go func() {
		strings <- "123"
		//这里我关闭通道,我也是可以继续接收的
		//close(strings)
	}()
	go func() {
		time.Sleep(time.Second * 2)
		//这里延迟关闭,但是返回的bool类型还是true,是因为里面还有数据,所以通道关闭以后再从通道取数据
		//返回的可能是true,也可能是false
		e := <-strings
		fmt.Println(e)
		//上文中,我往通道提交了一个值,我这里开启的协程接收, 这里再从chan里取数据会永久阻塞,因为里面已经没有数据了。
		//直到有新的数据,我才可以继续执行。 注意,这个协程是永久阻塞的,主线程是直接退出的,如果我不是以协程的方式
		//从chan里取数据,则永久阻塞,主线程也会报错的
		d := <-strings
		fmt.Println("第二次chan获得的数据", d)
	}()

	go func() {
		//这里我再开一个协程,那么我上文的第二次从chan获得数据就可以疏通了！
		strings <- "456"
	}()
	//此处我是为了打印我看到的内容，故意做的时间sleep
	time.Sleep(time.Second * 100)
}

//再抽象
func AbstractWriteChan(s chan<- string) {
	s <- "123"
}

//再抽象读取
func AbstractReadChan(s <-chan string) {
	fmt.Println(<-s)
}

//这里的函数会永久阻塞,如果chan没有数据的话
func blockRun(strings chan string) {
	j, ok := <-strings
	fmt.Println(j)
	fmt.Println(ok)
}

//for-break--test  就类似与java的for--break break以后我就全部退出了
func forBreak() {
	// 接收方。
	for {
		fmt.Println("Receiver: closed channel")
		break
	}
	fmt.Printf("Receiver: received an element: %v\n", 1)
	fmt.Println("End.")
}
