<?php
/**
array array_diff_assoc ( array $array1 , array $array2 [, array $... ] )

array1
从这个数组进行比较

array2
被比较的数组

...
更多被比较的数组
**/
$array1 = array("a" => "green", "b" => "brown", "c" => "blue", "red");
$array2 = array("a" => "green", "yellow", "red");
$result = array_diff_assoc($array1, $array2);
print_r($result);


/**
Array
(
    [b] => brown
    [c] => blue
    [0] => red
)
**/
