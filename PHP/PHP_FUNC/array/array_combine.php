<?php
/**
array array_combine ( array $keys , array $values )
keys
将被作为新数组的键。非法的值将会被转换为字符串类型（string）。

values
将被作为 Array 的值。
**/

$arr_k = array('a', 'b', 'c');
$arr_v = array(1, 2, 3);

print_r(array_combine($arr_k, $arr_v));


$a = array('green', 'red', 'yellow');
$b = array('avocado', 'apple', 'banana');
$c = array_combine($a, $b);

print_r($c);


/**
Array
(
    [a] => 1
    [b] => 2
    [c] => 3
)
Array
(
    [green] => avocado
    [red] => apple
    [yellow] => banana
)
**/
