<?php
/**
 * Created by PhpStorm.
 * User: duanfuhao
 * Date: 2019/8/21
 * Time: 17:58
 */

require_once "./SingleLinkList.php";

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
$list->insert('c');
$list->insert('b');
$list->insert('a');
//$list->printNode();
var_dump($list->isLinkPalindrome($list));
//var_dump($list->reverseLinkedList($list->head->next));
