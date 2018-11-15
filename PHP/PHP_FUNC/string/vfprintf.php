<?php
/**
int vfprintf ( resource $handle , string $format , array $args )
**/

if(!($fp = fopen('data.txt','w')))
    return;

$t = date("Y:m:d",time());
list($year,$month,$day) = explode(':', $t);
vfprintf($fp, "%04d-%02d-%02d", array($year, $month, $day));
