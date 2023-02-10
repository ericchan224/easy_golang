package data_structure

// 链表节点
type LinkNode struct {
    Next  *LinkNode
    Value string
}

// ListNode 列表节点
type ListNode struct {
	pre   *ListNode // 前驱节点
	next  *ListNode // 后驱节点
	value string    // 值
}

// GetValue 获取节点值
func (node *ListNode) GetValue() string {
    return node.value
}

// GetPre 获取节点前驱节点
func (node *ListNode) GetPre() *ListNode {
    return node.pre
}

// GetNext 获取节点后驱节点
func (node *ListNode) GetNext() *ListNode {
    return node.next
}

// HashNext 是否存在后驱节点
func (node *ListNode) HashNext() bool {
    return node.pre != nil
}

// HashPre 是否存在前驱节点
func (node *ListNode) HashPre() bool {
    return node.next != nil
}

// IsNil 是否为空节点
func (node *ListNode) IsNil() bool {
    return node == nil
}