## PHP文件操作

### 文件信息相关函数
filetype()-获取文件类型  
filesize()-获取文件大小 返回文件大小的字节数  
filectime()-文件创建时间  
filemtime()-文件修改时间  
fileatime()-文件最后访问时间  
fileinode()-文件的inode  

检测文件权限  
file_readable()  
file_writeable()  
is_executable()  

#### 文件路径相关  
pathinfo()
basename() -返回路径中的文件名部分  
dirname() -返回路径  
file_exists() -检测文件或目录是否存在  

### 文件相关操作
创建 删除 等等  
```
touch() -创建文件  
unlink() -删除文件  
rename($file,$newfile) -重命名文件  
copy($file,$dest) -拷贝文件 --拷贝远程文件需要开启allow_url_fopen 默认是on
```
### 打开，读取文件
```
fopen 打开指定文件 如果打开二进制文件 最好加 b
fread() 读取文件内容  
ftell() 返回文件指针读/写的位置  
fseek() 在文件指针中定位       
fclose() 关闭文件句柄
```

### 写入文件
```
fwrite($handle,string $string[,int $length]) 写入文件
rewind() 重置文件指针  
ftruncate($handle,$size) 将文件截断为指定长度

fgetc() -读取一个字符
fgets() -读取一行  
fgetss() -忽略标签
```

### 操作csv文件
fgetcsv($handle) --读取CSV  
fputcsv()


### 简化操作
file_get_contents()  
file_put_contents()


json_encode 转换成json  
json_decode
