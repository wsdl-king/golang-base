package main

import (
	"errors"
	//"errors"
	"fmt"
	"math"
)

//极客时间21讲, panic函数,recover函数 还有defer语句
// panic 英文解释 是恐慌, 上节课讲述的是错误 这节课要展示的是panic(异常) 老子意料之外的异常草
func main() {
	//squareRoot := int(math.Sqrt(float64(9)))
	//fmt.Print(squareRoot)
	//GetPrimes(49)

	//不捕获错误 老子直接给你终止程序 然后逆向栈出打印错误内容
	fmt.Println("进入main")
	//caller1()
	recoverT()
	fmt.Println("exit main")
}

//现在开始把我要学的一部分给列出来

//参照极客时间git article20 -- 这个我也不知道干啥的 %>_<% 我只学的是 panic recover defer的东东呀 臣妾坐不了主啊
// GetPrimes 用于获取小于或等于参数max的所有质数。 有点意思 很巧妙
// 本函数使用的是爱拉托逊斯筛选法（Sieve Of Eratosthenes）。
func GetPrimes(max int) []int {
	if max <= 1 {
		return []int{}
	}
	marks := make([]bool, max)
	var count int
	squareRoot := int(math.Sqrt(float64(max)))
	//质数检查逻辑 从2开始
	for i := 2; i <= squareRoot; i++ {
		if marks[i] == false {
			//保存切片位为false的才是质数
			for j := i * i; j < max; j += i {
				if marks[j] == false {
					marks[j] = true
					count++
				}
			}
		}
	}
	primes := make([]int, 0, max-count)
	for i := 2; i < max; i++ {
		if marks[i] == false {
			primes = append(primes, i)
		}
	}
	return primes
}

func caller1() {
	fmt.Println("Enter function caller1.")
	caller2()
	fmt.Println("Exit function caller1.")
}

func caller2() {
	fmt.Println("Enter function caller2.")
	s1 := []int{0, 1, 2, 3, 4}
	panic(errors.New("something wrong")) // 认为的触发panic
	e5 := s1[5]
	_ = e5
	fmt.Println("Exit function caller2.")
}

// recover练习, 其中有 panic,defer,recover
func recoverT() {
	fmt.Println("Enter function main.")

	// defer func 最后执行
	defer func() {
		fmt.Println("Enter defer function.")

		// recover函数的正确用法。
		if p := recover(); p != nil {
			fmt.Printf("panic: %s\n", p)
		}

		fmt.Println("Exit defer function.")
	}()

	// recover函数的错误用法。
	fmt.Printf("no panic: %v\n", recover())

	// 引发panic。
	//panic(errors.New("something wrong"))

	// recover函数的错误用法。
	// Unreachable 不可达的意思
	p := recover()
	fmt.Printf("panic: %s\n", p)

	fmt.Println("Exit function main.")

}
