package main

import (
	"fmt"
	"io"
	"strings"
)

//极客时间37讲,讲的是strings包与字符串操作
// ps: 一般在包里调用New***函数生成新的指针对象
// 与string值相比,string.Builder类型的值有哪些优势?
// 回答: 已存在的内容不可变,但可以拼接更多内容.  减少了内存分配和内存拷贝的次数 . 可将内容重置,可重用值.
// builder 使用后不能扩容,否则会panic. builder内部有一个字节数组指针,外部是不可以进行插手的,,同时,它也是线程不安全的,
// 我append的话 是在一块新的内存区域
// Reader的使用,reader有个读指针,下次读会从指针以后开始读取,可以使用size-len的方法计算出读指针,也可以用seek函数增加读指针.
func main() {
	//buildergrow()
	//simpleBuilder()
	//builderRest()
	//buildPass()
	simplieReader()
}

//string.builder扩容,扩容用到的方法是grow(n) 扩荣策略是当前剩余长度不足n就会扩容2倍+n,如果大于等于n,那么不会扩容
func buildergrow() {
	var builder1 strings.Builder
	fmt.Println("Grow the builder ...")
	//没有内容不会扩容
	builder1.Grow(10)
	fmt.Printf("The first  contents in the builder is %d.\n", builder1.Len())
	fmt.Println()
	builder1.WriteString("qws")
	fmt.Printf("The second   contents in the builder is %d.\n", builder1.Len())
}
func simpleBuilder() {
	// 示例1。
	var builder1 strings.Builder
	builder1.WriteString("A Builder is used to efficiently build a string using Write methods.")
	fmt.Printf("The first output(%d):\n%q\n", builder1.Len(), builder1.String())
	fmt.Println()
	//空字节
	builder1.WriteByte(' ')
	//builder1.WriteString("It minimizes memory copying. The zero value is ready to use.")
	//builder1.Write([]byte{'\n', '\n'})
	//builder1.WriteString("Do not copy a non-zero Builder.")
	fmt.Printf("The second output(%d):\n\"%s\"\n", builder1.Len(), builder1.String())
	fmt.Println()

}
func builderRest() {
	var builder1 strings.Builder
	builder1.WriteString("A Builder is used to efficiently build a string using Write methods.")
	fmt.Printf("The first output(%d):\n%q\n", builder1.Len(), builder1.String())
	// reset清空值
	builder1.Reset()
	fmt.Printf("The end output(%d):\n%q\n", builder1.Len(), builder1.String())
}

//builder传递
func buildPass() {
	// 示例1.
	var builder1 strings.Builder
	builder1.Grow(1)

	//值参数传递builder会引发panic
	f1 := func(b strings.Builder) {
		//b.Grow(1) // 这里会引发panic。
	}
	f1(builder1)
	//chan值传递的时候也会出现问题
	ch1 := make(chan strings.Builder, 1)
	ch1 <- builder1
	builder2 := <-ch1
	//builder2.Grow(1) // 这里会引发panic。
	_ = builder2

	//直接值传递也会panic
	builder3 := builder1
	//builder3.Grow(1) // 这里会引发panic。
	_ = builder3

	// 示例2。
	//这里面是不会panic的,底层指针指向同一个数组
	f2 := func(bp *strings.Builder) {
		(*bp).Grow(1) // 这里虽然不会引发panic，但不是并发安全的。
		builder4 := *bp
		//指针值传递,也会panic
		//builder4.Grow(1) // 这里会引发panic。
		_ = builder4
	}
	f2(&builder1)

	//重置为0以后,传递值扩容不会引起panic
	builder1.Reset()
	builder5 := builder1
	builder5.Grow(1) // 这里不会引发panic。
}

//builder的读取
func simplieReader() {
	// 示例1。 声明一个Reader指针值
	reader1 := strings.NewReader(
		"NewReader returns a new Reader reading from s. " +
			"It is similar to bytes.NewBufferString but more efficient and read-only.")
	fmt.Printf("The size of reader: %d\n", reader1.Size())
	//目前的读索引是0,未读任何数字, 计算读索引需要用到 reader.size-reader.len
	fmt.Printf("The reading index in reader: %d\n",
		reader1.Size()-int64(reader1.Len()))
	//读了47个值,然后我的读指针到47了
	buf1 := make([]byte, 47)
	n, _ := reader1.Read(buf1)
	fmt.Printf("%d bytes were read. (call Read)\n", n)
	fmt.Printf("The reading index in reader: %d\n",
		reader1.Size()-int64(reader1.Len()))
	//这是剩余长度 ,Size是总大小
	fmt.Println(reader1.Len())
	fmt.Println()

	// 示例2。
	buf2 := make([]byte, 21)
	offset1 := int64(64)
	n, _ = reader1.ReadAt(buf2, offset1)
	fmt.Printf("%d bytes were read. (call ReadAt, offset: %d)\n", n, offset1)
	//不影响我原来的读索引
	fmt.Printf("The reading index in reader: %d\n",
		reader1.Size()-int64(reader1.Len()))
	fmt.Println()

	// 示例3。
	offset2 := int64(17)
	//手动计算读索引+17
	expectedIndex := reader1.Size() - int64(reader1.Len()) + offset2
	fmt.Printf("Seek with offset %d and whence %d ...\n", offset2, io.SeekCurrent)
	//直接调用函数使读索引+17
	readingIndex, _ := reader1.Seek(offset2, io.SeekCurrent)
	fmt.Printf("The reading index in reader: %d (returned by Seek)\n", readingIndex)
	fmt.Printf("The reading index in reader: %d (computed by me)\n", expectedIndex)

	//再读了21个字节,然后读索引+21=85
	n, _ = reader1.Read(buf2)
	fmt.Printf("%d bytes were read. (call Read)\n", n)
	fmt.Printf("The reading index in reader: %d\n",
		reader1.Size()-int64(reader1.Len()))
}
