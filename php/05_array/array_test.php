<?php
/**
 * Created by PhpStorm.
 * User: leo
 * Date: 2018/10/5
 * Time: 9:13
 */
require "./array.php";
$myArr1 = new MyArray(10);
for($i=0;$i<9;$i++) {
    $myArr1->insert($i, $i+1);
}
$myArr1->printData();
try {
    $code = $myArr1->insert(6, 999);
    $myArr1->printData();
    $value = $myArr1->delete(6);
    var_dump($value);
    $myArr1->printData();
    $code = $myArr1->insert(11, 999);
    var_dump($code);
    $myArr1->printData();
    $value = $myArr1->delete(0);
    var_dump(0,$value);
    $myArr1->printData();
    $value = $myArr1->find(0);
    echo "find at 0: value:{$value}\n";
} catch (Exception $e) {
    echo $e->getMessage();
}