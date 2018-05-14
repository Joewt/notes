# PHP变量和常量

## PHP 变量  

判断变量类型最好用is_*  
变量函数库来判断

|类型|函数|
|:--|:--|
|整形|is_int() is_integer()|
|浮点型|is_float()|
|字符串类型|is_string()|
|布尔类型|is_bool()|
|标量类型|is_scalar()|
|空null|is_null()|
|数组|is_array()|
|对象|is_object()|
|资源|is_resource()|
|数值型或字符串|is_numeric()|


## PHP常量
* php常量是一个简单值的标识符，常量一经定义在脚本执行期间是不能改变

  
常量分类  
* 系统常量
  * PHP_VERSION 等
* 自定义常量
  * define定义  define('NAME','Joe')
    * 常量名不加$
    * 常量名最好大写
    * 常量作用域是全局
    * 常量一经定义不可改变
    * 常量值可以使标量类型，也可以是数组
  * const关键字定义常量 const NAME = 'joe'
    * 通过constant()获取常量的值
    * defined()检测常量是否存在
    * get_defined_constants() 获取已定义的所有常量
* 魔术常量
  * __LINE__ 得到当前行号
  * __FILE__ 获取文件路径名
  * __DIR__ 得到文件的完整绝对路径
  * __FUNCTION__ 得到当前函数名
  * __CLASS__ 得到当前类名
  * __METHOD__ 得到当前类的方法名
  * __TRAIT__ 获取trait名
  * __NAMESAPACE__ 获取命名空间



## PHP预定义变量
PHP提供的变量，全局变量  
分类

* 超全局变量 $GLOBALS
* 服务器和执行环境信息变量 $_SERVERS
* 环境变量  $_ENV
* HTTP cookie $_COOKIE
* HTTP session变量  $_SESSION
* 文件上传信息  $_FILES
* $_GET
* $_POST

