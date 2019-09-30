<?php
/**
 * Created by PhpStorm.
 * User: duanfuhao
 * Date: 2019/9/26
 * Time: 14:59
 */
function binarysearch($arr,$val) {
    $len = count($arr);
    if ($len <= 1) return $arr;
    $low = 0;
    $high = $len - 1;
    $mid = int($low + ($high - $low) / 2);
    while ($low <= $high) {
        if ($arr[$mid] < $val) {
            $low = $mid + 1;
        } elseif ($arr[$mid] > $val) {
            $high = $mid - 1;
        } else {
            return $mid;
        }
    }
    return -1;
}
