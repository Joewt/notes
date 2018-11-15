<?php

$email  = 'name@example.com';
$domain = strstr($email, '@');
echo $domain."\n"; 

$user = strstr($email, '@', true); 
echo $user; 
