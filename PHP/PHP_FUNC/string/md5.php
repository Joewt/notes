<?php

$str = 'joe';

echo md5($str)."\n";
$str_file = 'README.md';

echo md5_file($str_file);
