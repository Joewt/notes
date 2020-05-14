<?php
/**
array array_intersect_key ( array $array1 , array $array2 [, array $... ] )

**/

$array1 = array('blue'  => 1, 'red'  => 2, 'green'  => 3, 'purple' => 4);
$array2 = array('green' => 5, 'blue' => 6, 'yellow' => 7, 'cyan'   => 8);

print_r(array_intersect_key($array1, $array2));


/**
Array
(
    [blue] => 1
    [green] => 3
)
**/
