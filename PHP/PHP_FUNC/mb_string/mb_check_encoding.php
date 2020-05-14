<?php
/**
bool mb_check_encoding ([ string $var = NULL [, string $encoding = mb_internal_encoding() ]] )
检查指定的字节流在指定的编码里是否有效。它能有效避免所谓的“无效编码攻击（Invalid Encoding Attack）”。
**/


$string = "\x00\x81";

$encoding = "Shift_JIS";

if(mb_check_encoding($string, $encoding)){
    echo 'yes';
} else {
    echo 'no';
}
