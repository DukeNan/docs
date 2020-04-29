# gevent使用指导

> 文章转载于：[文章链接](https://sdiehl.github.io/gevent-tutorial/)



## 核心部分

### Greenlets

gevent中使用的主要模式是**Greenlet**，这是作为C扩展模块提供给Python的轻量协程。Greenlets都在主程序的OS进程内运行，但是是协同安排的。

> 在任何给定时间，只有一个greenlet正在运行。

这与由`multiprocessing`或`threading`库提供的任何真正的并行性构造不同， 它们执行旋转进程和POSIX线程，这些进程和操作系统由操作系统调度并且是真正并行的。

### 同步与异步执行

