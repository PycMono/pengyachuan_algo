package array

import "fmt"

// Array 数组结构体
type Array struct {
	// 数据
	data []interface{}

	// 长度
	len int
}

// Find 查找数据
func (a *Array) Find(index int) (interface{}, error) {
	if a.isOutOfRange(index) {
		return nil, fmt.Errorf("out of index range")
	}

	return a.data[index], nil
}

// Insert 插入数据
func (a *Array) Insert(index int, v interface{}) error {
	if a.isOutOfRange(index) {
		return fmt.Errorf("out of index range")
	}

	for j := a.Len(); j > index; j-- {
		a.data[j] = a.data[j-1]
	}

	a.data[index] = v
	a.len++
	return nil
}

// Delete 删除数据
func (a *Array) Delete(index int) error {
	if a.isOutOfRange(index) {
		return fmt.Errorf("out of index range")
	}

	for j := index; j < a.Len()-1; j++ {
		a.data[j] = a.data[j+1]
	}

	a.len--
	return nil
}

// Len 返回数组长度
func (a *Array) Len() int {
	return a.len
}

// isOutOfRange 是否越界
func (a *Array) isOutOfRange(index int) bool {
	return index >= cap(a.data)
}

//打印数列
func (a *Array) Print() {
	var format string
	for i := 0; i < a.Len(); i++ {
		format += fmt.Sprintf("|%+v", a.data[i])
	}
	fmt.Println(format)
}

// NewArray 创建数组
func NewArray(len int) *Array {
	if len <= 0 {
		return nil
	}

	return &Array{
		data: make([]interface{}, len, len),
		len:  0,
	}
}
