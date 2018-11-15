<?php
/**
string quotemeta ( string $str )
返回 在下面这些特殊字符前加 反斜线(\) 转义后的字符串。 这些特殊字符包含：

. \ + * ? [ ^ ] ( $ )
**/


$str = './+*?,/[]^() sdfasf ()';
echo quotemeta($str);
