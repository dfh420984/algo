<?php
/**
 * Created by PhpStorm.
 * User: duanfuhao
 * Date: 2019/9/6
 * Time: 10:31
 */

/**
 * 桶排序
 * @param $arr
 * @param $bucketNum 桶的数量
 * @param $bucketMaxSize 每个桶的大小
 * @author duanfuhao@smzdm.com
 * @date 2019/9/6
 */
function bucketSort($arr, $bucketNum, $bucketMaxSize)
{
    $len = count($arr);
    if ($len <= 1) {
        return $arr;
    }
    //计算每个桶的范围
    $min = min($arr); //取出数组最小值
    $max = max($arr); //取出最大值
    $bucketRange = ($max - $min + 1) / $bucketNum; //计算每个桶得的数据范围
    $buckets = [];
    //将原数据分配到每个桶中
    foreach ($arr as $v) {
        $index = floor(($v - $min) / $bucketRange);
        $buckets[$index][] = $v;
    }
    //开始遍历每个桶中数据进行排序
    $result = [];
    for ($i = 0; $i < $bucketNum; $i++) {
        $bucket = $buckets[$i];
        $length = count($bucket);
        if ($length == 0) {
            continue;
        }
        if ($length > $bucketMaxSize) { //桶内数据超过设定，桶内元素个数，继续分桶
            $bucket = bucketSort($bucket, $bucketNum, $bucketMaxSize);
        }
        //开始排序桶内数据
        quicksort($bucket,0,$length-1);
        //将数据合并到新数组中
        $result = array_merge($result,$bucket);
    }
    return $result;
}

function quicksort(&$arr,$l,$r) {
    if ($l >= $r) {
        return;
    }
    //寻找分区点
    $q = partion($arr,$l,$r);
    //开始分区
    quicksort($arr,$l,$q-1);
    quicksort($arr,$q+1,$r);
}

function partion(&$arr,$l,$r) {
    $i = $l;
    $j = $l;
    //将最右边的数设为比较基准点
    $privot = $arr[$r];
    //小于基准点放左边，大于基准点放右边
    for (; $j < $r; $j++) {
        if ($arr[$j] < $privot) {
            if ($i != $j) {
                $temp = $arr[$j];
                $arr[$j] = $arr[$i];
                $arr[$i] = $temp;
            }
            $i++;
        }
    }
    //循环完毕后，找到最终基准点，并和比较基准点交换数据
    $temp = $arr[$i];
    $arr[$i] = $arr[$r];
    $arr[$r] = $temp;
    return $i;
}

$arr = [11, 23, 45, 67, 88, 99, 22, 34, 56, 78, 90, 12, 34, 5, 6, 91, 92, 93, 93, 94, 95, 94, 95, 96, 97, 98, 99, 100];
$bucketNum = 5;
$bucketMaxSize = 5;
$res = bucketSort($arr, $bucketNum, $bucketMaxSize);
var_dump($res);