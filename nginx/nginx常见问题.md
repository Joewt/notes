### 1. 相同server_name多个虚拟主机优先级访问
会先匹配最先配置的  
### 2. location匹配优先级
= 进行普通字符精确匹配，也就是完整匹配
^~ 表示普通字符匹配，使用前缀匹配
~ \~* 表示执行一个正则匹配()
### 3. try_files使用
按顺序检查文件是否存在
```
location / {
  try_files $uri $uri/ /index.php;
}

#下面的配置 先匹配root 对应的目录 如果没有 则转到相应的page下的相应配置
location / {
  root /opt/app/code;
  tyr_files $uri @page;
}
location @page {
  proxy_pass http://127.0.0.1:8080;
}
```
### 3. nginx alias 和 root 区别
```
location /request_path/image/{
  root/local_path/image/;
}
http://www.sername.com/request_path/image/cat.png

/local_path/image/request_path/image/cat.png


location /requeste_path/image/{
  alias /local_path/image/;
}
http://www.sername.com/request_path/image/cat.png

/local_path/image/cat.png
```
### 4. 如何获取用户的真实IP地址
通过在第一级代理设置一个x_real_ip变量
set x_real_ip=\$remote_addr  
$x_real_ip=IP
### 5. nginx中常见的错误码
Nginx: 413 Request Entity Too Large
1. 用户上传文件限制 client_max_body_size  
503 bad gateway
2. 后端服务无响应
504 Gateway Time-out
3. 后端服务执行超时