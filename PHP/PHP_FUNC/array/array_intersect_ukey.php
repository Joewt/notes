<?php
/**
array array_intersect_ukey ( array $array1 , array $array2 [, array $... ], callable $key_compare_func )
**/

function cmp($key1, $key2)
{
    if ($key1 == $key2)
        return 0;
    else if ($key1 > $key2)
        return 1;
    else 
        return -1;
}

$array1 = array('blue'  => 1, 'red'  => 2, 'green'  => 3, 'purple' => 4);
$array2 = array('green' => 5, 'blue' => 6, 'yellow' => 7, 'cyan'   => 8);

print_r(array_intersect_ukey($array1, $array2, "cmp"));


/**
Array
(
    [blue] => 1
    [green] => 3
)
**/
