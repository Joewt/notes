<?php
/**
array array_intersect_uassoc ( array $array1 , array $array2 [, array $... ], callable $key_compare_func )
**/

$array1 = array("a" => "green", "b" => "brown", "c" => "blue", "red");
$array2 = array("a" => "GREEN", "B" => "brown", "yellow", "red");

print_r(array_intersect_uassoc($array1, $array2, "strcasecmp"));

/**
Array
(
    [b] => brown
)
**/
