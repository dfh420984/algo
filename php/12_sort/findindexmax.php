<?php

/**
 * 查找一个数组中第N大的数
 * 利用快速排序查找positon的方法，将大于最后一个元素的全部数据放在左边，小于最后一个元素的全部数据放在右边
 * 如果第一次找到的postion为q（下标），那么以为着q左边有q个元素且左边的数全部大于a[q],
 * 那么如果q+1=N，则a[q]为第N大数
 * 如果q+1>N,则第N大数在q左边，反之在q右边
 * @param $n要寻找的第几大元素
 */

function findMaxN(&$arr, $n)
{
    $len = count($arr);
    if ($len <= 1) {
        return false;
    }
    //寻找分区点
    $pos = partion($arr, 0, $len - 1);
    while ($pos + 1 != $n) {
        if ($pos + 1 > $n) { //在左边寻找
            $pos = partion($arr, 0, $pos - 1);
        } else {
            $pos = partion($arr, $pos + 1, $len - 1);
        }
    }
    $res = $arr[$pos];
    return $res;
}

function partion(&$arr, $start, $end) {
    $i = $start;
    $j = $start;
    for (; $j < $end; $j++) { //交换数据
        if ($arr[$j] > $arr[$end]) { //交换$j,$i位置
            if ($j != $i) {
                $temp = $arr[$j];
                $arr[$j] = $arr[$i];
                $arr[$i] = $temp;
            }
            $i++;
        }
    }
    //交换最后的$i和$end数据
    $temp = $arr[$end];
    $arr[$end] = $arr[$i];
    $arr[$i] = $temp;
    return $i;
}
$arr = [6,3,1,5,7,4,2,9,11,8];
$res = findMaxN($arr, 4);
var_dump($arr);
var_dump($res);