<?php
/**
array array_filter ( array $array [, callable $callback [, int $flag = 0 ]] )
array
要循环的数组

callback
使用的回调函数

如果没有提供 callback 函数， 将删除 array 中所有等值为 FALSE 的条目。更多信息见转换为布尔值。

flag
决定callback接收的参数形式:

ARRAY_FILTER_USE_KEY - callback接受键名作为的唯一参数
ARRAY_FILTER_USE_BOTH - callback同时接受键名和键值
**/

function call($var){
    return ($var & 1); 
}

$arr = array(1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20);


print_r(array_filter($arr, "call"));


/**
Array
(
    [0] => 1
    [2] => 3
    [4] => 5
    [6] => 7
    [8] => 9
    [10] => 11
    [12] => 13
    [14] => 15
    [16] => 17
    [18] => 19
)
**/
