package data_structure

import (
	"sync"
)

// DoubleList 双端列表，双端队列
type DoubleList struct {
	head *ListNode  // 指向链表头部
	tail *ListNode  // 指向链表尾部
	len  int        // 列表长度
	lock sync.Mutex // 为了进行并发安全pop弹出操作
}

// Len 返回列表长度
func (list *DoubleList) Len() int {
	return list.len
}

// AddNodeFromHead 从头部开始，添加节点到第N+1个元素之前，N=0表示添加到第一个元素之前，表示新节点成为新的头部，N=1表示添加到第二个元素之前，以此类推
func (list *DoubleList) AddNodeFromHead(n int, v string) {
	// 加并发锁
	list.lock.Lock()
	defer list.lock.Unlock()

	// 如果索引超过或等于列表长度，一定找不到，直接panic
	if n != 0 && n >= list.len {
		panic("index out")
	}

	// 先找出头部
	node := list.head

	// 往后遍历拿到第 N+1 个位置的元素
	for i := 1; i <= n; i++ {
		node = node.next
	}

	// 新节点
	newNode := new(ListNode)
	newNode.value = v

	// 如果定位到的节点为空，表示列表为空，将新节点设置为新头部和新尾部
	if node.IsNil() {
		list.head = newNode
		list.tail = newNode
	} else {
		// 定位到的节点，它的前驱
		pre := node.pre

		// 如果定位到的节点前驱为nil，那么定位到的节点为链表头部，需要换头部
		if pre.IsNil() {
			// 将新节点链接在老头部之前
			newNode.next = node
			node.pre = newNode
			// 新节点成为头部
			list.head = newNode
		} else {
			// 将新节点插入到定位到的节点之前
			// 定位到的节点的前驱节点 pre 现在链接到新节点前
			pre.next = newNode
			newNode.pre = pre

			// 定位到的节点链接到新节点之后
			newNode.next = node
			node.pre = newNode
		}

	}

	// 列表长度+1
	list.len = list.len + 1
}

// AddNodeFromTail 从尾部开始，添加节点到第N+1个元素之后，N=0表示添加到第一个元素之后，表示新节点成为新的尾部，N=1表示添加到第二个元素之后，以此类推
func (list *DoubleList) AddNodeFromTail(n int, v string) {
	// 加并发锁
	list.lock.Lock()
	defer list.lock.Unlock()

	// 如果索引超过或等于列表长度，一定找不到，直接panic
	if n != 0 && n >= list.len {
		panic("index out")
	}

	// 先找出尾部
	node := list.tail

	// 往前遍历拿到第 N+1 个位置的元素
	for i := 1; i <= n; i++ {
		node = node.pre
	}

	// 新节点
	newNode := new(ListNode)
	newNode.value = v

	// 如果定位到的节点为空，表示列表为空，将新节点设置为新头部和新尾部
	if node.IsNil() {
		list.head = newNode
		list.tail = newNode
	} else {
		// 定位到的节点，它的后驱
		next := node.next

		// 如果定位到的节点后驱为nil，那么定位到的节点为链表尾部，需要换尾部
		if next.IsNil() {
			// 将新节点链接在老尾部之后
			node.next = newNode
			newNode.pre = node

			// 新节点成为尾部
			list.tail = newNode
		} else {
			// 将新节点插入到定位到的节点之后
			// 新节点链接到定位到的节点之后
			newNode.pre = node
			node.next = newNode

			// 定位到的节点的后驱节点链接在新节点之后
			newNode.next = next
			next.pre = newNode

		}

	}

	// 列表长度+1
	list.len = list.len + 1
}

// First 返回列表链表头结点
func (list *DoubleList) First() *ListNode {
	return list.head
}

// Last 返回列表链表尾结点
func (list *DoubleList) Last() *ListNode {
	return list.tail
}

// IndexFromHead 从头部开始往后找，获取第N+1个位置的节点，索引从0开始。
func (list *DoubleList) IndexFromHead(n int) *ListNode {
	// 索引超过或等于列表长度，一定找不到，返回空指针
	if n >= list.len {
		return nil
	}

	// 获取头部节点
	node := list.head

	// 往后遍历拿到第 N+1 个位置的元素
	for i := 1; i <= n; i++ {
		node = node.next
	}

	return node
}

// IndexFromTail 从尾部开始往前找，获取第N+1个位置的节点，索引从0开始。
func (list *DoubleList) IndexFromTail(n int) *ListNode {
	// 索引超过或等于列表长度，一定找不到，返回空指针
	if n >= list.len {
		return nil
	}

	// 获取尾部节点
	node := list.tail

	// 往前遍历拿到第 N+1 个位置的元素
	for i := 1; i <= n; i++ {
		node = node.pre
	}

	return node
}

// PopFromHead 从头部开始往后找，获取第N+1个位置的节点，并移除返回
func (list *DoubleList) PopFromHead(n int) *ListNode {
	// 加并发锁
	list.lock.Lock()
	defer list.lock.Unlock()

	// 索引超过或等于列表长度，一定找不到，返回空指针
	if n >= list.len {
		return nil
	}

	// 获取头部
	node := list.head

	// 往后遍历拿到第 N+1 个位置的元素
	for i := 1; i <= n; i++ {
		node = node.next
	}

	// 移除的节点的前驱和后驱
	pre := node.pre
	next := node.next

	// 如果前驱和后驱都为nil，那么移除的节点为链表唯一节点
	if pre.IsNil() && next.IsNil() {
		list.head = nil
		list.tail = nil
	} else if pre.IsNil() {
		// 表示移除的是头部节点，那么下一个节点成为头节点
		list.head = next
		next.pre = nil
	} else if next.IsNil() {
		// 表示移除的是尾部节点，那么上一个节点成为尾节点
		list.tail = pre
		pre.next = nil
	} else {
		// 移除的是中间节点
		pre.next = next
		next.pre = pre
	}

	// 节点减一
	list.len = list.len - 1
	return node
}

// PopFromTail 从尾部开始往前找，获取第N+1个位置的节点，并移除返回
func (list *DoubleList) PopFromTail(n int) *ListNode {
	// 加并发锁
	list.lock.Lock()
	defer list.lock.Unlock()

	// 索引超过或等于列表长度，一定找不到，返回空指针
	if n >= list.len {
		return nil
	}

	// 获取尾部
	node := list.tail

	// 往前遍历拿到第 N+1 个位置的元素
	for i := 1; i <= n; i++ {
		node = node.pre
	}

	// 移除的节点的前驱和后驱
	pre := node.pre
	next := node.next

	// 如果前驱和后驱都为nil，那么移除的节点为链表唯一节点
	if pre.IsNil() && next.IsNil() {
		list.head = nil
		list.tail = nil
	} else if pre.IsNil() {
		// 表示移除的是头部节点，那么下一个节点成为头节点
		list.head = next
		next.pre = nil
	} else if next.IsNil() {
		// 表示移除的是尾部节点，那么上一个节点成为尾节点
		list.tail = pre
		pre.next = nil
	} else {
		// 移除的是中间节点
		pre.next = next
		next.pre = pre
	}

	// 节点减一
	list.len = list.len - 1
	return node
}
