<?php
/**
 * Created by PhpStorm.
 * User: duanfuhao
 * Date: 2019/8/29
 * Time: 16:09
 */
require_once './QueueOnLinkedList.php';

$stack = new QueueOnLinkedList();
$stack->enqueue(1);
$stack->enqueue(2);
$stack->enqueue(3);
$stack->enqueue(4);
$stack->printSelf();
echo $stack->dequeue().PHP_EOL;
echo $stack->dequeue().PHP_EOL;
echo $stack->dequeue().PHP_EOL;
echo $stack->dequeue().PHP_EOL;
$stack->printSelf();
