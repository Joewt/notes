# PHP系统函数

### 字符串相关函数

strlen 函数
描述： 用于获取字符长度
语法： int strlen(string $str)  
```
$str = 'abcdef';
echo strlen($str);
```

#### 大小写转换函数  
strtolower函数  
描述：将字符串转换成小写  
语法：string strtolower(string $str)  

strtoupper函数  
描述：将字符串转换成大写  
语法：string strtoupper(string $str)  

ucfirst函数  
描述：将句子首字母转换为大写
语法：string ucfirst(string $str);

ucwords函数  
描述：将每个单词的首字母转换为大写  
语法：string ucwords(string $str);  


#### 字符串替换函数

```
str_replace函数  
描述：实现字符替换，区分大小写  
语法：mixed str_replace(mixed $search,mixed replace,mixed $subject, [int &$count]);  

str_ireplace函数
描述：实现字符替换，不区分大小写  
语法：mixed str_ireplace(mixed $search,mixed replace,mixed $subject, [int &$count]); 
```

#### 与html实体相关的函数

```
htmlspecialchars函数  
描述：预定义的字符转换为html实体  
语法：strng htmlspecialchars(string $string,[,int $flags =ENT_COMPAT])  
说明：flag参数规定如何处理引号，其值可以为：
ENT_COMPAT - 默认，仅编码双引号  
ENT_QUOTES - 编码双引号和单引号  
ENT_NOQUOTES - 不编码任何引号  
```

#### 删除空格或其他字符相关的函数  

trim函数
描述：实现删除字符串开始和结束的位置删除空格或者其他字符  
语法：string trim(string $str)


#### 字符串位置相关的函数  

strpos-返回字符在另一个字符第一次出现的位置  
stripos-忽略大小写  

strrpos-返回的是最后一次出现的位置  
strripos-忽略大小写

#### 字符串截取函数

```
substr函数  
描述：实现字符串截取  
语法：string substr(string $string ,int $start,[,$length])

strstr函数
描述：将搜索一个字符串在另一个字符串中的第一次出现的位置，区分大小写
语法：string strstr(string $haystack, mixed $needle)
说明：函数返回字符串剩余的部分

stristr函数
不区分大小写

strrchr函数
描述：将搜索字符串在另一个字符串中最后一次出现的位置
语法：string strrchr(string $haystack, mixed $needle)

```

#### 翻转字符串
strrev函数
描述：反转字符串
语法：string strrev(string $string)

#### 打乱字符串
str_shuffle函数
描述：随机打乱字符串  
语法：string str_shuffle(string $str)


#### 分割字符串 
```
explode函数
描述：使用一个字符串分割另一个字符串  
语法：array explode (string $delimiter,string $string [,int $limit])

$str = 'a|b|c|d';
$arr = explode('|',$str)

implode函数
描述：将一个一维数组的值转换为字符串  
语法：


```

#### 字符串格式化
sprintf函数  
描述：格式化字符串  
语法：string sprintf(string $format,[,mixed $args[,mixed $...]])  
说明：
如果%符号多余arg参数，则必须使用占位符，占位符位于%符号之后，由数字和"\$"组成

### PHP数学函数

#### 向下或向上取整

floor函数  
描述：将实现舍一取整  
语法：float floor(float $value)  

ceil函数  
描述：将实现进一取整  
语法：float ceil(float $value)  

#### 幂运算、平方根  

pow函数  
描述：幂指数运算  
语法：number pow(number $base, number $exp)  

sqrt函数  
描述：平方根  
语法：float sqrt(float $avg)  

#### 最大值，最小值

max函数  
描述：求最大值  
语法：mixed max(mixed $value,mixed $value,...)  

min函数  
描述：求最小值  
语法：mixed min(mixed $value,mixed $value,...)

#### 随机数  
rand函数  
描述：将产生随机数  
语法：int rand(int $min,int $max) 

mt_rand函数  
描述：将产生一个更好的随机数  
语法：int mt_rand(int $min,int $max)  

#### 四舍五入  
```
round函数  
描述：实现四舍五入的功能  
语法：float round(float $val[,int $precision=0])

precision-表示小数点的位数

```

#### 数字格式化  
number_format函数 
描述：数字格式化
语法：string number_format(float $val)  

#### 浮点数余数  
fmod函数  
描述：将返回除法的浮点数余数  
语法：float fmod(flaot $x, float $y)

### 日期函数库  


#### 格式化日期 
```
date函数
描述：格式化一个本地时间/日期  
语法：string date(string format,[,int timestamp])
```
|format字符|说明|
|:--|:--|
|Y|4位数字完整表示的年份|
|y|2位数字表示的年份|
|F|月份，完整的文本格式|
|M|3个字母缩写表示的月份|
|m|数字表示的月份，有前导0|
|n|数字表示的月份，没有前导0|
|d|月份中的第几天，有前导0|
|j|月份中的第几天，没有前导0|
|I|星期几，完整的文本格式|
|D|星期中的第几天，文本表示，3个字母|
|w|星期中的第几天，数字表示|
|H|小时，24小时格式，有前导0|
|i|有前导0的分钟数|
|s|秒数，有前导0|

通过date_default_timezone_set('Asia/Shanghai')修改时区  
通过date_default_timezone_get()获取时区  
也可以通过修改配置文件更改默认时区  


#### 时间戳和字符串转时间戳

```
time()获取时间戳
strtotime函数
描述：将字符串转换成Unix时间戳  
语法：int strtotime(string $time[,int $now = time()])
```
日期时间格式  
|||
|:--|:--|
|dayname|sunday\|monday\|tuesday\|wednesday\|thursday\|friday\|saturday\|sun\|mon\|tue\|wed\|thu\|fri\|sat|
|daynext|weekday\|weekdays|
|number|\[+\|-]?\|\[0~9]+|
|ordinal|first\|second\|third\|fourth\|fifth\|sixth\|seventh\|eight\|ninth\|tenth\|eleventh\|twelfth\|next\|last\|previous\|this|
|reltext|next\|last\|previous\|this|
|unit|((sec\|second\|min\|minute\|hour\|day\|fortnight\|forthnight\|month\|year)s?\|weeks\|daytext|  

以天为基础的格式  
|||
|:--|:--|
|yesterday|昨天午夜|
|midnight|午夜|
|today|今天|
|noon|中午|
|tomorrow|明天午夜|
|first day of ??|某月第一天|
|last day of ??|某月最后一天|

如何获取上一月的日期

```
date('Y:m:d H:i:s',strtotime('last day of -1 month'))

```

#### 获取时间戳和微秒数 
```
microtime函数
描述：返回当前Unix时间戳和微秒数  
语法：mixed microtime([bool $get_as float])
```

#### 生成唯一ID  
```
uniqid函数  
描述：生成唯一ID  
语法：string uniqid([string $prefix="",[,bool $more_entropy = false]])

md5(uniqid(microtime().mt_rand()))

```

