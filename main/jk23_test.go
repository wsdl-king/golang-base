package main

import (
	"fmt"
	"testing"
	"time"
)

//笔记
// 测试函数需要满足的条件
// 1 测试源码文件(命令行源码和库源码),必须以_test.go结尾
// 2 功能测试中 需要 Test开头并且参数为(t *testing.T)
// 3 性能测试中 需要 Benchmark开头并且参数为(b *testing.B)
// 4 示例测试中 需要以Example 为开头 参数不做要求

//极客时间23讲 --go 程序测试基础知识
//函数的命名需要_test.go
// 测试分为三种  功能测试 基准测试 示例测试
// 说一句至理名言吧
// 人是否会进步以及进步得有多快,依赖的恰恰就是对自我的否定,这包括否定的深刻与否,
// 以及否定自我的频率如何. 这其实就是"不破不立" 这个词表达的含义
func main() {

	//这里我就写示例函数吧
	ExampleCat_Name()
}

type Cat2 struct {
	name           string // 名字。
	scientificName string // 学名。
	category       string // 动物学基本分类。
}

//错误示例
//测试函数需要以Test为前缀
//这里 这个函数 两个错误 函数的名称不是Test前缀 函数的参数 没有*testing.T
// 当我运行 go test 的时候 会告诉我 检查不到test函数
//go test jk23_test.go
//ok  	command-line-arguments	0.001s [no tests to run]
func Cat_Name() {
	var cat = Cat2{"qws", "qws1", "qws2"}
	fmt.Print(cat.name)
}

//正确示例 这个是 功能测试函数  go test 可以检查到 也可以执行
func TestCat_Name(t *testing.T) {
	time.Sleep(time.Second * 1)
	cat := Cat2{"qws", "qws1", "qws2"}
	fmt.Print(cat.name)
}
func TestCat_Name2(t *testing.T) {
	cat := Cat2{"qws", "qws1", "qws2"}
	fmt.Print(cat.name)
}
func ExampleCat_Name() {
	time.Sleep(time.Second * 2)
	fmt.Print("示例函数")
}

// 性能测试 请移步到 test的包
