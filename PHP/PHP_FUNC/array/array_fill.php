<?php
/**
array array_fill ( int $start_index , int $num , mixed $value )
start_index
返回的数组的第一个索引值。

如果 start_index 是负数， 那么返回的数组的第一个索引将会是 start_index ，而后面索引则从0开始。 (参见 例子)。

num
插入元素的数量。 必须大于或等于 0。

value
用来填充的值。
**/

$arr = array_fill(2,10, "joe");
print_r($arr);

/**
Array
(
    [2] => joe
    [3] => joe
    [4] => joe
    [5] => joe
    [6] => joe
    [7] => joe
    [8] => joe
    [9] => joe
    [10] => joe
    [11] => joe
)
**/
