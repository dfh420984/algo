<?php
/**
 * Created by PhpStorm.
 * User: duanfuhao
 * Date: 2019/8/21
 * Time: 11:09
 */

class MyArray {
    public $data;
    public $length;
    public $capacity;

    public function __construct($capacity)
    {
        $this->data = array();
        $this->capacity = $capacity; //数组最大容量
        $this->length = 0;  //数组现有数据长度
    }

    //判断数组是否已满
    public function isFull() {
        if ($this->length >= $this->capacity) {
            return true;
        }
        return false;
    }

    //判断下标是否越界
    public function isIndexOutOfRange($index) {
        if ($index < 0) {
            return true;
        }
        if ($index >= $this->capacity) { //下标是否越界
            return true;
        }
        return false;
    }

    //将数据插入索引指定位置
    public function insert($index,$val) {
        if ($this->isFull()) {
            throw new Exception("数组已满");
        }

        if ($this->isIndexOutOfRange($index)) {
            throw new Exception("下标越界 index:".$index);
        }
        //插入指定位置索引数据后，在此之后得索引如果有数据，需要交换
        for ($i = $this->length-1; $index <= $i; $i--) {
            $this->data[$i+1] = $this->data[$i];
        }
        $this->data[$index] = $val;
        $this->length++;
        return true;
    }

    //更加下标查找值
    public function find($index) {
        if ($this->isIndexOutOfRange($index)) {
            throw new Exception("下标越界 index:".$index);
        }
        return $this->data[$index];
    }

    //删除指定索引
    public function delete($index) {
        if ($this->isIndexOutOfRange($index)) {
            throw new Exception("下标越界 index:".$index);
        }
        //删除指定索引后，后面元素需要往前移动，保持数组连续性
        $val = $this->data[$index];
        for ($i = $index; $i < $this->length-1; $i++) {
            $this->data[$i] = $this->data[$i+1];
        }
        $this->length--;
        return $val;
    }

    //打印数组
    public function printData() {
        $format = "";
        for ($i = 0; $i < $this->length; $i++) {
            $format .= "|" . $this->data[$i];
        }
        print($format . "\n");
    }
}