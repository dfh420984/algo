<?php
/**
 * Created by PhpStorm.
 * User: duanfuhao
 * Date: 2019/8/21
 * Time: 14:56
 */

/**
 * 单向链表节点
 */

class SingleLinkListNode {
    public $data; //节点数据部分
    public $next;  //下一个节点指针

    public function __construct($data = null)
    {
        $this->data = $data;
        $this->next = null;
    }
}