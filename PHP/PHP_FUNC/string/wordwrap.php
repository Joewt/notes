<?php
/**
string wordwrap ( string $str [, int $width = 75 [, string $break = "\n" [, bool $cut = FALSE ]]] )
str
输入字符串。

width
列宽度。

break
使用可选的 break 参数打断字符串。

cut
如果 cut 设置为 TRUE，字符串总是在指定的 width 或者之前位置被打断。因此，如果有的单词宽度超过了给定的宽度，它将被分隔开来。（参见第二个范例）。 当它是 FALSE ，函数不会分割单词，哪怕 width 小于单词宽度。
**/


$text = "The quick brown fox jumped over the lazy dog.";

$new = wordwrap($text, 5);

echo $new;
