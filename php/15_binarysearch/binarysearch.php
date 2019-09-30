<?php
/**
 * https://github.com/wangzheng0822/algo/blob/master/php/15_binary/binary.php
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
    $mid = (int)($low + ($high - $low) / 2);
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

function bsearch($arr,$val) {
    return bsearch_rec($arr,$val,0,count($arr) - 1);
}

function bsearch_rec($arr,$val,$l,$h) {
    if ($l > $h) return -1;
    $mid = int($l + ($h - $l) / 2);
    if ($arr[$mid] < $val) {
        bsearch_rec($arr,$val,$mid + 1,$h);
    } elseif ($arr[$mid] > $val) {
        bsearch_rec($arr,$val,$l,$mid - 1);
    } else {
        return $mid;
    }
}

/**
 * 二分法求平方根,精确到小数点后6位
 * @param $val
 * @param $double
 * @author duanfuhao@smzdm.com
 * @date 2019/9/30
 */
function square($val, $double) {
    if ($val < 0) { //必须是正数
        return false;
    }
    $low = 0;
    $up = $val;
    if ($val > 0 && $val < 1) {
        $low = $val;
        $up = 1;
    }
    $mid = $low + ($up - $low) / 2;
    while ($up - $low > $double) {
        if ($mid * $mid > $val) {
            $up = $mid;
        } elseif ($mid * $mid < $val) {
            $low = $mid;
        } else {
            return $mid;
        }
        $mid = $low + ($up - $low) / 2;
    }
    return $mid;
}

var_dump(square(9.9, 0.000001));
