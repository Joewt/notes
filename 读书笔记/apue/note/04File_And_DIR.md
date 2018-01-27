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
