package main

import (
	"fmt"

	"github.com/chenchen224/easy_golang/basic/one/diy"
)

/**
1,包管理：go mod init， 然后通过import导入package diy
只能通过main函数提供程序入口
**/

/**
2,Golang语言可以先声明变量，再赋值，也可以直接创建一个带值的变量。如：
**/
// 声明并初始化三个值
var i, j, k = 1, 2, 3

// 我们也可以定义常量
const s = 2

func main() {
	_ = diy.Diy{}

	// 声明后再赋值
	// var i int64
	// i = 3

	// 直接赋值，创建一个新的变量
	// j := 5

	//3, 定义基本数据类型
	p := true                             // bool
	a := 3                                // int
	b := 6.0                              // float64
	c := "hi"                             // string
	d := [3]string{"1", "2", "3"}         // array，基本不用到
	e := []int64{1, 2, 3}                 // slice（在运用切片时有些地方需要注意一下）, 可以通过append添加元素，然后将结果再赋给自己本身，切片底层有个固定大小的数组，当数组容量不够时会生成一个新的更大的数组。
	f := map[string]int64{"a": 3, "b": 4} // map
	fmt.Printf("type:%T:%v\n", p, p)
	fmt.Printf("type:%T:%v\n", a, a)
	fmt.Printf("type:%T:%v\n", b, b)
	fmt.Printf("type:%T:%v\n", c, c)
	fmt.Printf("type:%T:%v\n", d, d)
	fmt.Printf("type:%T:%v\n", e, e)
	// 切片如何切
	fmt.Printf("type:%T:%v\n", e, e[:1])
	fmt.Printf("type:%T:%v\n", e, e[1:])
	fmt.Printf("type:%T:%v\n", e, e[:])
	fmt.Printf("type:%T:%v\n", f, f)
}
