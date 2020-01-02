package linkedlist

import "fmt"

// LinkedList 链表信息
type LinkedList struct {
	// 头节点信息
	head *ListNode

	// 链表长度
	len int
}

// InsertToHead 头部插入
func (l *LinkedList) InsertToHead(v interface{}) bool {
	return l.InsertAfter(l.head, v)
}

// InsertToTail 尾部插入
func (l *LinkedList) InsertToTail(v interface{}) bool {
	// 尾部插入
	cur := l.head
	for cur.next != nil {
		cur = cur.next
	}

	return l.InsertAfter(cur, v)
}

// InsertBefore 在某一个节点之前插入数据
func (l *LinkedList) InsertBefore(node *ListNode, v interface{}) bool {
	if node == nil || l.head == node {
		return false
	}

	// 遍历所有节点，找到和node对应的节点
	cur := l.head.next
	pre := l.head
	for cur != nil {
		if cur == node {
			break
		}

		pre = cur
		cur = cur.next
	}

	if nil == cur {
		return false
	}

	newNode := NewListNode(v)
	pre.next = newNode
	newNode.next = cur
	l.len++
	return true
}

// InsertAfter 某一个节点之后插入数据
func (l *LinkedList) InsertAfter(node *ListNode, v interface{}) bool {
	if node == nil {
		return false
	}

	newNode := NewListNode(v)
	oldNext := node.next
	newNode.next = oldNext
	node.next = newNode
	l.len++
	return true
}

// FindByIndex 查找数据
func (l *LinkedList) FindByIndex(index int) (*ListNode, bool) {
	if l.head == nil || index > l.len {
		return nil, false
	}

	cur := l.head.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}

	return cur, true
}

// DeleteNode 删除节点
func (l *LinkedList) DeleteNode(node *ListNode) bool {
	if l.head == nil {
		return false
	}

	cur := l.head.next
	pre := l.head
	for cur != nil {
		if cur == node {
			break
		}

		pre = cur
		cur = cur.next
	}

	if cur != node {
		return false
	}

	pre.next = node.next
	node = nil
	l.len--

	return true
}

// Print 打印链表
func (l *LinkedList) Print() {
	if l.head == nil {
		return
	}

	cur := l.head.next
	format := ""
	for nil != cur {
		format += fmt.Sprintf("%+v", cur.GetValue())
		cur = cur.next
		if nil != cur {
			format += "->"
		}
	}
	fmt.Println(format)
}

// NewLinkedList 创建新的LinkedList对象
func NewLinkedList() *LinkedList {
	return &LinkedList{
		head: NewListNode(0), // 头结点默认为0，不然后面不好处理
		len:  0,
	}
}
