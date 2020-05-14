<?php
/**
array array_keys ( array $array [, mixed $search_value = null [, bool $strict = false ]] )
input
一个数组，包含了要返回的键。

search_value
如果指定了这个参数，只有包含这些值的键才会返回。

strict
判断在搜索的时候是否该使用严格的比较（===）。
**/

$arr = array(0 => 100, "color" => "red");
print_r(array_keys($arr));

$arr = array('a'=>11, 'b'=>22, 'c'=>33);
print_r(array_keys($arr));


$array = array("blue", "red", "green", "blue", "blue");
print_r(array_keys($array, "blue"));

/**
Array
(
    [0] => 0
    [1] => color
)
Array
(
    [0] => a
    [1] => b
    [2] => c
)
Array
(
    [0] => 0
    [1] => 3
    [2] => 4
)
**/
