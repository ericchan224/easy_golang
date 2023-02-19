package main

import (
	"fmt"
	"log"
)

func main() {
	// 新建结构体，值
	g := Diy{
		A: 2,
		//b: 4.0, // 小写成员不能导出
	}
	g.Set2(1, 1)
	log.Println("", g)
}

/*
*
TODO: 值类型：值变量
引用类型：指针变量，引用变量
*
*/
func test1() {
	// a,b 是一个值
	a := 5
	b := 6

	fmt.Println("a的值：", a)

	// 指针变量 c 存储的是变量 a 的内存地址
	c := &a
	fmt.Println("a的内存地址：", c)

	// 指针变量不允许直接赋值，需要使用 * 获取引用
	// c = 4

	// 将指针变量 c 指向的内存里面的值设置为4
	*c = 4
	fmt.Println("a的值：", a)

	// 指针变量 c 现在存储的是变量 b 的内存地址
	c = &b
	fmt.Println("b的内存地址：", c)

	// 将指针变量 c 指向的内存里面的值设置为8
	*c = 8
	fmt.Println("a的值：", a)
	fmt.Println("b的值：", b)

	// 把指针变量 c 赋予 c1, c1 是一个引用变量，存的只是指针地址，他们两个现在是独立的了
	c1 := c
	fmt.Println("c的内存地址：", c)
	fmt.Println("c1的内存地址：", c1)

	// 将指针变量 c 指向的内存里面的值设置为9
	*c = 9
	fmt.Println("c指向的内存地址的值", *c)
	fmt.Println("c1指向的内存地址的值", *c1)

	// 指针变量 c 现在存储的是变量 a 的内存地址，但 c1 还是不变
	c = &a
	fmt.Println("c的内存地址：", c)
	fmt.Println("c1的内存地址：", c1)
}

/*
*

	TODO: Golang 支持我们定义自己的数据类型，结构体
	引用和值类型的结构体有何区别的?
	我们知道函数内和函数外的变量是独立的，当传参数进函数的时候，参数是值拷贝，函数里的变量被约束在函数体内，就算修改了函数里传入的变量的值，函数外也发现不了。
	但引用类型的变量，传入函数时，虽然也是传值，但拷贝的是引用类型的内存地址，可以说拷贝了一个引用，这个引用指向了函数体外的某个结构体，使用这个引用在函数里修改结构体的值，外面函数也会发现。
	如果传入的不是引用类型的结构体，而是值类型的结构体，那么会完整拷贝一份结构体，该结构体和原来的结构体就没有关系了。
	内置的数据类型切片 slice 和字典 map 都是引用类型，不需要任何额外操作，所以传递这两种类型作为函数参数，是比较危险的，开发的时候需要谨慎操作。

*
*/
type Diy struct {
	A int64   // 大写导出成员
	b float64 // 小写不可以导出
}

/**
	TODO: 结构体可以和函数绑定，也就是说这个函数只能被该结构体使用，这种函数称为结构体方法：
**/
// 引用结构体的方法，引用传递，会改变原有结构体的值
func (diy *Diy) Set(a int64, b float64) {
	diy.A = a
	diy.b = b
	return
}

// 值结构体的方法，值传递，不会改变原有结构体的值
func (diy Diy) Set2(a int64, b float64) {
	diy.A = a
	diy.b = b
	return
}

// 结构体 g 是值类型，本来不能调用 Set 方法，但是 Golang 帮忙转换了，我们毫无感知，然后值类型就变成了引用类型。同理，k 是引用类型，照样可以使用 Set2 方法。
// 前面我们也说过，函数传入引用，函数里修改该引用对应的值，函数外也会发现。
// 结构体的方法也是一样，不过范围扩散了结构体本身，方法里可以修改结构体本身，但是如果结构体是值，那么修改后，外面的世界是发现不了的。
func test2() {
	// 新建结构体，值
	g := Diy{
		A: 2,
		//b: 4.0, // 小写成员不能导出
	}
	g.Set2(1, 1)
	g.Set(1, 1)
	log.Println("", g)
}

/**
	TODO: 关键字 new 主要用来创建一个引用类型的结构体，只有结构体可以用。
	关键字 make 是用来创建和初始化一个切片或者字典。我们可以直接赋值来使用：
**/
func test() {
	_ = new(Diy)
	arr := make([]int, 0, 0)
	arr = append(arr, []int{1,2}...)
	_ = cap(arr)
}
/**
	TODO: 函数是代码片段的一个封装，方法是将函数和结构体绑定。
	但是 Golang 里面有一些内置语法，不是函数，也不是方法，比如 append，cap，len，make，这是一种语法特征。
	语法特征是高级语言提供的，内部帮你隐藏了如何分配内存等细节。
**/
