package main

import (
	"container/list"
	ring2 "container/ring"
	"fmt"
)

func main() {
	//text1:认识list链表的一些操作
	//text2:链表的取长度的时间复杂度是O(1),得益于它特殊的Element+len的组成
	UnderStandList()
	//text3:认识环的一些操作
	//text4:Ring环的取长度时间复杂度是O(N)
	//UnderStandRing()

	//笔记:
}

func UnderStandRing() {
	// Ring 这种结构体呢 只有三部分组成 prev-Ring,next-Ring,value(interface{})
	var ring ring2.Ring
	//ring初始化一个长度为1的循环链表
	fmt.Println("初始化ring的长度", ring.Len())
	ring.Value = -1
	//此处也是延迟加载机制
	next := ring.Next()
	//此处是创建一个长度为2的环
	i := ring2.New(2)
	//创造完长度为2的环以后,我这里做了两次值的复制，注意哦这里跟链表不一样，这里没有所谓的根节点
	i.Value = 7
	i.Next().Value = 8
	fmt.Println(i.Len())
	//这里我测试下link方法,这里把两个Ring给链接起来,这里也会进行延迟加载机制
	next.Link(i)
	fmt.Println("这里next-Ring的长度为", next.Len())
	move := next.Move(2)
	//我们可以发现长度没有变化，move方法不是删除Ring的元素哦，只是移动一个指针!
	fmt.Println("移动完毕以后的新链表的长度为", move.Len())
	//这里我们看prev方法,这里也是会进行初始化的
	prev := next.Prev()
	// 我们可以发现长度没有变化，prev方法不是删除Ring的元素哦，只是往前移动一个指针，然后返回一个指针对象
	fmt.Println("prev的长度为", prev.Len())
	//这里我同链表进行link,看发生了什么，结果
	//这里减少个好理解的方法， 这里next.Move(x)是返回的新的Ring，那我就在原链表与新的Ring相链接
	//那就是 -1和7相链接，原来的链表就成为了 -1,7,8没有变化，然后呢生成的需要next一位
	//如果x是2 则原来链表就是-1,8 生成的新的就是7
	link := next.Link(next.Move(2))
	fmt.Println(link)
	fmt.Println("此时的next,就是被删除的元素", next)
	//fmt.Println("被删除的新的环链表的长度为", link.Len())
	//这里我们看unlink方法，删除后面的n%next.len的元素，返回的元素是被删除的元素
	//unlink := next.Unlink(3)
	//fmt.Println("查看被删除元素的长度", unlink.Len())
	//fmt.Println(next.Value)
	//fmt.Println(ring)

}

func UnderStandList() {
	//tip1: 认识list
	var list list.List
	//简单认识下 list的结构 Element 和 int
	//Element是双向链表，里面具有 prev 前Element ; next 后Element,*list 此元素所属的链表,interface{} 就是值嘛
	//初始化一个链表的长度是0
	fmt.Println(list)
	fmt.Println("初始化链表的长度为", list.Len())
	// tip2: list的一些常用方法
	//给 list 添加元素
	list.PushBack(1)
	pushBack := list.PushBack(2)
	fmt.Println("在链表最后插入一个element后的返回值", pushBack)
	back := list.Back()
	fmt.Println("list 最后端的元素为:", back)
	front := list.Front()
	fmt.Println("list 最前端的元素为:", front)
	//打印添加完元素的list
	fmt.Println("添加完的list", list)
	fmt.Printf("添加完链表的长度为:%d", list.Len())
	//有点意思， 哈哈哈 这个 golang中的list就是个双向链表，此例子中我添加两个元素
	//号外号外，这里的list添加元素不同于java,这里的list添加元素只能使用list.pushBack(value)这样的值哦!
	//当我添加了两个元素，你可以发现，此时的结构就像一个双向链表 而且是循环的 list.prev就是最后一个元素

	//tip 2: 此处我使用list的其他方法,这里我在指定位置最后端元素插入一个值为3
	list.InsertAfter(3, back)
	//此时链表也发生了变化,长度变成了3
	fmt.Println("打印在最后端插入的值的链表以后的链表指向的最后的值", list.Back())
	//tip3: 此处我使用list的其他方法，这里我在指定位置最前端元素插入一个值为0
	list.InsertBefore(0, front)
	fmt.Println("打印在最前端插入值的链表以后的链表指向的值", list.Front())
	//tip4:此处我使用list的其他方法，这里我在最前端元素插入一个值为-1
	list.PushFront(-1)
	fmt.Println("打印最新链表的前端值", list.Front())
	//tip5:此处我使用list的其他方法，这里我在最后端元素插入一个值为4
	list.PushBack("4")
	fmt.Println("打印最新链表的前端值", list.Back())
	element0 := list.Front()
	element1 := element0.Next()
	element2 := element1.Next()
	fmt.Println("我看一下头元素", element0)
	fmt.Println("我看一下头元素的下一个元素", element1)
	fmt.Println("我看一下原始的list", list)
	//tip6: 此处练习链表的 move方法,我将element-1元素移动到element0前，本来0在-1前面的
	//list.MoveBefore(element1,element0)
	// tip7: 我将 0 移动 1后面  入参A,B B是调整的位置 mark位置，围绕这个位置做的一些操作
	//list.MoveAfter(element1, element2)
	fmt.Println(element2)
	//tip8: 这里我学习将给定的元素移动到链表的头 0移动到表头
	//list.MoveToFront(element1)
	//tip9: 这里我学习将给定的元素移动到链表的尾部 0移动到尾部
	list.MoveToBack(element1)
}
