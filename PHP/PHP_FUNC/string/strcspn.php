<?php
/**
int strcspn ( string $str1 , string $str2 [, int $start [, int $length ]] )
返回 str1 中，所有字符都不存在于 str2 范围的起始子字符串的长度。
**/


$a = strcspn('abcd',  'apple');
$b = strcspn('abcd',  'banana');
$c = strcspn('hello', 'l');
$d = strcspn('hello', 'world');

var_dump($a);
var_dump($b);
var_dump($c);
var_dump($d);
