<?php


function print_($val)
{
    echo $val."\n";
}
$transport = array('foot', 'bike', 'car', 'plane');
$mode = current($transport); 
print_($mode);
$mode = next($transport);    
print_($mode);
$mode = current($transport); 
print_($mode);
$mode = prev($transport);    
print_($mode);
$mode = end($transport);
print_($mode);
$mode = current($transport); 
print_($mode);

array_push($transport, "xxx");
$mode = each($transport);
print_r($mode);
print_(current($transport));
reset($transport);
print_(current($transport));


shuffle($transport);

print_r($transport);
