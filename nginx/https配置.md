### 1. 生成秘钥和CA证书
1. 生成key秘钥
2. 生成证书签名请求文件(csr文件) 发送给第三方
3. 生成证书签名文件(CA文件)

----
openssl genrsa -idea -out joe.key 1024  
openssl req -new -key joe.key -out joe.csr  
openssl x509 -req -days 3650 -in joe.csr -signkey joe.key -out joe.crt

### 2. Nginx的HTTPS语法配置

Syntax: ssl on | off;  
Default: ssl off;  
Context:http,server  

Syntax: ssl_certificate file;  
Default: --  
Context:http,server  

Syntax: ssl_certificate_key file;  
Default: --  
Context:http, server  
```
server
 {
   listen       443;
   server_name  xxx.xxx.xxx.xxx;

   keepalive_timeout 100;

   ssl on;
   ssl_session_cache   shared:SSL:10m;
   ssl_session_timeout 10m;

   #ssl_certificate /etc/nginx/ssl_key/joe.crt;
   ssl_certificate /etc/nginx/ssl_key/joe_apple.crt;
   ssl_certificate_key /etc/nginx/ssl_key/joe.key;
   #ssl_certificate_key /etc/nginx/ssl_key/joe_nopass.key;

   index index.html index.htm;
   location / {
       root  /opt/app/code;
   }
}
```

#### 3. 配置苹果要求的证书
a、服务器所有的连接使用TLS1.2以上版本(openssl 1.0.2)

b、HTTPS证书必须使用SHA256以上哈希算法签名

c、HTTPS证书必须使用RSA2048位或ECC256位以上公钥算法

d、使用前向加密技术

升级joe.key
openssl req -days 36500 -x509 -sha256 -nodes -newkey rsa:2048 -keyout joe.key -out joe.crt
