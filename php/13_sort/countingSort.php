<?php
/**
 * Created by PhpStorm.
 * User: duanfuhao
 * Date: 2019/9/6
 * Time: 14:37
 */

/**
 * * 计数排序
 * 五分制 [0,5]
 * $arr = [0,1,5,3,2,4,1,2,4,2,1,4,4];
 * 13个人
 * @param $arr
 * @author duanfuhao@smzdm.com
 * @date 2019/9/6
 */
function countingSort($arr)
{
    $len = count($arr);
    if ($len <= 1) {
        return $arr;
    }
    $bucketArr = [];
    $max = max($arr);
    $min = min($arr);
    //初始化桶
    for ($i = $min; $i <= $max; $i++) {
        $bucketArr[$i] = 0;
    }
    //计算每个元素，计算每个元素桶内元素个数
    foreach ($arr as $k => $v) {
        $bucketArr[$v]++;
    }
    //顺序求和 $arr[$i] = $arr[$i-1] + $arr[$i]
    for ($i = 1; $i <= $max; $i++) {
        $bucketArr[$i] = $bucketArr[$i] + $bucketArr[$i-1];
    }
    //开始排序
    $temp = [];
    foreach ($arr as $k=>$v) {
        $bucketArr[$v]--;
        $temp[$bucketArr[$v]] = $v;
    }
    //copy
    for ($i=0;$i<$len;$i++) {
        $arr[$i] = $temp[$i];
    }
    return $arr;
}

$arr = [0,1,5,3,2,4,1,2,4,2,1,4,4];
var_dump(countingSort($arr));