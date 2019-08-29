<?php
/**
 * Created by PhpStorm.
 * User: duanfuhao
 * Date: 2019/8/29
 * Time: 15:30
 */

//节点
class SingleLinkListNode {
    public $data; //节点数据部分
    public $next;  //下一个节点指针

    public function __construct($data = null)
    {
        $this->data = $data;
        $this->next = null;
    }
}

//基于链表实现栈
class StackOnLinkedList {
    public $topNode;
    public $length;

    public function __construct()
    {
        $this->topNode = new SingleLinkListNode();
        $this->length = 0;
    }

    //入栈
    public function push($data) {
        $node = new SingleLinkListNode($data);
        $this->pushNode($node);
    }

    public function pushNode($node) {
        $cur = $this->topNode;
        while ($cur->next != null) {
            $cur = $cur->next;
        }
        $cur->next = $node;
        $this->length++;
    }

    //出栈
    public function pop() {
        if ($this->length == 0) {
            return null;
        }
        $data= $this->topNode->next->data;
        $this->topNode->next = $this->topNode->next->next;
        $this->length--;
        return $data;
    }

    public function getLen() {
        return $this->length;
    }

    //打印
    public function printSelf() {
        if ($this->isEmpty()) {
            return null;
        } else {
            $curNode = $this->topNode;
            while ($curNode->next != null) {
                echo $curNode->next->data . ' -> ';
                $curNode = $curNode->next;
            }
            echo 'NULL' . PHP_EOL;
        }
    }

    /**
     * 判断栈是否为空
     *
     * @return bool
     */
    public function isEmpty()
    {
        return $this->length > 0 ? false : true;
    }
}