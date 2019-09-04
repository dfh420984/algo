<?php
/**
 * Created by PhpStorm.
 * User: duanfuhao
 * Date: 2019/9/2
 * Time: 10:25
 */

/**
 * 方式一
 * 快速排序
 * @param $arr
 * @author duanfuhao@smzdm.com
 * @date 2019/9/2
 */
function quicksort(&$arr, $l, $r) {
    if ($l > $r) {
        return false;
    }
    //设置前后游标
    $i = $l;
    $j = $r;
    //假设基准点
    $temp = $arr[$l];
    while ($i < $j) {
        //从右往左，依次递减，找出第一个比基准值小的元素
        while ($arr[$j] >= $temp && $i < $j) {
            $j--;
        }
        //从左往右，依次递增，找到比基准大的元素
        while ($arr[$i] <= $temp && $i < $j) {
            $i++;
        }
        //如果满足条件则交换
        if ($i<$j) {
            $t = $arr[$j];
            $arr[$j] = $arr[$i];
            $arr[$i] = $t;
        }
    }
    //交换基准点数据(前后游标相遇)
    $arr[$l] = $arr[$i];
    $arr[$i] = $temp;
    //递归调用左半数组
    quicksort($arr, $l, $j-1);
    //递归调用右半数组
    quicksort($arr, $j+1, $r);
}

function quicksort2($arr) {
    $len = count($arr);
    if ($len <= 1) {
        return $arr;
    }
    $privot = $arr[0]; //去第一个值作为基准比较点
    $left = [];
    $right = [];
    for ($i = 1; $i < $len; $i++) {
        if ($arr[$i] > $privot) { //比基准点大，放右边数组
            $right[] = $arr[$i];
        } else {
            $left[] = $arr[$i];
        }
    }
    $left = quicksort2($left);
    $right = quicksort2($right);
    return array_merge($left,[$privot],$right);
}

$arr = [1,32,6,8,7,10];
$new_arr = quicksort2($arr);
var_dump($new_arr);