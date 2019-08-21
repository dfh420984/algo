<?php
/**
 * Created by PhpStorm.
 * User: duanfuhao
 * Date: 2019/8/21
 * Time: 17:58
 */

require_once "./SingleLinkList.php";

//判断链表中保存的字符串是否是回文
function isLinkPalindrome($list) {
    if ($list->getLength() <= 1) {
        return false;
    }
    //定义快慢两个指针
    $pre = null;
    $slow = $list->head->next;
    $fast = $list->head->next;
    $remainNode = null;

    // 找单链表中点 以及 反转前半部分链表
    while ($fast != null && $fast->next != null) {
        $fast = $fast->next->next;
        // 单链表反转关键代码 三个指针
        $remainNode = $slow->next;
        $slow->next = $pre;
        $pre = $slow;
        $slow = $remainNode;
    }
    // 链表长度为偶数的情况
    if ($fast != null) {
        $slow = $slow->next;
    }
    // 开始逐个比较
    while ($slow != null) {
        if ($slow->data != $pre->data) {
            return false;
        }
        $slow = $slow->next;
        $pre = $pre->next;
    }
    return true;
}

$list = new SingleLinkList();
//$list->insert('a');
//$list->insert('b');
//$list->insert('c');
//print_r($list->getPreNode($list->getNodeByIndex(2)));
//$list->deleteNode('b');
//$list->buildHasCircleList();
//var_dump($list->isCircleLink($list->head));
//print_r($list->getLength());
//$list->printNode();
//$list->printCircleNode();
$list->insert('a');
$list->insert('b');
$list->insert('c');
$list->insert('c');
$list->insert('b');
$list->insert('a');
//$list->printNode();
var_dump(isLinkPalindrome($list));
