package main

import (
	"bytes"
	"fmt"
)

//极客时间38讲,讲述的bytes包与字节串操作
//  strings包主要面向的是Unicode字符和经过UTF-8编码的字符串,bytes包面对的则主要是字节和字节切片
// strings包和bytes包最大的不同就是,我可以通过size-len 得到已读的strings,但是无法得到已读的bytes
// strings和bytes的len都是未读的长度,bytes包里的诸多操作都是基于 "已读长度"的
func main() {
	//simplebuff()
	//truncateTest()
	//simpleTrune()
	simple2()
}

//简单的byte buffer例子--字节序列缓冲区
func simplebuff() {
	var buffer1 bytes.Buffer
	//39个字符
	contents := "Simple byte buffer for marshaling data."
	fmt.Printf("Writing contents %q ...\n", contents)
	buffer1.WriteString(contents)
	// 切片扩容策略 2的幂次方 32*2=64
	fmt.Printf("The length of buffer: %d\n", buffer1.Len())
	fmt.Printf("The capacity of buffer: %d\n", buffer1.Cap())

	//读取7个字节
	p1 := make([]byte, 7)
	n, _ := buffer1.Read(p1)
	fmt.Printf("%d bytes were read. (call Read)\n", n)
	// len返回的是未读的字节长度和strings一样
	fmt.Printf("The length of buffer: %d\n", buffer1.Len())
	fmt.Printf("The capacity of buffer: %d\n", buffer1.Cap())
}

//这里练习截断方法
func truncateTest() {
	var buffer1 bytes.Buffer
	contents := "abcde"
	buffer1.WriteString(contents)
	p1 := make([]byte, 1)
	buffer1.Read(p1)
	//在截断时,需要保留头部3个字节
	// 这里的头部是未读部分的头部,我先读了一个 然后保存3个,一旦超过4就不行了
	buffer1.Truncate(4)
	fmt.Println(buffer1.Len())
	fmt.Println(buffer1)
}

//这里练习回退方法
func simpleTrune() {
	var buffer1 bytes.Buffer
	contents := "大帅哥"
	buffer1.WriteString(contents)
	//p1 := make([]byte, 8)
	//buffer1.Read(p1)
	buffer1.ReadRune()
	fmt.Println(buffer1.Len())
	//如果我没有读过就回退,我直接给你返回一个error
	//unreadByte := buffer1.UnreadByte()
	//fmt.Println(unreadByte)
	//这里我回退的是上一次被读取的Unicode字符所占用的字节数
	buffer1.UnreadRune()
	fmt.Println(buffer1.Len())
}
func readFromTest() {
	var buffer1 bytes.Buffer
	contents := "abcde"
	buffer1.WriteString(contents)
	//从一个实现io.Reader接口的r，把r里的内容读到缓冲器里，n返回读的数量
	//buffer1.ReadFrom()
}

//
func simple2() {
	// 示例1。
	var contents string
	//这样默认生成的长度就是8
	buffer1 := bytes.NewBufferString("1")
	//var buffer1 bytes.Buffer
	//这样默认生成的是64
	//buffer1.WriteString("a")

	fmt.Printf("The length of new buffer with contents %q: %d\n",
		contents, buffer1.Len())
	fmt.Printf("The capacity of new buffer with contents %q: %d\n",
		contents, buffer1.Cap())
	fmt.Println()

	contents = "12345"
	fmt.Printf("Write contents %q ...\n", contents)
	buffer1.WriteString(contents)
	fmt.Printf("The length of buffer: %d\n", buffer1.Len())
	fmt.Printf("The capacity of buffer: %d\n", buffer1.Cap())
	fmt.Println()
}
