### apue 配置问题

使用的是centos7 系统 编译源码的时候发现如下错误
```
gcc -ansi -I../include -Wall -DLINUX -D_GNU_SOURCE  barrier.c -o barrier  -L../lib -lapue -pthread -lrt -lbsd
/tmp/ccF6O2ef.o: In function `thr_fn':
barrier.c:(.text+0x80): undefined reference to `heapsort'
collect2: error: ld returned 1 exit status
make[1]: *** [barrier] Error 1
make[1]: Leaving directory `/home/learnApue/apue.3e/threads'
make: *** [all] Error 1
```
解决办法 [这里](https://pkgs.org/) 下载两个包  
下载的为libbsd-devel-0.6.0-3.el7.elrepo.x86_64.rpm和libbsd-devel-0.6.0-3.el7.elrepo.x86_64.rpm
按照顺序按照   
安装方法 rpm -Uvh xxxx.rpm  

这样就可以使用了