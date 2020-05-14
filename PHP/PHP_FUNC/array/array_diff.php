<?php
/**
array array_diff ( array $array1 , array $array2 [, array $... ] )
array1
要被对比的数组

array2
和这个数组进行比较

...
更多相比较的数组

返回值

返回一个数组，该数组包括了所有在 array1 中但是不在任何其它参数数组中的值。注意键名保留不变。
**/

$array1 = array("a" => "green", "red", "blue", "red");
$array2 = array("b" => "green", "yellow", "red");
print_r(array_diff($array1, $array2));

/**
Array
(
    [1] => blue
)
**/
