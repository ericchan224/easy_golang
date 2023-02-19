package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

/*
*
TODO: 协程
*
*/
func Hu() {
	// 使用睡眠模仿一些耗时
	time.Sleep(2 * time.Second)
	fmt.Println("after 2 second hu!!!")
}

func main() {

	// 将会堵塞
	//Hu()

	// 开启新的协程，不会堵塞
	go Hu()

	go func() {
		log.Println("aaa")
	}()

	fmt.Println("start hu, wait...")

	// 必须死循环，不然主协程退出了，程序就结束了
	for {
		time.Sleep(1 * time.Second)
	}

}

/*
*
TODO: 信道
*
*/
func Hu2(ch chan int) {
	// 使用睡眠模仿一些耗时
	time.Sleep(2 * time.Second)
	fmt.Println("after 2 second hu!!!")

	// 执行语句后，通知主协程已经完成操作
	ch <- 1000
}

func test() {
	// 新建一个没有缓冲的信道
	ch := make(chan int)

	// 将信道传入函数，开启协程
	go Hu2(ch)
	fmt.Println("start hu, wait...")

	// 从空缓冲的信道读取 int，将会堵塞，直到有消息到来
	v := <-ch
	fmt.Println("receive:", v)
}

func Receive(ch chan int) {
	// 先等几秒后再接收消息
	time.Sleep(2 * time.Second)
	for {
		select {
		case v, ok := <-ch:
			// 接收信道里面的消息，接收后缓冲就充足了

			// 信道被关闭了，退出
			if !ok {
				fmt.Println("chan close,receive:", v)
				return
			}
			// 打印
			fmt.Println("receive:", v)
		}
	}
}

func Send(ch chan int) {
	// 发到第11个时，会卡住，因为信道满了
	for i := 0; i < 13; i++ {
		ch <- i
		fmt.Println("send:", i)
	}
	// 打印完毕，关闭信道
	close(ch)
}

func test2() {
	// 新建一个10个缓冲的信道
	ch := make(chan int, 10)

	// 将信道传入函数，开启协程
	go Receive(ch)
	go Send(ch)

	// 必须死循环，不然主协程退出了，程序就结束了
	for {
		time.Sleep(1 * time.Second)
	}
}

func test3() {
	buffedChan := make(chan int, 2)
	buffedChan <- 2
	buffedChan <- 3
	close(buffedChan) // 关闭后才能for打印出，否则死锁

	//close(buffedChan) // 不能重复关闭
	//buffedChan <- 4  // 关闭后就不能再送数据了，但是之前的数据还在
	for i := range buffedChan { // 必须关闭，否则死锁
		fmt.Println(i)
	}
}

/*
*
TODO: 锁实现并发安全
*
*/
type Money struct {
	lock   sync.Mutex // 锁
	amount int64
}

// 加钱
func (m *Money) Add(i int64) {
	// 加锁
	m.lock.Lock()

	// 在该函数结束后执行
	defer m.lock.Unlock()
	m.amount = m.amount + i
}

// 减钱
func (m *Money) Minute(i int64) {
	// 加锁
	m.lock.Lock()

	// 在该函数结束后执行
	defer m.lock.Unlock()

	// 钱足才能减
	if m.amount >= i {
		m.amount = m.amount - i
	}
}
