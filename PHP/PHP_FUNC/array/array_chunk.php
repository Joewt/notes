<?php
/**
array array_chunk ( array $array , int $size [, bool $preserve_keys = false ] )
array
需要操作的数组

size
每个数组的单元数目

preserve_keys
设为 TRUE，可以使 PHP 保留输入数组中原来的键名。如果你指定了 FALSE，那每个结果数组将用从零开始的新数字索引。默认值是 FALSE。
**/


$arr = array('a', 'b', 'c', 'd', 'e');
print_r(array_chunk($arr, 2));
print_r(array_chunk($arr, 2, TRUE));
