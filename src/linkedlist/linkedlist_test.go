package linkedlist

import "testing"

func TestName(t *testing.T) {
	linkedList := NewLinkedList()
	linkedList.InsertToHead(10)
	linkedList.InsertToTail(30)
	linkedList.InsertToHead(40)
	linkedList.InsertToHead(50)
	linkedList.InsertToTail(100)
	linkedList.Print()

	node, ok := linkedList.FindByIndex(2)
	linkedList.InsertBefore(node, 1000)
	_ = ok
	t.Log(node.v)
	linkedList.DeleteNode(node)
	linkedList.Print()
}
