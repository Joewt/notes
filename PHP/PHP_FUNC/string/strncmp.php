<?php
$s1 = 'joe hello world';
$s2 = 'joe hello';

echo strncmp($s1, $s2, 2)."\n";
echo strncmp($s1, $s2, 10);
