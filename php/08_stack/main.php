<?php
/**
 * Created by PhpStorm.
 * User: duanfuhao
 * Date: 2019/8/29
 * Time: 16:09
 */
require_once './StackOnLinkedList.php';

$stack = new StackOnLinkedList();
$stack->push(1);
$stack->push(2);
$stack->push(3);
$stack->push(4);
$stack->printSelf();
echo $stack->pop().PHP_EOL;
echo $stack->pop().PHP_EOL;
echo $stack->pop().PHP_EOL;
echo $stack->pop().PHP_EOL;
$stack->printSelf();
