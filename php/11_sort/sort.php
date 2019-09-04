<?php
/**
 * Created by PhpStorm.
 * User: duanfuhao
 * Date: 2019/8/30
 * Time: 16:59
 */

$arr = [5,8,5,2,9,6,11];

/**
 * 冒泡排序
 * @param $arr
 * @return mixed
 * @author duanfuhao@smzdm.com
 * @date 2019/8/30
 */
function bubbleSort($arr) {
    $len = count($arr);
    for($i = 0; $i < $len; $i++) {
        // 提前退出冒泡循环的标志位
        $flag = false;
        for ($j = 0; $j < $len-$i-1; $j++) {
            if ($arr[$j] > $arr[$j+1]) {
                $tmp = $arr[$j+1];
                $arr[$j+1] = $arr[$j];
                $arr[$j] = $tmp;
                $flag = true;
            }
        }
        //没有数据交换，退出
        if (!$flag) break;
    }
    return $arr;
}
var_dump(bubbleSort($arr));

/**
 * 插入排序分成已排序和未排序区间
 * @param $arr
 * @return mixed
 * @author duanfuhao@smzdm.com
 * @date 2019/8/30
 */
function insertSort($arr) {
    $len = count($arr);
    for ($i = 1; $i < $len; $i++) { //待排序数据
        $val = $arr[$i]; //待插入数据
        $j = $i - 1;
        for (; $j >= 0; $j--) { //从后往前去和已排序数据比较
            if ($arr[$j] > $val) {
                $arr[$j + 1] = $arr[$j];
            } else {
                break;
            }
        }
        $arr[$j+1] = $val;
    }
    return $arr;
}
var_dump(insertSort($arr));

/**
 * 选择排序
 * @param $arr
 * @return mixed
 * @author duanfuhao@smzdm.com
 * @date 2019/8/30
 */
function selectSort($arr) {
    $len = count($arr);
    for ($i = 0; $i < $len; $i++) {
        $min = $i; //假设最小值下标
        for ($j = $i + 1; $j < $len; $j++) {
            if ($arr[$min] > $arr[$j]) {
                $min = $j;
            }
        }
        //找到最小值小标后，交换值
        if ($min != $i) {
            $tmp = $arr[$min];
            $arr[$min] = $arr[$i];
            $arr[$i] = $tmp;
        }
    }
    return $arr;
}
var_dump(selectSort($arr));