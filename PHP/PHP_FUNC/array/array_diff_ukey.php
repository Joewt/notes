<?php
/**
array array_diff_ukey ( array $array1 , array $array2 [, array $... ], callable $key_compare_func )

array1
The array to compare from

array2
An array to compare against

...
More arrays to compare against

key_compare_func
在第一个参数小于，等于或大于第二个参数时，该比较函数必须相应地返回一个小于，等于或大于 0 的整数。

int callback ( mixed $a, mixed $b )
**/

function key_compare_func($key1, $key2)
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

var_dump(array_diff_ukey($array1, $array2, 'key_compare_func'));

/**
array(2) {
  ["red"]=>
  int(2)
  ["purple"]=>
  int(4)
}
**/
