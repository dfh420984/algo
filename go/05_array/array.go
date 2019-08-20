package array

import (
	"errors"
	"fmt"
)

//Array 声明一个Array结构体
type Array struct {
	data   []int
	length uint
}

/*
*数组的插入，查找，越界，删除等操作
 */

//NewArray 初始化数组内存
func NewArray(capacity uint) *Array {
	if capacity == 0 {
		return nil
	}
	return &Array{
		data:   make([]int, capacity, capacity),
		length: 0,
	}
}

//获取数组长度
func (this *Array) Len() uint {
	return this.length
}

//判断索引是否越界
func (this *Array) IsIndexOutOfRange(index uint) bool {
	if index >= uint(cap(this.data)) { //超过数组最大容量，越界了
		return true
	}
	return false
}

//通过索引查找数组
func (this *Array) Find(index uint) (int, error) {
	if this.IsIndexOutOfRange(index) {
		return 0, errors.New("out of range")
	}
	return this.data[index], nil
}

//插入数值到索引指定位置
func (this *Array) Insert(index uint, val int) error {
	if this.Len() == uint(cap(this.data)) { //数组已满
		return errors.New("full array")
	}

	if index != this.length && this.IsIndexOutOfRange(index) {
		return errors.New("out of range")
	}

	for i := this.length; i > index; i-- {
		this.data[i] = this.data[i-1]
	}

	this.data[index] = val
	this.length++
	return nil
}

//在数组尾部插入
func (this *Array) InsertTotail(val int) error {
	return this.Insert(this.Len(), val)
}

//删除数组指定索引
func (this *Array) Delete(index uint) (int, error) {
	if this.IsIndexOutOfRange(index) {
		return 0, errors.New("out of range")
	}

	val := this.data[index]
	//取出数组元素后交换下标元素值
	for i := index; i < this.Len()-1; i++ {
		this.data[i] = this.data[i+1]
	}
	this.length--
	return val, nil
}

func (this *Array) Print() {
	var format string
	for i := uint(0); i < this.Len(); i++ {
		format += fmt.Sprintf("|%+v", this.data[i])
	}
	fmt.Println(format)
}
