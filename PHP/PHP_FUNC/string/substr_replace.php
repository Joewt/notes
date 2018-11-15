<?php
/**
mixed substr_replace ( mixed $string , mixed $replacement , mixed $start [, mixed $length ] )
substr_replace() 在字符串 string 的副本中将由 start 和可选的 length 参数限定的子字符串使用 replacement 进行替换。
**/
$str = 'ABCDEFG';
echo substr_replace($str, 'abc', 0, 4);
