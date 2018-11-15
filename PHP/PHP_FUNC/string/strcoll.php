<?php

/**
int strcoll ( string $str1 , string $str2 )
**/


$a = 'a';
$b = 'B';

print strcmp ($a, $b) . "\n"; // prints 1

setlocale (LC_COLLATE, 'C');
print "C: " . strcoll ($a, $b) . "\n"; // prints 1

setlocale (LC_COLLATE, 'de_DE');
print "de_DE: " . strcoll ($a, $b) . "\n"; // prints -2

setlocale (LC_COLLATE, 'de_CH');
print "de_CH: " . strcoll ($a, $b) . "\n"; // prints -2

setlocale (LC_COLLATE, 'en_US');
print "en_US: " . strcoll ($a, $b) . "\n"; // prints -2
