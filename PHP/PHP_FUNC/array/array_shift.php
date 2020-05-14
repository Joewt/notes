<?php
/**
mixed array_shift ( array &$array )
**/

$arr = array(1,2,3,4,5,6,7,8,9,10,11,12,13);
print_r($arr);
print_r(array_shift($arr));
print_r($arr);


echo "\n";
print_r(array_unshift($arr, 1));
print_r($arr);
