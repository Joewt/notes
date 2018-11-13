<?php
/**
array array_intersect_assoc ( array $array1 , array $array2 [, array $... ] )
array1
要检查的主值。

array2
要比较的数组。

...
要对比的数组变量的列表。

返回值

返回数组，该数组包含了所有在 array1 中也同时出现在所有其它参数数组中的值。
**/

$array1 = array("a" => "green", "b" => "brown", "c" => "blue", "red");
$array2 = array("a" => "green", "b" => "yellow", "blue", "red");
$result_array = array_intersect_assoc($array1, $array2);
print_r($result_array);

/**
Array
(
    [a] => green
)
**/
