<?php
$array = array('lastname', 'email', 'phone');

echo implode(':', $array);
echo "\n";

echo join('--', $array);
