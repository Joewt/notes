# 文件和目录  

本章主要讲解了文件系统的其他特征和文件的性质。


## 4.2 函数stat，fstat，fstatat和lstat
```
#include <sys/stat.h>
int stat(const char * restrict pathname,struct stat *restrict buf);
int fstat(int fd,struct stat *buf);
int lstat(const char *restrict pathname,struct stat *restrict buf);
int fstatat(int fd,const char *restrict pathname,struct stat *restrict buf,int flag);
```
成功返回0，失败返回-1
* stat 函数返回文件有关的信息结构
* fstat 函数获得已在描述符fd上打开文件的有关信息
* lstat 函数类似于stat，当命名的文件是一个符号链接时，lstat返回该符号链接的有关信息，而不是该符号链接引用的文件的信息
* fstatat 函数为一个相对于当前打开目录（由fd参数指定）的路径名返回文件统计信息。flag参数控制着是否跟随着一个符号链接。

buf参数是一个指针，指向一个我们必须提供的结构 在mac系统上该结构的定义如下
```

struct stat {
     dev_t       st_dev;     /* [XSI] ID of device containing file */
     ino_t       st_ino;     /* [XSI] File serial number */
     mode_t      st_mode;    /* [XSI] Mode of file (see below) */
     nlink_t     st_nlink;   /* [XSI] Number of hard links */
     uid_t       st_uid;     /* [XSI] User ID of the file */
     gid_t       st_gid;     /* [XSI] Group ID of the file */
     dev_t       st_rdev;    /* [XSI] Device ID */
 #if !defined(_POSIX_C_SOURCE) || defined(_DARWIN_C_SOURCE)
     struct  timespec st_atimespec;  /* time of last access */
     struct  timespec st_mtimespec;  /* time of last data modification */
     struct  timespec st_ctimespec;  /* time of last status change */
 #else
     time_t      st_atime;   /* [XSI] Time of last access */
     long        st_atimensec;   /* nsec of last access */
     time_t      st_mtime;   /* [XSI] Last data modification time */
     long        st_mtimensec;   /* last data modification nsec */
     time_t      st_ctime;   /* [XSI] Time of last status change */
     long        st_ctimensec;   /* nsec of last status change */
 #endif
     off_t       st_size;    /* [XSI] file size, in bytes */
     blkcnt_t    st_blocks;  /* [XSI] blocks allocated for file */
     blksize_t   st_blksize; /* [XSI] optimal blocksize for I/O */
     __uint32_t  st_flags;   /* user defined flags for file */
     __uint32_t  st_gen;     /* file generation number */
     __int32_t   st_lspare;  /* RESERVED: DO NOT USE! */
     __int64_t   st_qspare[2];   /* RESERVED: DO NOT USE! */
 };
```


文件类型  

|宏|文件类型|
|:--|:--|
|S_ISREG() |普通文件|
|S_ISDIR() |目录文件|
|S_ISCHR() |字符特殊文件|
|S_ISBLK() |块特殊文件|
|S_ISFIFO()|普通或FIFO|
|S_ISLNK() |符号链接|
|S_ISSOCK()|套接字|

取命令行参数，然后根据每个命令行参数打印文件类型
```
/*
 对每个命令行参数打印文件类型
*/
#include "apue.h"
int main(int argc, char *argv[])
{
    int i;
    struct stat buf;
    char *ptr;
    for(i = 1; i < argc; i++) {
        printf("%s: ",argv[i]);
        if (lstat(argv[i],&buf) < 0) {
            err_sys("lstat error");
            continue;
        }
        printf("%hu\n",buf.st_mode);
        printf("%d\n",S_ISREG(buf.st_mode));
        if (S_ISREG(buf.st_mode))
            ptr = "regular";
        else if (S_ISDIR(buf.st_mode))
            ptr = "directory";
        else if (S_ISCHR(buf.st_mode))
            ptr = "character special";
        else if (S_ISBLK(buf.st_mode))
            ptr = "block special";
        else if (S_ISFIFO(buf.st_mode))
            ptr = "fifo";
        else if (S_ISLNK(buf.st_mode))
            ptr = "symbolic link";
        else if (S_ISSOCK(buf.st_mode))
            ptr = "socket";
        else
            ptr = "unknown mode";
        printf("%s\n",ptr);
}
exit(0);
}
//这些宏的定义如下
/*
#define S_ISBLK(m)  (((m) & S_IFMT) == S_IFBLK)
#define S_ISCHR(m)  (((m) & S_IFMT) == S_IFCHR)
#define S_ISDIR(m)  (((m) & S_IFMT) == S_IFDIR)
#define S_ISFIFO(m) (((m) & S_IFMT) == S_IFIFO)
#define S_ISREG(m)  (((m) & S_IFMT) == S_IFREG)
#define S_ISLNK(m)  (((m) & S_IFMT) == S_IFLNK)
#define S_ISSOCK(m) (((m) & S_IFMT) == S_IFSOCK)
#if !defined(_POSIX_C_SOURCE) || defined(_DARWIN_C_SOURCE)
#define S_ISWHT(m)  (((m) & S_IFMT) == S_IFWHT)
#endif
*/
```

## 4.5 文件访问权限
除了文件外，目录字符特殊文件都有访问权限  
每个文件有9个访问权限位，分为如下3类  
在 <sys/stat.h>中定义  

|st_mode屏蔽|含义|
|:--|:--|
| S_IRUSR  |用户读   |
|S_IWUSR   |用户写   |
|S_IXUSR   |用户执行   |

|st_mode屏蔽|含义|
|:--   |:--   |
|S_IRGRP   |组读   |
|S_IWGRP   |组写   |
|S_IXGRP   |组执行   |

|st_mode屏蔽|含义|
|:--|:--|
|S_IROTH   |其他读   |
|S_IWOTH   |其他写   |
|S_IXOTH   |其他执行   |
* 当用名字打开一任意类型的文件时，对改名字包含的每一个目录，包括它可能隐含的当前工作目录都应具有执行权限
    - 对目录的读权限，可以列出目录的文件名列表
* 对于一个文件的读权限决定了我们是否能够打开现有文件进行读操作。这与open函数的O_RDONLY和O_RDWR标志决定
* 对于一个文件的写权限决定了我们是否能够打开现有文件进行写操作。这与open函数的O_WRONLY和O_RDWR标志决定
* 为了在open函数中对一个文件指定O_TRUNC标志，必须对该文件具有写权限
* 在一个目录中创建一个新文件，需对该目录有写权限和执行权限
* 为了删除一个现有的文件，必须对包含该文件的目录具有写和执行权限，对该文件本身不需要有读、写权限
*

## 4.7 函数access和faccessat
当用open函数打开文件时，内核以进程的有效用户id和有效组id为基础执行其访问文件权限测试。但有时，进程希望按照其实际用户id和实际组id来测试访问能力。access函数和faccessat函数式按照实际id进行测试
```
#include <unistd.h>
int access(const char* pathname,int mode);
int faccessat(int fd,const char* pathaname,int mode,int flag);
//成功返回0，出错返回-1
```
如果测试文件存在，mode就为F_OK,否则mode为如下图中所列常量的按位或。

|mode|说明|
|:--|:--|
| R_OK  |测试读权限  |
| W_OK  |测试写权限  |
| X_Ok  |测试执行权限|
代码示例
```
#include "apue.h"
#include <fcntl.h>
int main(int argc, char *argv[])
{
    if (argc != 2)
        err_quit("参数不够");

    if (access(argv[1],F_OK) < 0)
        err_ret("权限错误: %s",argv[1]);
    else
        printf("读取权限 OK\n");

    if (open(argv[1],O_RDONLY) < 0)
        err_ret("打开文件错误: %s",argv[1]);
    else
        printf("打开文件 OK\n");

    exit(0);
}
```
## 4.8 函数umask
umask为进程设置文件模式创建屏蔽字,并返回之前的值(没有出错返回)
```
#include <sys/stat.h>
mode_t umask(mode_t cmask);
```
cmask 参数有9个访问权限位若干个按位"或"构成  
常见的几种umask值
* 002 阻止其他用户写入你的文件
* 022 阻止同组成员和其他用户写入你的文件
* 027 阻止同组成员写你的文件以及其他用户读、写、执行你的文件

## 4.12 文件长度
stat 结构成员st_size 表示以字节为单位的文件长度,改字段只对普通文件，目录文件，链接文件有作用  
* 普通文件，其文件长度可以为0，从文件开始到end-of-file。
* 目录文件，是一个数的整数倍(例如16或512)
* 链接文件，文件名中的实际字节数

## 4.13 文件截断

```
#include <unistd.h>
int truncate(const char *pathname,off_t length);
int ftruncate(int fd, off_t length);
//成功返回0，失败返回-1

```

如果文件长度大于length，则length后的数据不能访问，如果小于，文件长度增加。增加的这段填充0(可能在文件中创建了一个空洞)

## 4.14 文件系统



## 4.15 函数link,linkat,unlink,unlinkat,remove
任何一个文件可以有多个目录项指向其i节点。创建一个现有文件的链接的方法是使用link或linkat函数  
```
#include <unistd.h>
int link(const char* existingpath, const char* newpath);
int linkat(int efd, const char* existingpath, int nfd, const char* newpath, int flag);
//成功返回0，出错返回-1

```
* 这俩函数创建一个新目录项newpath，引用现有文件existingpath。如果newpath已经存在则出错。只创建newpath中最后一个分量。  
* linkat函数，现有文件和新的路径名都是通过efd和existingpath、nfd和newpath指定。nfd和efd是文件描述符，如果目录是相对目录就根据文件描述符进行计算，如果是绝对路径该参数将会忽略。  
* 当现有文件是符号链接时，由flag参数来控制linkat函数是创建现有符号链接的链接还是创建执行现有符号链接指向的文件的链接。如果在flag参数中设置了AT_SYMLINK_FOLLOW标志，就创建指向符号链接目标的链接。如果这个标志被清除则创建一个指向符号链接本身的链接。  

删除一个目录项，可以调用unlink函数
```
#incude <unistd.h>
int unlink(const char* pathname);
int unlinkat(int fd, const char* pathname, int flag);
//成功返回0，出错返回-1
```
删除目录项，并将由pathname所引用文件链接计数减1。只有当链接计数达到0时，该文件才被删除。  
为了解除对文件的链接，必须对包含该目录项的目录具有写和执行权限。如果该目录设置了粘着位，则对该目录必须具有写权限，并且要具备下面三个条件之一：
* 拥有该文件
* 拥有该目录
* 具有超级用户权限

unlink通常用于确保即使程序崩溃，它所创建的临时文件也不会遗留下来。进程用open和creat创建一个文件，然后立即调用unlink，因为该文件仍然是打开的，所以不会将其内容删除。只有当进程关闭该文件或终止时，该文件的内容才被删除。  
#### remove函数
```
#include <stdio.h>
int remove(const char* pathname);
//成功返回0，失败返回-1
```
对于文件remove与unlink的功能类似，对于目录与rmdir相同  

## 4.16 函数rename与renameat
文件或目录可以使用rename和renameat重命名  
```
#include <stdio.h>
int rename(const char* oldname, const char* newname);
int renameat(int oldfd, const char* oldname, int newfd, const char* newname);
//成功返回0，失败返回-1
```
oldname参数是 目录、文件还是符号链接有以下几点说明
1. 如果oldname指的是一个文件而不是目录，那么为该文件或符号链接重命名。在这种情况下如果newname已经存在，则它不能引用一个目录。如果newname已经存在并且不是一个目录，则先将该目录项删除然后将oldname重命名为newname。对包含oldname的目录以及newname的目录，调用进程必须有写权限，因为要更改该目录
2. 如果oldname指的是一个目录，那么为该目录重命名。如果newname已经存在，则它必须引用一个目录，而且该目录是空目录(空目录指只包含.和..项)。如果newname存在，则先将其删除，然后将oldname重命名为newname。另外，当为一个目录重命名时，newname不能包含oldname作为其路径前缀。例如不能将/usr/foo重命名为/usr/foo/test,因为旧名字是新名字的路径前缀，因而不能讲其删除。
3. 如果oldname和newname引用的是符号链接，则处理的是符号链接本身，而不是符号链接引用的文件  
4. 不能对.和..重命名
5. 如果oldname和newname引用同一个文件，则不做任何改动 返回

## 4.17 符号链接(软链接)
符号链接是对一个文件的间接指针，引用符号链接是避免硬链接的一些限制
* 硬链接通常要求链接和文件在同一个文件系统中
* 只有超级用户才能创建指向目录的硬链接(在底层文件系统的支持下)

符号链接一般用于将一个文件或整个目录结构移到系统的另一个位置。符号链接有着自己的inode号和用户数据块
## 4.19 文件的时间
对每个文件通常维护3个时间段  
|字段|说明|例子|ls选项|
|--|--|--|--|
|st_atim   |文件数据的最后访问时间   |read   |-u   |
|st_mtim   |文件数据的最后修改时间   |write  |默认  |
|st_ctim   |i节点状态的最后更改时间   |chmod,chown   |-c   |
