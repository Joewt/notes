<?php

/**

mixed array_reduce ( array $array , callable $callback [, mixed $initial = NULL ] )

array
输入的 array。

callback
mixed callback ( mixed $carry , mixed $item )
carry
携带上次迭代里的值； 如果本次迭代是第一次，那么这个值是 initial。

item
携带了本次迭代的值。

initial
如果指定了可选参数 initial，该参数将在处理开始前使用，或者当处理结束，数组为空时的最后一个结果。

**/


function sum($carry, $item)
{
    var_dump($carry, $item);
    $carry *= $item;
    return $carry;
}

$arr = array(1,2,3,4,5);

print_r(array_reduce($arr, "sum"));
