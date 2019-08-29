<?php
/**
 * 循环队列
 * Created by PhpStorm.
 * User: duanfuhao
 * Date: 2019/8/29
 * Time: 18:33
 */
class LoopQueue{
    public $max; //数组元素大小
    public $head;
    public $tail;
    public $data;

    public function __construct($max)
    {
        $this->data = [];
        $this->max = $max;
        $this->head = 0;
        $this->tail = 0;
    }

    //入队
    public function enqueue($data) {
        if (($this->tail + 1) % $this->max == $this->head) { //队满
            echo "tail:".$this->tail.PHP_EOL;
            echo "head:".$this->head.PHP_EOL;
            return false;
        }
        $this->data[$this->tail] = $data;
        $this->tail = ($this->tail + 1) % $this->max;
        return true;
    }

    //出队
    public function dequeue() {
        if ($this->head == $this->tail) {
            return false;
        }
        $res = $this->data[$this->head];
        unset($this->data[$this->head]);
        $this->head = ($this->head + 1) % $this->max;
        return $res;
    }

    public function getLength()
    {
        return ($this->tail - $this->head + $this->max) % $this->max;
    }
}

$queue = new LoopQueue(4);
// var_dump($queue);
$queue->enQueue(1);
$queue->enQueue(2);
$queue->enQueue(3);
$queue->enQueue(4);
var_dump($queue->getLength());
echo $queue->deQueue().PHP_EOL;
echo $queue->deQueue().PHP_EOL;
echo $queue->deQueue().PHP_EOL;
echo $queue->deQueue().PHP_EOL;