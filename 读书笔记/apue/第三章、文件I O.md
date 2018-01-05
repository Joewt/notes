## 3.2 文件描述符
&emsp;&emsp;对于内核而言，所有打开的文件都通过文件描述符引用，文件描述符是一个非负整数。  
&emsp;&emsp;**UNIX**系统shell把文件描述符0与进程的标准输入关联，1与标准输出关联，2与标准错误关联  
通常 把0，1，2替换成常量 STDIN_FILENO、STDOUT_FILENO、STDERR_FILENO


---

## 3.3 函数open和openat  
```
#include <fcntl.h>
int open(const char *path, int oflag, ... /* mode_t mode */);
int openat(int fd, const *path, int oflag, ... /* mode_t mode */);
// 最后一个参数 ... ISO C用这种方法表明余下的参数的数量及其类型是可变的。
//open 函数 只有当创建新文件时才使用最后这个参数 
//函数的返回值：若成功，返回文件描述符，若出错，返回-1
```
path 参数是要打开或创建文件的名字。oflag参数 用来说明该函数的多个选项。  
有以下一个或多个常亮进行 “或”运算构成oflag参数 （这些常量定义在<fcntl.h>）
下面这些常量是必须的
- O_RDONLY Open for reading only.
- O_WRONLY Open for writing only.
- O_RDWR   Open for reading and writing.
- O_EXEC   Open for execute only.
- O_SEARCH Open for search only(applies to directories).

这5个常量中必须指定一个且只能指定一个。还有一些常量是可选的(慢慢补充)
* O_APPEND   每次写都追加到文件末尾

## 3.6 函数lseek
```
#include <unistd.h>
off_t lseek(int fd, off_t offset, int whence);
//返回值：若成功，返回新的文件的偏移量，若出错，返回-1
```

whence 参数
* whence是 SEEK_SET,将文件偏移量设置为据文件开始处offset个字节
* whence是 SEEK_CUR,将文件偏移量设置为当前值加上offset,offset可为正或负
* whence是 SEEK_END,将文件偏移量设置为文件长度加上offset,offset可为正或负
在比较lseek的返回值时要谨慎，不要测试它是否小于0，而应该测试是否等于-1  

```
#include "apue.h"
int main(void)
{
    if (lseek(STDIN_FILENO, 0, SEEK_CUR) == -1)
        printf("cannot seek\n");
    else
        printf("seek ok\n");
    
    exit(0);
}
```

## 3.7 函数read
```
//read 函数从打开的文件中读取数据
#include <unistd.h>
ssize_t read(int fd,void *buf,size_t nbytes);
//返回值：读到的字节数，如果到达文件末尾返回0，若出错返回-1
```

## 3.8 函数write

调用write函数向打开的文件写数据
```
#include <unistd.h>
ssize_t write(int fd, const void *buf, size_t nbytes);
//成功 返回已写入的字节数,失败返回-1
```
write 出错的一个常见原因是磁盘已写满 或者超过了一个给定进程的文件长度限制  
普通文件 写操作从当前文件偏移量开始。打开文件时指定了O_APPEND选项(O_APPEND每次写都追加到文件末尾),  
每次写操作将文件偏移量设置在文件的当前结尾处,在一次写成功后，该文件偏移量增加实际写的字节数
