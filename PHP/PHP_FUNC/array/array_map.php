<?php
/**
array array_map ( callable $callback , array $array1 [, array $... ] )
**/

function call($var)
{
    return $var * $var;
}


$arr = array(1,2,3,4,5,6,7,8,9,10);
print_r(array_map("call", $arr));

/**
Array
(
    [0] => 1
    [1] => 4
    [2] => 9
    [3] => 16
    [4] => 25
    [5] => 36
    [6] => 49
    [7] => 64
    [8] => 81
    [9] => 100
)
**/
