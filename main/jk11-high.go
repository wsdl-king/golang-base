package main

//极客时间第11高级补充章练习,这里练习了select-case的表达式
//当select-case　运行的时候,会自上而下运行一次case上的双边表达式,注意哦　每次的select case　都会运行一次双边表达式
//运行完双边表达式以后开始从 select-case里寻找合适的出口,如果所有的case不符合并且没有default,则线程永远阻塞
//直到有另外一个协程"唤醒"它.
import (
	"fmt"
	"time"
)

var channels = [3]chan int{
	nil,
	make(chan int),
	nil,
}

var numbers = []int{1, 2, 3}

func main() {
	seniorChanGo()
}

func seniorChanGo() {
	go func() {
		//for{
		a := <-channels[1]
		fmt.Println("一个协程消费管道的数据,查看select的值", a)
		//}
	}()
	seniorChan()
	time.Sleep(time.Second * 10)
}

//高级管道中select case　的应用
func seniorChan() {
	for {
		select {
		case getChan(0) <- getNumber(0):
			fmt.Println("The first candidate case is selected.")
		case getChan(1) <- getNumber(1):
			fmt.Println("The second candidate case is selected.")
		case getChan(2) <- getNumber(2):
			fmt.Println("The third candidate case is selected")
		default:
			fmt.Println("No candidate case is selected!")
		}
	}

}

func getNumber(i int) int {
	fmt.Printf("numbers[%d]\n", i)
	return numbers[i]
}

func getChan(i int) chan int {
	fmt.Printf("channels[%d]\n", i)
	return channels[i]
}
