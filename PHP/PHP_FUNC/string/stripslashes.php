<?php
/**
string stripslashes ( string $str )

str
输入字符串。

返回值

返回一个去除转义反斜线后的字符串（\' 转换为 ' 等等）。双反斜线（\\）被转换为单个反斜线（\）
**/
$str = "Is your name O\'reilly?";

echo stripslashes($str);
