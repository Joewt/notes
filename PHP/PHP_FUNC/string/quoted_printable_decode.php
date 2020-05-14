<?php
$str = 'Hello=0Aworld.';

$new_str = quoted_printable_decode($str);
echo $new_str;
echo "\n";


echo quoted_printable_encode($new_str);
