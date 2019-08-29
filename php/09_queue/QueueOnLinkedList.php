<?php
/**
 * 基于链表实现队列
 * Created by PhpStorm.
 * User: duanfuhao
 * Date: 2019/8/29
 * Time: 18:02
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

class QueueOnLinkedList {
    public $length;
    //头节点
    public $head;
    //尾节点
    public $tail;

    public function __construct()
    {
        $this->head = new SingleLinkListNode();
        $this->tail = $this->head;
        $this->length = 0;
    }

    //入队
    public function enqueue($data) {
        $newNode = new SingleLinkListNode($data);
        $this->tail->next = $newNode;
        $this->tail = $newNode;
        $this->length++;
    }

    //出队
    public function dequeue() {
        if (0 == $this->length) {
            return false;
        }
        $node = $this->head->next;
        $this->head->next = $this->head->next->next;
        $this->length--;
        return $node->data;
    }

    /**
     * 获取队列长度
     *
     * @return int
     */
    public function getLength()
    {
        return $this->length;
    }

    public function printSelf() {
        if (0 == $this->length) {
            echo 'empty queue' . PHP_EOL;
            return;
        }
        echo 'head.next -> ';
        $curNode = $this->head;
        while ($curNode->next) {
            echo $curNode->next->data . ' -> ';
            $curNode = $curNode->next;
        }
        echo 'NULL' . PHP_EOL;
    }
}