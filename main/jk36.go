package main

import "fmt"

//极客时间36讲,讲的是unicode和字符编码
// go语言文件必须是UTF-8
// string 转换成 []rune类型的情况下,其中的字符串会被拆分成一个一个的Unicode字符
func main() {
	rangeT()

}

//笔记ASCII编码---它是由美国国家标准学会制定的 单字节字符 编码方案,可用于基于文本的数据交换.
//ASCII编码使用单个byte的二进制标示一个字符,标准的ASCII编码最高位作为奇偶校验位,而扩展的ASCII则将此位也用于表示字符
//我们所说的Unicode编码规范,实际上是另一个更加通用的,针对书面字符和文本的字符编码标准,为世界上现存的所有自然语言中的
//每一个字符,都设定了一个唯一的二进制编码.

//笔记Unicode: Unicode编码规范通常使用十六进制表示法来标示Unicode代码点的整数值.使用U+前缀 a就标示为U+61
//UTF 表示字符与字节序列之间的转换方式 -右边是编码单元.
// 问题:一个string类型的值在底层是如何被表达的
// 答案: 一个string类型的值是由一些列 相对应的Unicode代码点 的UTF-8编码值来表达的
// go语言中 一个string类型的值 可以被拆分成一个包含 多个字符 的序列,也可以被拆分成一个包含 多个字节 的序列.
// 前者 用rune(int32)元素类型的切片, 后者 用byte切片表示
// 拆成 rune切片  1个字符标示一个int32值.  拆成 byte切片 (字符的话 3byte个标示1个)

//总结局: 一个string类型的值在底层就是一个能够表达 若干个 UTF-8编码值 的 字节序列.

//range遍历字符串
func rangeT() {
	//会先把被遍历的字符串拆成一个字节序列(每一个字符都拆成一个字节序列),然后找出字节序列中的每一个UTF-8编码值
	// 遍历的字节索引 也要注意 区分 1和3 ,相邻的Unicode字符的索引值并不一定是连续的,取决于前一个Unicode字符是否为单字节字符
	str := "Go爱好者"
	for i, c := range str {
		fmt.Printf("%d: %q [% x]\n", i, c, []byte(string(c)))
	}

}

//打印字节序列吧还有区别rune和byte 切片的不同
func simplePrint() {
	str := "Go爱好者"
	fmt.Printf("The string: %q\n", str)
	fmt.Printf("  => runes(char): %q\n", []rune(str))
	fmt.Printf("  => runes(hex): %x\n", []rune(str))
	fmt.Printf("  => bytes(hex): [% x]\n", []byte(str))
}
