<?php
/**
array array_diff_key ( array $array1 , array $array2 [, array $... ] )
array1
从这个数组进行比较

array2
针对此数组进行比较

...
更多比较数组
**/

$array1 = array('blue'  => 1, 'red'  => 2, 'green'  => 3, 'purple' => 4);
$array2 = array('green' => 5, 'blue' => 6, 'yellow' => 7, 'cyan'   => 8);

var_dump(array_diff_key($array1, $array2));


/**

array(2) {
  ["red"]=>
  int(2)
  ["purple"]=>
  int(4)
}

**/
