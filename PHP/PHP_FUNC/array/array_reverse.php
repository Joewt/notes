<?php
/**
array array_reverse ( array $array [, bool $preserve_keys = false ] )
array
输入的数组。

preserve_keys
如果设置为 TRUE 会保留数字的键。 非数字的键则不受这个设置的影响，总是会被保留。
**/


$arr = array('a'=>1, array('b'=>2, 'c'=>3), array('color'=>'red'), 2,3,4,5,6);

print_r(array_reverse($arr));
print_r(array_reverse($arr, TRUE));
