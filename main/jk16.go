package main

import (
	"fmt"
	"time"
)

//极客时间第16讲  协程方面的知识复习及补充
// 至理名言: 不要用共享内存的方式来进行数据通讯,应该用数据通讯的方式来共享内存
func main() {
	//go1()
	//go2()
	//go3()
	//go4()
}

//  我运行我的主main 你运行你的小go 但是我主线程运行完以后要停顿 ..偶买噶
//	小go去那外层通信函数的值,拿到的可能都是10 也可能很多10 夹杂着幸存者
// 小 go包装成G的过程中要去主线程取值的,但是一般你取到的都是10.........
func go1() {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(time.Second * 1)
}

//第二种 go 函数的运行 顺序打印 0,1,2,3,4......
func go2() {
	// 运行的go 不知道是哪个go
	// 我运行完以后for以后,我就停顿一秒 这时候我开启1个go,这个go已经加入了G队列,然后因为主main停顿
	//所以我这一个go 会取值 打印  结果为顺序打印
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
		time.Sleep(time.Second * 1)
	}
}

//第三种go 函数的运行 其实这个什么也不打印
func go3() {
	//这个就是 我运行我的主main 你运行你的小go
	//我主main 一定是很快的运行完的,你的小go是个渣渣 包装成G到线程里 太慢了 我主线程不等你了 不会打印任何
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
}

//乱序打印
func go4() {
	for i := 0; i < 10; i++ {
		// 每次运行完一个for 都要将i值取出来,等待G包装 到G队列
		go func(f int) {
			//这里只有取内部 go函数的参数 才是乱序
			fmt.Println(f)
		}(i) // i是形参,这个形参是读取的外部主程序的i 然后映射到f
	}
	time.Sleep(time.Microsecond * 14)
}
