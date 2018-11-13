<?php
/**
array array_column ( array $input , mixed $column_key [, mixed $index_key = null ] )
input
需要取出数组列的多维数组。 如果提供的是包含一组对象的数组，只有 public 属性会被直接取出。 为了也能取出 private 和 protected 属性，类必须实现 __get() 和 __isset() 魔术方法。

column_key
需要返回值的列，它可以是索引数组的列索引，或者是关联数组的列的键，也可以是属性名。 也可以是NULL，此时将返回整个数组（配合index_key参数来重置数组键的时候，非常管用）

index_key
作为返回数组的索引/键的列，它可以是该列的整数索引，或者字符串键值。
**/

$arr = array(
    array(
        'id' => 2135,
        'first_name' => 'John',
        'last_name' => 'Doe',
    ),
    array(
        'id' => 3245,
        'first_name' => 'Sally',
        'last_name' => 'Smith',
    ),
    array(
        'id' => 5342,
        'first_name' => 'Jane',
        'last_name' => 'Jones',
    ),
    array(
        'id' => 5623,
        'first_name' => 'Peter',
        'last_name' => 'Doe',
    )
);
 
$first_names = array_column($arr, 'first_name');
print_r($first_names);

$last_name = array_column($arr, 'last_name', 'id');
print_r($last_name);

/**
Array
(
    [0] => John
    [1] => Sally
    [2] => Jane
    [3] => Peter
)
Array
(
    [2135] => Doe
    [3245] => Smith
    [5342] => Jones
    [5623] => Doe
)
**/
