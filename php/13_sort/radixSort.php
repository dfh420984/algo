<?php
/**
 * Created by PhpStorm.
 * User: duanfuhao
 * Date: 2019/9/6
 * Time: 15:10
 */

/**
 * 基数排序
 * 先根据个位排序、百位、千位........
 * @param $arr
 * @author duanfuhao@smzdm.com
 * @date 2019/9/6
 */
function radixSort(&$arr,$loop) {
    $divisor = pow(10,$loop);
    //初始化10个桶
    $buckets = (new \SplFixedArray(10))->toArray();
    foreach ($arr as $k=>$v) {
        $index = ($v/$divisor)%10; //计算该数字在哪个桶中
        $buckets[$index][] = $v;
    }
    $k=0;
    //从桶中取出数字
    for ($i = 0; $i < count($buckets); $i++) {
        while (count($buckets[$i]) > 0) {
            $arr[$k++] = array_shift($buckets[$i]);
        }
    }
}

$arr = [1234,4321,12,31,412];
//求出最大数
$max = max($arr);
//算出比较次数,决定循环次数
$len = strlen($max);

for($i=0; $i<$len;$i++) {
    radixSort($arr,$i);
}

var_dump($arr);

