package main

import (
	"bytes"
	"fmt"
)

//极客时间39章的练习,bytes包与字节串操作
//课后习题: strings.builder和bytes.buffer的String方法做对比,前者更高效,源码利用了指针操作.
func main() {
	//simpleT()
	ncxl()
}

//扩容策略---这里先简单记录一下,剩余长度够 我就不扩容,如果不够就扩容2m+n
func simpleT() {
	//1.剩余长度
	var contents string
	//这样默认生成的长度就是8
	buffer1 := bytes.NewBufferString("1")
	fmt.Printf("The length of new buffer with contents %q: %d\n",
		contents, buffer1.Len())
	fmt.Printf("The capacity of new buffer with contents %q: %d\n",
		contents, buffer1.Cap())
	fmt.Println()
	//这里演示我直接增加8个字符
	// 进行扩容 原来容量*2+新的长度n
	buffer1.WriteString("abcabcab1")
	fmt.Printf("The length of new buffer with contents %q: %d\n",
		contents, buffer1.Len())
	fmt.Printf("The capacity of new buffer with contents %q: %d\n",
		contents, buffer1.Cap())
}

// 问题2 bytes.buffer中的方法造成的内存泄露
// 我通过bytes方法拿到了当时未读的字节切片,然后我对当时的字节切片做切片操作就可以以不正当的方法拿到一个切片
// 但是底层都是引用的一个数组 我曹 这特么太危险了 想要解决你就深度拷贝 直接for循环然后copy
func ncxl() {
	contents := "ab"
	buffer1 := bytes.NewBufferString(contents)
	fmt.Printf("The capacity of new buffer with contents %q: %d\n",
		contents, buffer1.Cap()) // 内容容器的容量为：8。
	//Next的方法就返回一个由原来buffer的下 2 个字节组成的切片
	unreadBytes := buffer1.Next(2)
	//unreadBytes := buffer1.Bytes()
	fmt.Printf("The unread bytes of the buffer: %v\n", unreadBytes) // 未读内容为：[97 98]。

	buffer1.WriteString("cdefg")
	fmt.Printf("The capacity of buffer: %d\n", buffer1.Cap()) // 内容容器的容量仍为：8。
	unreadBytes = unreadBytes[:cap(unreadBytes)]
	fmt.Printf("The unread bytes of the buffer: %v\n", unreadBytes) // 基于前面获取到的结果值可得，未读内容为：[97 98 99 100 101 102 103 0]。
	//危险了 草..
	//直接对切片进行修改 索引6处 我直接增加一个byte
	unreadBytes[len(unreadBytes)-2] = byte('X') // 'X'的 ASCII 编码为 88。
	//未读内容也发生了变化, 所以 这个东西要慎用
	fmt.Printf("The unread bytes of the buffer: %v\n", buffer1.Bytes()) // 未读内容变为了：[97 98 99 100 101 102 88]。
}
