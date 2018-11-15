<?php

$str = 'abcdefdfsdf';
$str1 = 'abcEEF';

echo strncasecmp($str, $str1, 5);
