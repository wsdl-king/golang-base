package main

import (
	"context"
	"fmt"
	"time"
	//"time"
)

//极客时间32章,有关Context内容的讲解
//  笔记, Context就是一颗树,父子树, 恩.. 还有   ctx.Done() 是除了context.value意外都存在的
// 内部实现就是一个管道,可以管控goroutine的分发,当我的<-ctx.Done()返回值了,也就意味着 超时了,或者手动cancel了
// 我所在的goroutine也就不阻塞了
// 总结: 含有数据的不可以被撤销,不含数据的可以被撤销,父子Context之间的传播是深度优先
func main() {
	//deadlockContext()
	//timeoutContext()
	deadContext()
	//simpleContextValue()
}

//简单的Context的应用
func deadlockContext() {
	//创建根节点,根节点是不带任何数据的
	//返回两个值,一个是当前节点,因为我现在是background所以我就是空的父节点,另一个是取消函数,当我取消以后,
	// 我的当前ctx的done也就不阻塞了
	rootCtx, _ := context.WithCancel(context.Background())
	_, cancel := context.WithCancel(rootCtx)
	go func() {
		defer func() {
			//fmt.Println("父类取消其子节点也会跟着取消!")
			// 我的子节点取消,但是我的父类进行阻塞
			cancel()
		}()
	}()
	//这边的话就会死锁
	<-rootCtx.Done()
	e := rootCtx.Err()
	// 表明死你个傻逼自己取消的
	fmt.Println("我的错误是", e)
}

//我自己释放的goroutine
//超时Context 然后我自己也可以取消
func timeoutContext() {
	//三秒超时,管控我的子 goroutine
	timeout, cancel := context.WithTimeout(context.Background(), time.Second*3)
	go func() {
		defer func() {
			time.Sleep(time.Second * 2)
			fmt.Println("超时2秒")
			//这里老子自己释放
			cancel()
			time.Sleep(time.Second * 2)
		}()

	}()
	//超时自动释放
	<-timeout.Done()
	e := timeout.Err()
	// 表明死你个傻逼自己取消的
	fmt.Println("我的错误是", e)
}

//这个是设置具体的超时时间,到时间以后退出
func deadContext() {
	rootCtx, _ := context.WithCancel(context.Background())
	//三秒以后超时
	add := time.Now().Add(time.Second * 3)
	fmt.Println(add)
	deadline, _ := context.WithDeadline(rootCtx, add)

	go func() {
		fmt.Println("老子定义一个不可取消的超时老子认为阻塞2秒")
		//我自己退出
		//阻塞2s
		time.Sleep(time.Second * 2)
		fmt.Println("阻塞2秒完毕")
	}()
	//等待阻塞
	<-deadline.Done()
	//打印出我的终止时间
	deadline2, ok := deadline.Deadline()
	fmt.Println(deadline2)
	fmt.Println(ok)
	e := deadline.Err()
	// 表明死你个傻逼自己取消的
	fmt.Println("我的错误是", e)
}

//你可以看到,在使用context.value的方法的时候,我的子类可以去找父类的值,但是我的父类是不可以找子类的值的
// 没有提供通过选定key然后改变value值的方法,我们只能新建 %>_<%,撤销的话 就需要用到父类,并且父类也是可撤销的!
func simpleContextValue() {
	//可撤销的,返回两个值,你一看就知道
	rootCtx, _ := context.WithCancel(context.Background())
	//设计root的一个value
	//不可撤销的,返回一个值, 撤销信号在父类传播的时候,遇到他们会直接跨过,然后将信号传递给他们的子值
	withValue := context.WithValue(rootCtx, "root", "value")
	i := context.WithValue(withValue, "qws", "123")
	//我衍生出一个非value的context
	h, _ := context.WithCancel(i)
	//我用父节点找子节点的key,是得不到数据的
	fmt.Println(withValue.Value("qws"))
	//我用子节点找父节点的值
	fmt.Println(i.Value("root"))
	//子节点自己的值
	fmt.Println(i.Value("qws"))

	//打印最底层节点查找的value数据--注意只可以向上取值
	fmt.Println("一个没有数据的Context,但是他们存在相应的关系", h.Value("qws"))
}

//极客时间的例子
