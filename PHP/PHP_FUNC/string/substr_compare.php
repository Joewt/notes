<?php

/**
int substr_compare ( string $main_str , string $str , int $offset [, int $length [, bool $case_insensitivity = FALSE ]] )
**/

$str1 = 'abcdef';
$str2 = 'ab';
echo substr_compare($str1, $str2, 1, 2);
echo substr($str1, 1, 2);
