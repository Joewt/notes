<?php
/**
bool array_multisort ( array &$array1 [, mixed $array1_sort_order = SORT_ASC [, mixed $array1_sort_flags = SORT_REGULAR [, mixed $... ]]] )
array1
要排序的 array。

array1_sort_order
之前 array 参数要排列的顺序。 SORT_ASC 按照上升顺序排序， SORT_DESC 按照下降顺序排序。

此参数可以和 array1_sort_flags 互换，也可以完全删除，默认是 SORT_ASC 。

array1_sort_flags
为 array 参数设定选项：

排序类型标志：

SORT_REGULAR - 将项目按照通常方法比较（不修改类型）
SORT_NUMERIC - 按照数字大小比较
SORT_STRING - 按照字符串比较
SORT_LOCALE_STRING - 根据当前的本地化设置，按照字符串比较。 它会使用 locale 信息，可以通过 setlocale() 修改此信息。
SORT_NATURAL - 以字符串的"自然排序"，类似 natsort()
SORT_FLAG_CASE - 可以组合 (按位或 OR) SORT_STRING 或者 SORT_NATURAL 大小写不敏感的方式排序字符串。
参数可以和 array1_sort_order 交换或者省略，默认情况下是 SORT_REGULAR。

...
可选的选项，可提供更多数组，跟随在 sort order 和 sort flag 之后。 提供的数组和之前的数组要有相同数量的元素。 换言之，排序是按字典顺序排列的。
**/




$ar1 = array(10, 100, 100, 0);
$ar2 = array(1, 3, 2, 4);
array_multisort($ar1, $ar2);

print_r($ar1);
print_r($ar2);


/**
Array
(
    [0] => 0
    [1] => 10
    [2] => 100
    [3] => 100
)
Array
(
    [0] => 4
    [1] => 1
    [2] => 2
    [3] => 3
)
**/
