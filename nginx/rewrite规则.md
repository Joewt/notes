### 配置语法
  Syntax: rewrite regex replace[flag] 

  Default:

  Context:server,location,if

### 正则表达式
|.|匹配除换行以外的任意字符|
|:--|:--|
|?|重复0次或1次|
|+|重复1次或更多次|
|*|最少链接数，哪个机器连接数少就分发|
|\d|匹配数字|
|^|匹配字符串的开始|
|$|匹配字符串的介绍|
|{n}|重复n次|
|{n,}|重复n次或更多次|
|[c]|匹配单个字符c|
|[a-z]|匹配a-z小写字母的任意一个|

|\转义字符|
|:--|
|rewrite index\.php$/pages/maintain.html break|

|()用于匹配括号之间的内容，通过\$1,\$2调用|
|-|
```
  if($http_user_agent ~ MSIE) {
    rewrite ^(.*)$ /msie/$1 break;
  }
```
正则表达式了解透彻
终端测试命令pcretest


### flag
| last | 停止rewrite检测 会接着匹配规则  |
|:--|:--|
| break | 停止rewrite检测 不会匹配 返回404|
| redirect | 返回302临时重定向 |
| permanent | 返回301永久重定向 |

几个实例
```
  #如果是Chrome浏览器  访问 nginx则跳转到baidu.com  
  if($http_user_agent ~* Chrome) {
    rewrite ^/nginx http://baidu.com redirect;
  }
  # -f 判断 路径是否存在  !-f 表示 不存在 则跳转到对应目录
  if (!-f $request_filename) {
    rewrite ^(.*)$ http://www.baidu.com/$1 redirect;
  }

```
### Rewrite规则优先级
1. 执行server块的rewrite指令
2. 执行location匹配
3. 执行选定的location中的rewrite
