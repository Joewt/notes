<?php
/**
array array_unique ( array $array [, int $sort_flags = SORT_STRING ] )

array
输入的数组。

sort_flags
第二个可选参数sort_flags 可用于修改排序行为：

排序类型标记：

SORT_REGULAR - 按照通常方法比较（不修改类型）
SORT_NUMERIC - 按照数字形式比较
SORT_STRING - 按照字符串形式比较
SORT_LOCALE_STRING - 根据当前的本地化设置，按照字符串比较。
**/

$arr = array(1,1,2,2,3,4,5,6,7,7,8,9);
print_r(array_unique($arr));
