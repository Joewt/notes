# Nginx学习笔记
### 一、Nginx的安装
&emsp;&emsp;这里推荐使用yum安装如果是debian使用apt-get  

进入[官网](http://www.nginx.org)可以下载二进制包自行编译安装  
这里推荐一个快捷方法  
1. 创建一个文件 /etc/yum.repos.d/nginx.repo  
文件内容如下:
```
[nginx]
name=nginx repo
baseurl=http://nginx.org/packages/OS/OSRELEASE/$basearch/
gpgcheck=0
enabled=1
```
2. “OS” 用 “rhel” 或 “centos”取代  
“OSRELEASE”根据自己的系统版本号设置如果是6.x的改为6同理7.x的类似  
3. 执行yum install nginx安装即可

### 二、Nginx目录说明

<table>
  <tr>
    <th width=50%, bgcolor=#be77ff >路径</th>
    <th width=15%, bgcolor=#be77ff>类型</th>
    <th width="35%", bgcolor=#921aff>作用</th>
  </tr>
  <tr>
    <td bgcolor=#eeeeee> /etc/nginx/koi-utf<br>/etc/nginx/koi-win<br>/etc/nginx/win-utf </td>
    <td> 配置文件  </td>
    <td>  编码转换映射转化文件   </td>
  </tr>
  <tr>
    <td bgcolor=#eeeeee> /etc/nginx<br>/etc/nginx/nginx.conf<br>/etc/nginx/conf.d<br>/etc/nginx/conf.d/default.conf </td>
    <td> 目录、配置文件  </td>
    <td>  Nginx主配置文件   </td>
  </tr>
  <tr>
    <td bgcolor=#eeeeee> /etc/nginx/fastcgi_params<br>/etc/nginx/uwsgi_params<br>/etc/nginx/scgi_params</td>
    <td> 配置文件  </td>
    <td>  cgi配置相关，fastcgi配置   </td>
  </tr>
  <tr>
    <td bgcolor=#eeeeee>
    /etc/nginx/mime.types</td>
    <td>配置文件</td>
    <td>设置http协议d额Content-Type与扩展名对应关系</td>
  </tr>
  <tr>
    <td bgcolor=#eeeeee>
    /usr/sbin/nginx<br>/usr/sbin/nginx-debug</td>
    <td>
    命令
    </td>
    <td>
    Nginx服务的启动管理的终端命令
    </td>
  </tr>
</table>
部分编译选项
<table>
<tr>
  <th width=50%, bgcolor=#be77ff >编译选项</th>
  <th width=50%, bgcolor=#be77ff>作用</th>
</tr>
<tr>
  <th>--with-cc-opt=parameters</th>
  <th>设置额外的参数将被添加CFLAGS变量</th>
</tr>
<tr>
  <th>--with-ld-opt=parameters</th>
  <th>设置附加的参数，链接系统库</th>
</tr>
<tr>
  <th>--with-http_sub_module</th>
  <th>HTTP内容替换</th>
</tr>
<tr>
  <th>--with-http_random_index_module</th>
  <th>目录中选择一个随机主页</th>
</tr>
</table>

### 三、Nginx的基本使用
nginx -V 可以查看nginx安装了哪些模块  
nginx -tc /etc/nginx/nginx.conf 检查配置文件是否出错  
nginx -s reload -c /etc/nginx/nginx.conf 平滑重启nginx  
