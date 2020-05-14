<?php
/**
mixed array_search ( mixed $needle , array $haystack [, bool $strict = false ] )
needle
搜索的值。

Note:
如果 needle 是字符串，则比较以区分大小写的方式进行。
haystack
这个数组。

strict
如果可选的第三个参数 strict 为 TRUE，则 array_search() 将在 haystack 中检查完全相同的元素。 这意味着同样严格比较 haystack 里 needle 的 类型，并且对象需是同一个实例。
**/


$array = array(0 => 'blue', 1 => 'red', 2 => 'green', 3 => 'red');

print_r(array_search('green', $array)); 
print_r(array_search('red', $array));

print_r(array_search('xxx', $array));


