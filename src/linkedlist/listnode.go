package linkedlist

// ListNode 链表节点
type ListNode struct {
	// 值
	v interface{}

	// 下一个节点信息
	next *ListNode
}

// GetValue 获取当前节点的值
func (l* ListNode)GetValue()interface{}  {
	return l.v
}

// GetNextNode 获取下一个节点的指针
func (l* ListNode)GetNextNode()*ListNode {
	return l.next
}

// NewListNode 获取新的值
func NewListNode(v interface{})*ListNode  {
	return &ListNode{
		v:    v,
		next: nil,
	}
}