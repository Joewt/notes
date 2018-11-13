<?php
/**
bool array_key_exists ( mixed $key , array $array )
**/


$search_array = array('first' => 1, 'second' => 4);

if(array_key_exists("second", $search_array)){
    echo "YES";
} else {
    echo "NO";
}


/**

YES
**/
