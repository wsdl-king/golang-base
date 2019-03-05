package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

//极客时间43讲 bufio下半程----这章还需要再看几遍,领略其中的精髓
//主要问题 bufio.Reader类型的4种读取方法
func main() {

	// Peek  特点 就是偷嘛, 然后就读取了缓冲区的数据 但是也不会更改已读计数的值
	// Peek的话 你需要区别 是否Peek值大于缓冲区的值,报的错是不一样的
	// Read  特点 读取并且增加已读计数,当缓冲区没有数据,我会直接跳过缓冲区,从底层读取器中读取数据
	// 上面两个 把数据写入缓冲区的时候,都会及时的更新已经写计数的值
	// ReadSlice 读取间隔符,如果缓冲区填写满了,还是没有读到,那么我就直接的返回缓冲区的数据
	// ReadBytes  一次又一次的调用ReadSlice方法,直到找到分隔符
	// 返回切片结构的结构 都可能造成内存泄漏
	//jk84()
	//jk85()
	jk86()
}

//极客时间demo84
func jk84() {

	comment := "Package bufio implements buffered I/O. " +
		"It wraps an io.Reader or io.Writer object, " +
		"creating another object (Reader or Writer) that " +
		"also implements the interface but provides buffering and " +
		"some help for textual I/O."
	basicReader := strings.NewReader(comment)
	fmt.Printf("The size of basic reader: %d\n", basicReader.Size())
	fmt.Println()

	// 示例1。
	fmt.Println("New a buffered reader ...")
	//创建缓冲区 默认大小是4096
	reader1 := bufio.NewReader(basicReader)
	fmt.Printf("The default size of buffered reader: %d\n", reader1.Size())
	// 此时reader1的缓冲区还没有被填充。 初始化以后我是不会填充的, 然后我reader.Buffered标示缓冲区里已读数
	fmt.Printf("The number of unread bytes in the buffer: %d\n", reader1.Buffered())
	fmt.Println()

	//// 示例2。
	bytes, err := reader1.Peek(7)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("Peeked contents(%d): %q\n", len(bytes), bytes)
	//填充了,但是没有改变已读数
	fmt.Printf("The number of unread bytes in the buffer: %d\n", reader1.Buffered())
	fmt.Println()
	//
	//// 示例3。
	buf1 := make([]byte, 7)
	n, err := reader1.Read(buf1)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("Read contents(%d): %q\n", n, buf1)
	//改变已读数
	fmt.Printf("The number of unread bytes in the buffer: %d\n", reader1.Buffered())
	fmt.Println()
	//
	//// 示例4。
	//重置底层数据
	fmt.Printf("Reset the basic reader (size: %d) ...\n", len(comment))
	basicReader.Reset(comment)
	//重置缓冲区
	fmt.Printf("Reset the buffered reader (size: %d) ...\n", reader1.Size())
	reader1.Reset(basicReader)
	peekNum := len(comment) + 1
	fmt.Printf("Peek %d bytes ...\n", peekNum)
	bytes, err = reader1.Peek(peekNum)
	if err != nil {
		//这里报错了,因为我fill以后还是不满足peek的值,但是并未超过缓冲区的大小
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("The number of peeked bytes: %d\n", len(bytes))
	fmt.Println()

	// 示例5。
	fmt.Printf("Reset the basic reader (size: %d) ...\n", len(comment))
	//需要接收数据的 然后 重置读位置
	basicReader.Reset(comment)
	size := 300
	fmt.Printf("New a buffered reader with size %d ...\n", size)
	reader2 := bufio.NewReaderSize(basicReader, size)
	peekNum = size + 1
	fmt.Printf("Peek %d bytes ...\n", peekNum)

	//超过了缓冲区的大小,缓冲区是300 但是我peek是301 所以会报  bufio: buffer full
	bytes, err = reader2.Peek(peekNum)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("The number of peeked bytes: %d\n", len(bytes))
}

//极客时间demo85
func jk85() {
	comment := "Package bufio implements buffered I/O. " +
		"It wraps an io.Reader or io.Writer object, " +
		"creating another object (Reader or Writer) that " +
		"also implements the interface but provides buffering and " +
		"some help for textual I/O."
	////创建普通的Reader对象
	//basicReader := strings.NewReader(comment)
	//fmt.Printf("The size of basic reader: %d\n", basicReader.Size())
	//
	//size := 300
	//fmt.Printf("New a buffered reader with size %d ...\n", size)
	////缓冲区, 然后设置了缓冲区的大小是300
	//reader1 := bufio.NewReaderSize(basicReader, size)
	//fmt.Println()
	//
	//fmt.Print("[ About 'Peek' method ]\n\n")
	//// 示例1。
	//peekNum := 38
	//fmt.Printf("Peek %d bytes ...\n", peekNum)
	//bytes, err := reader1.Peek(peekNum)
	//if err != nil {
	//	fmt.Printf("error: %v\n", err)
	//}
	//fmt.Printf("Peeked contents(%d): %q\n", len(bytes), bytes)
	//fmt.Printf("The number of unread bytes in the buffer: %d\n", reader1.Buffered())
	//fmt.Println()
	//
	//fmt.Print("[ About 'Read' method ]\n\n")
	//// 示例2。
	//readNum := 38
	//buf1 := make([]byte, readNum)
	//fmt.Printf("Read %d bytes ...\n", readNum)
	//n, err := reader1.Read(buf1)
	//if err != nil {
	//	fmt.Printf("error: %v\n", err)
	//}
	//fmt.Printf("Read contents(%d): %q\n", n, buf1)
	//fmt.Printf("The number of unread bytes in the buffer: %d\n", reader1.Buffered())
	//fmt.Println()
	//
	//fmt.Print("[ About 'ReadSlice' method ]\n\n")
	//
	//
	////// 示例3。
	//fmt.Println("Reset the basic reader ...")
	//basicReader.Reset(comment)
	//fmt.Println("Reset the buffered reader ...")
	//reader1.Reset(basicReader)
	//fmt.Println()

	//delimiter := byte('(')
	//fmt.Printf("Read slice with delimiter %q...\n", delimiter)
	//line, err := reader1.ReadSlice(delimiter)
	//if err != nil {
	//	fmt.Printf("error: %v\n", err)
	//}
	//fmt.Printf("Read contents(%d): %q\n", len(line), line)
	//fmt.Printf("The number of unread bytes in the buffer: %d\n", reader1.Buffered())
	//fmt.Println()

	//找不到我就填充满缓冲区 然后全都读取
	//delimiter := byte('[')
	//fmt.Printf("Read slice with delimiter %q...\n", delimiter)
	//line, err := reader1.ReadSlice(delimiter)
	//if err != nil {
	//	fmt.Printf("error: %v\n", err)
	//}
	//fmt.Printf("Read contents(%d): %q\n", len(line), line)
	//fmt.Printf("The number of unread bytes in the buffer: %d\n", reader1.Buffered())
	//fmt.Println()

	// 示例4。
	//fmt.Println("Reset the basic reader ...")
	//basicReader.Reset(comment)
	//size = 200
	//fmt.Printf("New a buffered reader with size %d ...\n", size)
	//reader2 := bufio.NewReaderSize(basicReader, size)
	//fmt.Println()
	//
	//delimiter := byte('[')
	//fmt.Printf("Read slice with delimiter %q...\n", delimiter)
	//line, err := reader2.ReadSlice(delimiter)
	//if err != nil {
	//	fmt.Printf("error: %v\n", err)
	//}
	//fmt.Printf("Read contents(%d): %q\n", len(line), line)
	//fmt.Printf("The number of unread bytes in the buffer: %d\n", reader2.Buffered())
	//fmt.Println()
	//
	//fmt.Print("[ About 'ReadBytes' method ]\n\n")
	//// 示例5。
	//fmt.Println("Reset the basic reader ...")
	//basicReader.Reset(comment)
	//size = 200
	//fmt.Printf("New a buffered reader with size %d ...\n", size)
	//reader3 := bufio.NewReaderSize(basicReader, size)
	//fmt.Println()
	//
	//delimiter := byte('[')
	//fmt.Printf("Read bytes with delimiter %q...\n", delimiter)
	////读满填充 就不放弃
	//line, err := reader3.ReadBytes(delimiter)
	//if err != nil {
	//	fmt.Printf("error: %v\n", err)
	//}
	//fmt.Printf("Read contents(%d): %q\n", len(line), line)
	//fmt.Printf("The number of unread bytes in the buffer: %d\n", reader3.Buffered())
	//fmt.Println()
	//
	//// 示例6和示例7。
	//fmt.Print("[ About contents leak ]\n\n")
	showContentsLeak(comment)
}
func showContentsLeak(comment string) {
	// 示例6。
	basicReader := strings.NewReader(comment)
	fmt.Printf("The size of basic reader: %d\n", basicReader.Size())

	size := len(comment)
	fmt.Printf("New a buffered reader with size %d ...\n", size)
	reader4 := bufio.NewReaderSize(basicReader, size)
	fmt.Println()

	peekNum := 7
	fmt.Printf("Peek %d bytes ...\n", peekNum)
	bytes, err := reader4.Peek(peekNum)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("Peeked contents(%d): %q\n", len(bytes), bytes)
	fmt.Println()

	// 只要扩充一下之前拿到的字节切片bytes，
	// 就可以用它来读取甚至修改缓冲区中的后续内容。
	bytes = bytes[:cap(bytes)]
	fmt.Printf("The all of the contents in the buffer:\n%q\n", bytes)
	fmt.Println()

	blank := byte(' ')
	fmt.Println("Set blanks into the contents in the buffer ...")
	for _, i := range []int{55, 56, 57, 58, 66, 67, 68} {
		bytes[i] = blank
	}
	fmt.Println()

	peekNum = size
	fmt.Printf("Peek %d bytes ...\n", peekNum)
	bytes, err = reader4.Peek(peekNum)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("Peeked contents(%d):\n%q\n", len(bytes), bytes)
	fmt.Println()

	// 示例7。
	// ReadSlice方法也存在相同的问题。
	delimiter := byte(',')
	fmt.Printf("Read slice with delimiter %q...\n", delimiter)
	line, err := reader4.ReadSlice(delimiter)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("Read contents(%d): %q\n", len(line), line)
	fmt.Println()

	line = line[:cap(line)]
	fmt.Printf("The all of the contents in the buffer:\n%q\n", line)
	fmt.Println()

	underline := byte('_')
	fmt.Println("Set underlines into the contents in the buffer ...")
	for _, i := range []int{89, 92, 103} {
		line[i] = underline
	}
	fmt.Println()

	peekNum = size
	fmt.Printf("Peek %d bytes ...\n", peekNum)
	bytes, err = reader4.Peek(peekNum)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("Peeked contents(%d): %q\n", len(bytes), bytes)
}

func jk86() {
	comment := "Writer implements buffering for an io.Writer object. " +
		"If an error occurs writing to a Writer, " +
		"no more data will be accepted and all subsequent writes, " +
		"and Flush, will return the error. After all data has been written, " +
		"the client should call the Flush method to guarantee all data " +
		"has been forwarded to the underlying io.Writer."
	basicWriter1 := &strings.Builder{}
	//
	size := 300
	//fmt.Printf("New a buffered writer with size %d ...\n", size)
	////创建写缓冲区
	writer1 := bufio.NewWriterSize(basicWriter1, size)
	//fmt.Println()
	//
	//// 示例1。
	//begin, end := 0, 53
	//fmt.Printf("Write %d bytes into the writer ...\n", end-begin)
	//writer1.WriteString(comment[begin:end])
	//fmt.Printf("The number of buffered bytes: %d\n", writer1.Buffered())
	//fmt.Printf("The number of unused bytes in the buffer: %d\n",
	//	writer1.Available())
	//fmt.Println("Flush the buffer in the writer ...")
	////手动写到底层容器,所以我缓冲区可用的长度就复原了,然后,然后我缓冲区也没有可写数据了
	//writer1.Flush()
	//fmt.Printf("The number of buffered bytes: %d\n", writer1.Buffered())
	//fmt.Printf("The number of unused bytes in the buffer: %d\n",
	//	writer1.Available())
	//fmt.Println()

	//// 示例2。
	//begin, end = 0, 326
	//fmt.Printf("Write %d bytes into the writer ...\n", end-begin)
	//writer1.WriteString(comment[begin:end])
	//fmt.Printf("The number of buffered bytes: %d\n", writer1.Buffered())
	//fmt.Printf("The number of unused bytes in the buffer: %d\n",
	//	writer1.Available())
	//fmt.Println("Flush the buffer in the writer ...")
	//writer1.Flush()
	//fmt.Println()
	//
	//// 示例3。
	basicWriter2 := &bytes.Buffer{}
	fmt.Printf("Reset the writer with a bytes buffer(an implementation of io.ReaderFrom) ...\n")
	writer1.Reset(basicWriter2)
	reader := strings.NewReader(comment)
	fmt.Println("Read data from the reader ...")
	//超过缓冲区的大小 直接写入数据容器
	writer1.ReadFrom(reader)
	fmt.Printf("The number of buffered bytes: %d\n", writer1.Buffered())
	fmt.Printf("The number of unused bytes in the buffer: %d\n",
		writer1.Available())

}
