package main

import "fmt"

func main() {
	a := 1
	//TODO:, 判断语句
	if a > 0 {
		fmt.Println("a>0")
	} else if a== 0 {
		fmt.Println("a=0")
	} else {
		fmt.Println("a<0")
	}

	//TODO: 如果有过多的else if 建议使用switch case
	num := 4

    switch num {
    case 3:
        fmt.Println(3)
    case 4:
        fmt.Println(4)
    case 5:
        fmt.Println(5)
    default:
        fmt.Println("not found")
	}

	//TODO: 循环语句, golang只有for 循环
	// for 起始状态; 进入循环需满足的条件; 每次循环后执行的指令 {}
    for i := 9; i <= 10; i++ {
        fmt.Printf("i=%d\n", i)
    }
	// 死循环语句
    a = 0
    for {
        if a >= 10 {
            fmt.Println("out")
            // 退出循环
            break
        }

        a = a + 1
        if a > 5 {
            continue
        } 
        
        fmt.Println(a)
    }

	e := []int64{1, 2, 3}                 // slice
    f := map[string]int64{"a": 3, "b": 4} // map

    // 循环切片
    for k, v := range e {
        fmt.Println(k, v)
    }

    // 循环map
    for k, v := range f {
        fmt.Println(k, v)
    }
}
