<?php
/**
array array_intersect ( array $array1 , array $array2 [, array $... ] )
**/

$array1 = array("a" => "green", "red", "blue");
$array2 = array("b" => "green", "yellow", "red");
$result = array_intersect($array1, $array2);
print_r($result);

/**
Array
(
    [a] => green
    [0] => red
)
**/
