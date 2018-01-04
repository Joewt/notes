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

## 函数lseek
```
#include <unistd.h>
off_t lseek(int fd, off_t offset, int whence);
//返回值：若成功，返回新的文件的偏移量，若出错，返回-1
```