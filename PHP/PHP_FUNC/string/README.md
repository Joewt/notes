# 字符串函数
* [wordwrap](wordwrap.php) -- 打断字符串为指定数量的字串
* [printf](printf.php) -- 输出格式化字符串
* [fprintf](fprintf.php) -- 将格式化后的字符串写入到流
* [print](print.php) -- 输出字符串(只能输出一个)
* [echo](echo.php) -- 输出字符串(可以输出多个)
* [sprintf](sprintf.php) -- 返回一个格式化的字符串
* [vsprintf](vsprintf.php) -- 返回格式化字符串
* [vprintf](vprintf.php) -- 输出格式化字符串
* [vfprintf](vfprintf.php) -- 将格式化字符串写入流
* [sscanf](sscanf.php) -- 根据指定格式解析输入的字符
* [ucwords](ucwords.php) -- 将字符串中每个单词的首字母转换为大写
* [trim](trim.php) -- 去除字符串首尾处的空白字符（或者其他字符）
* [rtrim](rtrim.php) -- 删除字符串末端的空白字符（或者其他字符）
* [substr](substr.php) -- 返回字符串的子串
* [substr_replace](substr_replace.php) -- 替换字符串的子串
* [substr_count](substr_count.php) -- 计算字串出现的次数
* [substr_compare](substr_compare.php) -- 二进制安全比较函数
* [strtr](strtr.php) -- 转换指定字符
* [strtoupper](strtoupper.php) -- 将字符串转化为大写
* [strtolower](strtolower.php) -- 将字符串转化为小写
* [strtok](strtok.php) -- 标记分割字符串
* [strstr](strstr.php) -- 查找字符串的首次出现
* [strchr](strstr.php) -- 别名strstr
* [stristr](strstr.php) -- strstr 函数的忽略大小写版本
* [strspn](strspn.php) -- 计算字符串中全部字符都存在于指定字符集合中的第一段子串的长度。
* [strcspn](strcspn.php) -- 获取不匹配遮罩的起始子字符串的长度
* [strrpos](strrpos.php) --  计算指定字符串在目标字符串中最后一次出现的位置
* [strripos](strripos.php) --  计算指定字符串在目标字符串中最后一次出现的位置(不区分大小写)
* [strrev](strrev.php) -- 反转字符串
* [strrchr](strrchr.php) -- 查找指定字符在字符串中的最后一次出现
* [strpos](strpos.php) -- 查找字符串首次出现的位置
* [stripos](strpos.php) -- 查找字符串首次出现的位置（不区分大小写）
* [strpbrk](strpbrk.php) -- 在字符串中查找一组字符的任何一个字符
* [strcmp](strcmp.php) -- 二进制安全字符串比较
* [strcasecmp](strcasecmp.php) -- 二进制安全比较字符串（不区分大小写）
* [strncmp](strncmp.php) -- 二进制安全比较字符串开头的若干个字符
* [strncasecmp](strncasecmp.php) -- 二进制安全比较字符串开头的若干个字符（不区分大小写）
* [strnatcmp](strnatcmp.php) -- 使用自然排序算法比较字符串
* [strnatcasecmp](strnatcasecmp.php) -- 使用“自然顺序”算法比较字符串（不区分大小写） 
* [strlen](strlen.php) -- 获取字符串长度
* [stripslashes](stripslashes.php) -- 反引用一个引用字符串
* [strip_tags](strip_tags.php) -- 从字符串中去除 HTML 和 PHP 标记
* [strcoll](strcoll.php) -- 基于区域设置的字符串比较
* [str_word_count](str_word_count.php) -- 返回字符串中单词的使用情况
* [str_split](str_split.php) -- 将字符串转换为数组
* [str_shuffle](str_shuffle.php) -- 随机打乱一个字符串
* [str_rot13](str_rot13.php) -- 对字符串执行 ROT13 转换
* [str_replace](str_replace.php) -- 子字符串替换
* [str_ireplace](str_replace.php) -- 子字符串替换（忽略大小写）
* [str_repeat](str_repeat.php) -- 重复一个字符串
* [str_pad](str_pad.php) -- 使用另一个字符串填充字符串为指定长度
* [str_getcsv](str_getcsv.php) -- 解析csv字符串为一个数组
* [soundex](soundex.php) -- 返回字符串的soundex 键
* [similar_text](similar_text.php) -- 计算两个字符串的相似度
* [sha1](sha1.php) -- 计算字符串的 sha1 散列值
* [sha1_file](sha1_file.php) -- 计算文件的 sha1 散列值
* [setlocale](setlocale.php) -- 设置地区信息
* [quotemeta](quotemeta.php) -- 转义元字符集
* [quoted_printable_decode](quoted_printable_decode.php) -- 将 quoted-printable 字符串转换为 8-bit 字符串
* [quoted_printable_encode](quoted_printable_decode.php) -- 将 8-bit 字符串转换成 quoted-printable 字符串
* [ord](ord.php) -- 返回字符的 ASCII 码值
* [number_format](number_format.php) -- 以千位分隔符方式格式化一个数字
* [money_format](number_format.php) -- 将数字格式化成货币字符串
* [nl2br](nl2br.php) -- 在字符串所有新行之前插入 HTML 换行标记
* [md5](md5.php) -- 计算字符串的 MD5 散列值
* [md5_file](md5_file.php) -- 计算指定文件的 MD5 散列值
* [lcfirst](lcfirst.php) -- 使一个字符串的第一个字符小写
* [implode](implode.php) -- 将一个一维数组的值转化为字符串
* [join](implode.php) -- implode的别名
* [htmlspecialchars](htmlspecialchars.php) -- 将特殊字符转换为 HTML 实体
* [htmlspecialchars_decode](htmlspecialchars_decode.php) -- 将特殊的 HTML 实体转换回普通字符
* [htmlentities](htmlentities.php) -- 将字符转换为 HTML 转义字符
* [hex2bin](hex2bin.php) -- 转换十六进制字符串为二进制字符串
* [bin2hex](bin2hex.php) -- 把包含数据的二进制字符串转换为十六进制
* [explode](explode.php) -- 使用一个字符串分割另一个字符串
* [crypt](crypt.php) -- 单向字符串散列
* [crc32](crc32.php) -- 计算一个字符串的 crc32 多项式
* [count_chars](count_chars.php) -- 返回字符串所用字符的信息
* [chr](chr.php) -- 返回指定的字符
* [chunk_split](chunk_split.php) -- 将字符串分割成小块
* [chop](rtrim.php) -- rtrim的别名
* [addcslashes](addcslashes.php) -- 以 C 语言风格使用反斜线转义字符串中的字符
* [addslashes](addslashes.php) -- 使用反斜线引用字符串
