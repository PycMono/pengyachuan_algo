package array

import "testing"

func TestName(t *testing.T) {
	array := NewArray(10)
	array.Insert(0, 100)
	array.Insert(1, 100)
	array.Insert(3, 20)
	array.Insert(9, 20)
	array.Print()
	array.Delete(2)
	array.Delete(3)
	array.Print()
}
