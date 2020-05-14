<?php
/**
array compact ( mixed $varname1 [, mixed $... ] )
**/


$name = "joe";
$age = 18;

$res = compact("name", "age");
print_r($res);
