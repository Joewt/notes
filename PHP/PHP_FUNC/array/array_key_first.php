<?php
/**
mixed array_key_first ( array $array )
**/
if (!function_exists('array_key_first')) {
    /**
     * Gets the first key of an array
     *
     * @param array $array
     * @return mixed
     */
    function array_key_first(array $array)
    {
        if (count($array)) {
            reset($array);
            return key($array);
        }

        return null;
    }
}
$arr = array('a'=>1, 'b'=>2, 'c'=>3);
var_dump(array_key_first($arr));



/**
string(1) "a"
**/
