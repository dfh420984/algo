<?php
/**
 * Created by PhpStorm.
 * User: duanfuhao
 * Date: 2019/8/30
 * Time: 14:35
 */

/**
 * 电影院递归自己位置
 * 递推 f(1) = 1 f(n) = f(n) - 1
 */
function film($n) {
    if ($n == 1) {
        return 1;
    }
    return film($n-1) + 1;
}

var_dump(film(5));

/**
 * 假如这里有 n 个台阶，每次你可以跨 1 个台阶或者 2 个...
* 递推 f(1) = 1 f(2) = 2 f(n) = f(n-1) + f(n-2)
 */
function taijie($n) {
    if ($n == 1) {
        return 1;
    }
    if ($n == 2) {
        return 2;
    }
    $map = [];
    if (isset($map[$n])) {
        return $map[$n];
    }
    $res =  taijie($n-1) + taijie($n-2);
    $map[$n] = $res;
    return $res;
}

var_dump(taijie(4));

/**
 * 递归阶乘
 * f(1) = 1 f(n) = n*f(n-1)
 */

function jiecheng($n) {
    if ($n == 1) {
        return $n;
    }
    $res = $n * jiecheng($n-1);
    return $res;
}

var_dump(jiecheng(4));