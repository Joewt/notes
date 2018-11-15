<?php

/**
string htmlspecialchars ( string $string [, int $flags = ENT_COMPAT | ENT_HTML401 [, string $encoding = ini_get("default_charset") [, bool $double_encode = TRUE ]]] )
某类字符在 HTML 中有特殊用处，如需保持原意，需要用 HTML 实体来表达。 本函数会返回字符转义后的表达。 如需转换子字符串中所有关联的名称实体，使用 htmlentities() 代替本函数。

如果传入字符的字符编码和最终的文档是一致的，则用函数处理的输入适合绝大多数 HTML 文档环境。 然而，如果输入的字符编码和最终包含字符的文档是不一样的， 想要保留字符（以数字或名称实体的形式），本函数以及 htmlentities() （仅编码名称实体对应的子字符串）可能不够用。 这种情况可以使用 mb_encode_numericentity() 代替。

执行转换
字符	替换后
& (& 符号)	&amp;
" (双引号)	&quot;，除非设置了 ENT_NOQUOTES
' (单引号)	设置了 ENT_QUOTES 后， &#039; (如果是 ENT_HTML401) ，或者 &apos; (如果是 ENT_XML1、 ENT_XHTML 或 ENT_HTML5)。
< (小于)	&lt;
> (大于)	&gt;
**/



$str = '&<>>/""';

echo htmlspecialchars($str)."\n";


$new_str = htmlspecialchars($str);


echo htmlspecialchars_decode($new_str);
