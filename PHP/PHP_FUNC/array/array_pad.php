<?php
/**
array array_pad ( array $array , int $size , mixed $value )
**/


$arr = array(1,2,3,44,9);
print_r(array_pad($arr, 10, 111));

/**
Array
(
    [0] => 1
    [1] => 2
    [2] => 3
    [3] => 44
    [4] => 9
    [5] => 111
    [6] => 111
    [7] => 111
    [8] => 111
    [9] => 111
)
**/
