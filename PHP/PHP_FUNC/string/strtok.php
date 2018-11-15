<?php
/**
strtok() 将字符串 str 分割为若干子字符串，每个子字符串以 token 中的字符分割。这也就意味着，如果有个字符串是 "This is an example string"，你可以使用空格字符将这句话分割成独立的单词。

注意仅第一次调用 strtok 函数时使用 string 参数。后来每次调用 strtok，都将只使用 token 参数，因为它会记住它在字符串 string 中的位置。如果要重新开始分割一个新的字符串，你需要再次使用 string 来调用 strtok 函数，以便完成初始化工作。注意可以在 token 参数中使用多个字符。字符串将被该参数中任何一个字符分割。
**/


$string = "This is\tan example\nstring";
$t = strtok($string, "\n\t");
while ($tok !== false) {
    echo "Word=$tok\n";
    $tok = strtok(" \n\t");
}
