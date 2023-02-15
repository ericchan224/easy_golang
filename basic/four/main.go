package main

import (
	"fmt"
	"reflect"
)

/**
接口结构interface, 同样可以用断言和反射判断
**/
// A 定义一个接口，有一个方法
type A interface {
    Println()
}

// B 定义一个接口，有两个方法
type B interface {
    Println()
    Printf() int
}

// A1Instance 定义一个结构体
type A1Instance struct {
    Data string
}

// Println 结构体实现了Println()方法，现在它是一个 A 接口
func (a1 *A1Instance) Println() {
    fmt.Println("a1:", a1.Data)
}

// A2Instance 定义一个结构体
type A2Instance struct {
    Data string
}

// Println 结构体实现了Println()方法，现在它是一个 A 接口
func (a2 *A2Instance) Println() {
    fmt.Println("a2:", a2.Data)
}

// Printf 结构体实现了Printf()方法，现在它是一个 B 接口，它既是 A 又是 B 接口
func (a2 *A2Instance) Printf() int {
    fmt.Println("a2:", a2.Data)
    return 0
}

func main() {
    // 定义一个A接口类型的变量
    var a A

    // 将具体的结构体赋予该变量
    a = &A1Instance{Data: "i love you"}

    // 调用接口的方法
    a.Println()

    // 断言类型
    if v, ok := a.(*A1Instance); ok {
        fmt.Println(v)
    } else {
        fmt.Println("not a A1")
    }

    fmt.Println(reflect.TypeOf(a).String())

    // 将具体的结构体赋予该变量
    a = &A2Instance{Data: "i love you"}

    // 调用接口的方法
    a.Println()

    // 断言类型
    if v, ok := a.(*A1Instance); ok {
        fmt.Println(v)
    } else {
        fmt.Println("not a A1")
    }

    fmt.Println(reflect.TypeOf(a).String())

    // 定义一个B接口类型的变量
    var b B

    //b = &A1Instance{Data: "i love you"} // 不是 B 类型
    b = &A2Instance{Data: "i love you"}

    fmt.Println(b.Printf())
}

/**
数据类型interface{}, 现在可用any代替
**/
func test() {
	// 声明一个未知类型的 a，表明不知道是什么类型
	var a interface{}
	a = 2
	fmt.Printf("%T,%v\n", a, a)

	// 传入函数
	print(a)
	print(3)
	print("i love you")

	// 使用断言，判断是否是 int 数据类型
	v, ok := a.(int)
	if ok {
		fmt.Printf("a is int type,value is %d\n", v)
	}

	// 使用断言，判断变量类型
	switch a.(type) {
	case int:
		fmt.Println("a is type int")
	case string:
		fmt.Println("a is type string")
	default:
		fmt.Println("a not type found type")
	}

	// 使用反射找出变量类型
	t := reflect.TypeOf(a)
	fmt.Printf("a is type: %s", t.Name())
}
