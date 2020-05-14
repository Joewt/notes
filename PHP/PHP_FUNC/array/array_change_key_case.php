<?php
/**
array array_change_key_case ( array $array [, int $case = CASE_LOWER ] )
将array数组中的所有键名改为全大写或全小写
默认为CASE_LOWER
**/

$arr = array("OneONE"=>1, "TWoTWO"=>2);
print_r(array_change_key_case($arr, CASE_UPPER));


print_r(array_change_key_case($arr));

