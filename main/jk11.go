package main

import (
	"fmt"
	"math/rand"
	"time"
)

//极客时间第11章,选择器练习
func main() {
	//strings := make(chan string, 1)
	//单纯的入门选择器
	//Selector(strings)
	//goToLoopSelector 使用的是 goto var 的关键字 退出外层的for循环
	//goToLoopSelector(strings)
	//这里测试的是 break loop 退出外层的for循环
	//BreakLoopSelector(strings)
	//example测试了 for-select-case 及 break的情况
	//example1()
	//example2 测试了 chan 提前关闭后的情况
	//example2()
	//测试了for-range对chan的简单应用
	//ForRange()
	//补充ran.intn的知识
	addRandSeed()
}

func Selector(strings chan string) {
	for {
		//select 选择器里面的break只是退出选择器,并没有退出for循环
		select {
		case v1 := <-strings:
			fmt.Print(v1)
			break
		default:
			fmt.Println("not  value")
			break
		}

	}
}

//注意注意,select选择器,它只会选择一次,通常搭配外层的for循环进行使用,也就是说,如果只是select,
//select 只会挑选那个没有阻塞的chan进行输出,然后退出
func goToLoopSelector(strings chan string) {
	for {
		//select 选择器里面的break只是退出选择器,并没有退出for循环
		//这里使用 goto loop进行操作 可以退出for循环
		select {
		case v1 := <-strings:
			fmt.Print(v1)
			goto loop
		default:
			fmt.Println("not  value")
			goto loop
		}
	}
loop:
	fmt.Println("我退出了for循环")
}

func BreakLoopSelector(strings chan string) {
loop:
	for {
		//select 选择器里面的break只是退出选择器,并没有退出for循环
		//这里使用 goto loop进行操作 可以退出for循环
		select {
		case v1 := <-strings:
			fmt.Print(v1)
			break loop
		default:
			fmt.Println("not  value")
			break loop
		}
	}
	fmt.Println("我退出了for循环")
}

// 示例1。
//select --case 中加不加break是没有区别的,反正如果外部有for,我也是不会退出for的哈哈哈
//此例子 case 中的break 其实你可以随便的加与不加  不加的话 执行完毕以后就会退出的
//但是 如果 select-case 中 的内置代码块有break的话,那加与不加 就特别的有区别了！ 你可以看example2
func example1() {
	// 准备好几个通道数组
	intChannels := [3]chan int{
		make(chan int, 1),
		make(chan int, 1),
		make(chan int, 1),
	}
	var index int
	// 随机选择一个通道，并向它发送元素值。
	go func() {
		for {
			index = rand.Intn(3)
			intChannels[index] <- index
		}
	}()
	fmt.Printf("The index: %d\n", index)
	intChannels[index] <- index
	// 哪一个通道中有可取的元素值，哪个对应的分支就会被执行。
	for {
		time.Sleep(time.Second * 2)
		select {
		case <-intChannels[0]:
			fmt.Println("The first candidate case is selected.")
			break
		case <-intChannels[1]:
			fmt.Println("The second candidate case is selected.")
			break
		case elem := <-intChannels[2]:
			fmt.Printf("The third candidate case is selected, the element is %d.\n", elem)
			break
		default:
			fmt.Println("No candidate case is selected!")
			break
		}
	}

}

// 示例2。
func example2() {
	intChan := make(chan int)
	//  关闭通道以后,再从通道读取数据都是0,敲黑板画重点了!
	close(intChan)
	for {
		fmt.Println(<-intChan)
	}
}

//注意,这里的操作顺序,如果 协程放在 strings<-"123" 后面的话 永久阻塞,因为是一条线执行的程序,没有配对的chan操作
//所以,这里需要把协程放在strings<-"123"前面
func ForRange() {
	strings := make(chan string)
	go func() {
		for d := range strings {
			fmt.Println(d)
		}
	}()
	strings <- "123"
	strings <- "1232"
	time.Sleep(time.Second)
}

//补充一下randSeed和rand.Int部分的内容
//笔记: tip1:你看昂,如果我不加rand.seed函数,那么我rand.Ints单线程运行下永远是第一次的值,但是我for一直运行的话,就是好多次值了。
// tip2: 再看昂,如果我加了rand.seed 我多次单线程运行就可以获得不同的值,for一直运行也是不同的值,不过for循环太威猛,,导致很多次一样哈哈哈,你可以sleep
func addRandSeed() {
	//rand.Seed(time.Now().Unix())
	//intn := rand.Intn(3)
	//fmt.Println(intn)

	for {
		time.Sleep(time.Second * 1)
		rand.Seed(time.Now().Unix())
		intn := rand.Intn(3)
		fmt.Println(intn)
	}
}
