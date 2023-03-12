package main

import (
	"fmt"
	"io"
)

// Sayer 接口
type Sayer interface {
	say()
}

// Mover 接口
type Mover interface {
	move()
}

// 接口嵌套
type animal interface {
	Sayer
	Mover
}

type dog struct {
	name string
}

// 实现Sayer接口
func (d dog) say() {
	fmt.Printf("%s会叫汪汪汪\n", d.name)
}

// 实现Mover接口
func (d dog) move() {
	fmt.Printf("%s会动\n", d.name)
}

func main() {
	var _ io.Writer
	var x Sayer
	var y Mover
	var z animal

	var a = dog{name: "旺财"}
	x = a
	y = a
	z = a
	x.say()
	y.move()
	z.move()
}
