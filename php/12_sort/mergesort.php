<?php
/**
 * Created by PhpStorm.
 * User: duanfuhao
 * Date: 2019/9/2
 * Time: 15:24
 */

function mergesort(&$arr)
{
    $n = count($arr);
    return merge_sort_c($arr, 0, $n - 1);
}

/**
 * 归并排序
 * @param $arr
 * @param $l
 * @param $r
 * @author duanfuhao@smzdm.com
 * @date 2019/9/2
 */
function merge_sort_c(&$arr, $l, $r)
{
    //分解终止条件
    if ($l >= $r) {
        return ;
    }
    $m = (int)($l + ($r - $l) / 2);
    //分治递归
    merge_sort_c($arr, $l, $m);
    merge_sort_c($arr, $m + 1, $r);
    return merge($arr, $l, $m + 1,$r);
}

//排序，合并
function merge(&$arr, $l, $m,$r)
{
    $temp = [];
    $i = $l;
    $j = $m;
    $k = 0;
    while ($i <= $m && $j <= $r) {
        if ($arr[$i] <= $arr[$j]) {
            $temp[$k++] = $arr[$i++];
        } else {
            $temp[$k++] = $arr[$j++];
        }
    }
    // 判断哪个子数组中有剩余的数据
    while($m > $i) {
        $temp[$k++] = $arr[$i++];
    }
    while($r > $j) {
        $temp[$k++] = $arr[$j++];
    }
    //拷贝数据至原数组指定位置
    for($z=0;$z<count($temp);$z++) {
        $arr[$l+$z] = $temp[$z];
    }
}

$arr = [1, 32, 6, 8, 7, 10];
mergesort($arr);
var_dump($arr);