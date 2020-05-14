<?php
/**
array array_diff_uassoc ( array $array1 , array $array2 [, array $... ], callable $key_compare_func )
array1
待比较的数组

array2
和这个数组进行比较

...
更多比较的数组

key_compare_func
在第一个参数小于，等于或大于第二个参数时，该比较函数必须相应地返回一个小于，等于或大于 0 的整数。

int callback ( mixed $a, mixed $b )
**/


function key_compare_func($a, $b)
{
    if ($a === $b) {
        return 0;
    }
    return ($a > $b)? 1:-1;
}

$array1 = array("a" => "green", "b" => "brown", "c" => "blue", "red");
$array2 = array("a" => "green", "yellow", "red");
$result = array_diff_uassoc($array1, $array2, "key_compare_func");
print_r($result);


/**
Array
(
    [b] => brown
    [c] => blue
    [0] => red
)
**/
