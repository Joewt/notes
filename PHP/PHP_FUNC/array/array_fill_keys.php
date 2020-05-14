<?php
/**
array array_fill_keys ( array $keys , mixed $value )
keys
使用该数组的值作为键。非法值将被转换为字符串。

value
填充使用的值。
**/


$keys = array('foo', 5, 10, 'bar', 'a', 'b', 11);

print_r(array_fill_keys($keys, 'abcdef'));


/**
Array
(
    [foo] => abcdef
    [5] => abcdef
    [10] => abcdef
    [bar] => abcdef
    [a] => abcdef
    [b] => abcdef
    [11] => abcdef
)
**/

