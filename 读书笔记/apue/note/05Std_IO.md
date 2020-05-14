# 标准I/O库
主要讲解了标准I/O库的一些内容  

---

## 5.2 流和FILE对象
标准I/O库，他们的操作都是围绕流进行的  
字符分为单字节和多字节，流的定向决定了是所读、写的字符是单字节还是多字节的。当一个流被创建时，并没有定向。使用多字节I/O函数，流的定向就被设置为宽定向，单字节同理。只有两个函数可以改变流的定向，freopen和fwide  

```
#include <stdio.h>
#include <wchar.h>
int fwide(FILE *fp, int mode);
//若流是宽定向的，返回正值；若流是单字节定向的返回负值，流未定向，返回0
```
mode参数的说明  
* mode为负，fwide将使指定的流是单字节定向的
* mode为正，fwide将使指定的流是宽定向的
* mode为0，不设置流的定向，但返回标识该流定向的值

---

## 5.4 缓冲
标准I/O库提供缓冲的目的是尽可能减少read和write的调用。  
标准I/O提供了3种缓冲
1. 全缓冲。只有填满标准I/O缓冲区后才进行实际的I/O操作。对于驻留在磁盘上的文件通常是由标准I/O库实施全缓冲的。在一个流上执行第一次I/O操作是，相关I/O函数通常调用malloc获得需要的缓冲。
2. 行缓冲。当在输入输出中遇到换行符时，标准I/O库执行I/O操作。这允许我们一次输出一个字符(fputc),但只有在写了一行之后才进行实际的I/O操作。当涉及一个终端时(如标准输入和标准输出)，通常使用行缓冲。
3. 不带缓冲。标准I/O库不对字符进行缓冲存储。例如，若用标准I/O函数fputs写15个字符到不带缓冲的流中，我们就期望这15个字符能立即输出，就很有可能使用write函数将这些字符写到相关联的打开文件中。

标准错误流stderr通常是不带缓冲的。  
ISO C要求下列缓冲特征
* 当且仅当标准输入和标准输出并不指向交互设备时，他们才是全缓冲的
* 标准错误绝不是全缓冲的

很多系统默认使用下列类型的缓冲
* 标准错误是不带缓冲的
* 若是指向终端设备的流，则是行缓冲的；否则是全缓冲的

如果不喜欢系统默认的可使用下列两个函数中的一个更改缓冲类型

```
#include <stdiio.h>
void setbuf(FILE *restrict fp, char *restrict buf);
int setvbuf(FILE *restrict fp, char *restrict buf, int mode, size_t size);
```

在任何时候都可以强制冲洗一个流
```
#include <stdio.h>
int fflush(FILE *fp);
//成功返回0，失败返回EOF
```

---

## 5.5 打开流

下列3个函数打开一个标准I/O流
```
#include <stdio.h>
FILE *fopen(const char *restrict pathname, const char *restrict type);
FILE *freopen(const char *restrict pathname, const char *restrict type, FILE *restrict fp);
FILE *fdopen(int fd, const char *type);
//成功返回文件指针，失败返回NULL
```
这3个函数区别如下
1. fopen函数打开路径名为pathname的一个指定的文件
2. freopen函数在一个指定的流上打开一个指定的文件，如若该流已经打开，则先关闭该流。若流已经定向，则使用freopen清除定向。此函数一般用于将一个指定的文件打开为一个预定义的流：标准输入、标准输出或标准错误
3. fdopen函数取一个已有的文件描述符(我们可能从open、dup、dup2、fcntl、pipe、socket、socketpair或accept函数得到此文件描述符)，并使一个标准I/O流与该描述符相结合。此函数常用于由创建管道和网络通信通道函数返回的描述符。因为这些特殊类型文件不能使用标准I/O函数fopen打开，所以我们必须先调用设备专用函数以获得一个文件描述符，然后用fdopen使一个标准I/O流与该描述符相结合。

fdopn是 POSIX.1 具有的

type参数说明如下
|type|说明|open(2)标志|
|:--|:--|:--|
|r或rb   |为读而打开   |O_RDONLY|
|w或wb   |把文件截断至0长，或为写而创建   |O_WRONLY\|O_CREAT\|O_TRUNC|
|a或ab   |追加；为在文件尾而写打开，或为写而创建   |O_WRONLY\|O_CREAT\|O_APPEND|
|r+或r+b或rb+   |为读和写打开   |O_RDWR|
|w+或w+b或w+   |把文件截断为0长，或为读和写而打开   |O_RDWR\|O_CREAT\|O_TRUNC|
|a+或a+b或ab+   |为在文件尾读和写而打开或创建   |O_RDWR\|O_CREAT\|O_APPEND|

---

## 读和写流
一旦打开了流，则可在3种不同类型的非格式化I/O中进行选择，对其进行读、写操作  
1. 每次一个字符的I/O。一次读或写一个字符，如果流是带缓冲的，则标准I/O函数处理所有缓冲。
2. 每次一行的I/O。如果想要一次读或写一行，则使用fgets和fputs。每行都以一个换行符终止。当调用fgets时，应说明能处理的最大行长。
3. 直接I/O。fread和fwrite函数支持这种类型的I/O。每次I/O操作读或写某种数量对象，而每个对象具有指定的长度。这两个函数常用于从二进制文件中每次读或写一个结构。

#### 1. 输入函数
以下三个函数可用于一次读一个字符  
```
#include <stdio.h>
int getc(FILE *fp);
int fgetc(FILE *fp);
int getchar(void);
//成功返回下一个字符，失败或到达文件末尾，返回EOF
```
函数getchar()等同于 getc(stdin)。getc可被实现为宏，fgetc不能。意味着以下几点
1. getc的参数不应当是具有副作用的表达式，因此它可能会被计算多次
2. 因为fgetc一定是一个函数，所以它可以得到地址。这就允许将fgetc的地址作为一个参数传递给另一个函数
3. 调用fgetc所需时间很可能比调用getc要长，因为调用函数所需时间通常长于调用宏

不管是出错还是到达文件尾端，这3个函数都返回同样的值。为了区分这两种不同的情况，必须调用ferror或feof。
```
#include <stdio.h>
int ferror(FILE *fp);
int feof(FILE *fp);
//两个函数的返回值，若成功返回非0(真)；否则返回0(假)
void clearerr(FILE *fp);
```
在大多数实现中，为每个流在FILE对象中维护了两个标志:
* 出错标志
* 文件结束标志

调用clearerr可以清除这两个标志。  
从流中读取数据以后，可以调用ungetc将字符再压送回流中。
```
#include <stdiio.h>
int ungetc(int c,FILE *fp);
//成功返回c，出错返回EOF
```

#### 2. 输出函数
对应上面每个输入函数都有一个输出函数
```
#include <stdio.h>
int putc(int c, FILE *fp);
int fputc(int c, FILE *fp);
int putchar(int c);
//成功返回c，出错返回EOF
```
---

## 5.7 每次一行I/O
下面两个函数提供每次输入一行的功能
```
#include <stdio.h>
char *fgets(char *restrict buf, int n, FILE *restrict fp);
char *gets(char *buf);
//成功返回buf；若已到达文件尾端或出错，返回NULL
```
推荐使用fgets  
每个输入都有对应的输出
```
#include <stdio.h>
int fputs(const char* restrict str, FILE *restrict fp);
int puts(const char *str);
//成功返回非负值，失败返回EOF
```

---

## 5.9 二进制I/O
系统提供了一下函数执行二进制I/O操作
```
#include <stdio.h>
size_t fread(void *restrict ptr, size_t size, size_t nobj, FILE *restrict fp);
size_t fwrite(const void *restrict ptr, size_t size, size_t nobj, FILE *restrict fp);
//两个函数的返回值：读或写的对象数
```
---

## 5.10 定位流
有3种方式定位标准I/O流  
1. ftell和fseek函数。这俩函数都是假定文件的位置可以存放在一个长整数中
2. ftellofseeko函数。Single UNIX Specification引入了这两个函数，使文件偏移量可以不必一定使用长整形。他们使用off_t数据类型替代了长整形
3. fgetpos和fsetpos函数。这两个函数是由ISO C引入的。他们使用一个抽象数据类型fpos_t记录文件位置。这种数据类型可以根据需要定义一个足够大的数，用以记录文件位置。

需要移植到非UNIX系统上需要使用fgetpos和fsetpos函数


---

## 5.11格式化I/O
格式化输出是由5个printf函数处理
```
#include <stdio.h>
int printf(const char* restrict format, ...);
int fprintf(FIFE *restrict fp, const char* restrict format, ...);
int dprintf(int fd, const char* restrict format, ...);
//若成功返回输出字符数。出错返回负值
int sprintf(char* restrict buf, const char* restrict format, ...);
//成功返回存入数组的字符数，若编码出错返回负值
int snprintf(char* restrict buf, size_t n, const char* restrict format, ...);
若缓冲区足够大，返回将要存入数组的字符数，若编码出错返回负值
```
printf将格式化的数据写到标准输出，fprintf写至指定的流，dprintf写至指定的文件描述符，sprintf将格式化的字符送入数组buf中。sprintf在该数组的尾端自动加一个null字节，但该字符不包括在返回值中。  


#### 格式化输入
```
#incude <stdio.h>
int scanf(const char* restrict format, ...);
int fscanf(FIFE *restrict fp, const char* restrict format, ...);
int sscanf(const char* restrict buf, const char* restrict format, ...);
//3个函数的返回值，赋值的输入项数，若出错或任一转换前已到达文件尾端，返回EOF
```



## 5.13 临时文件
ISO C标准I/O库提供了两个函数以帮助创建临时文件
```
#include <stdiio.h>
char *tmpnam(char *ptr);
FILE *tmpfile(void);
//成功返回文件指针，出错返回null
```
tmpnam函数产生一个与现有文件名不同的一个有效路径名字符串。每次调用它时，都产生一个不同的路径名，最多调用次数是TMP_MAX(定义在<stdio.h>在mac系统该值为308915776)  
若ptr是NULL，则所产生的路径名存放在一个静态区中，指向该静态区的指针作为函数值返回。后续调用tmpnam时，会重写该静态区(这意味着，如果我们调用此函数多次，而且想保存路径名，则我们应该保存该文件的副本，而不是指针的副本)。如若ptr不是NULL，则认为它应该是指向长度至少是L_tmpnam个字符的数组(L_tmpnam定义在<stdio.h>中mac下为1024)。所产生的路径名存放在该数组中，ptr也作为函数值返回。  
tmpfile创建一个临时二进制文件(类型wb+)，在关闭该文件或程序结束时将自动删除这种文件。注意，UNIX对二进制文件不进行特殊区分。  
