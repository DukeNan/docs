## 事件循环

### 概念

在异步编程中事件的概念非常重要，因为这是异步编程的本质和根本所在。在计算系统中，可以生成事件的实体通常被称为**事件源**，而管理事件的实体称为**事件处理程序**。asyncio中有第三个实体称为**事件循环**`event loop`，它实现了管理事件的所有功能，事件循环在整个程序执行过程中循环运行，并跟踪进程中发生的所有事件，当主线程空闲时通过调用事件处理程序对它们进行排队和处理。当处理程序结束时，控件将传递到被调度的下一个事件。

### 方法：

Asyncio提供用于管理事件循环的方法如下：

- loop = get_event*_*loop()方法 - 获取当前上下文的事件循环
- loop.call_later(time_delay, callback, argument) - 在给定的时间延迟秒之后调用回调函数
- loop.call_soon(callback, argument) - 当控制返回到事件循环时调用
- loop.time() - 根据事件循环的内部时钟将当前时间返回
- asyncio.set_event_loop() - 设置当前上下文的事件循环为循环
- asyncio.new_event_loop() - 根据策略的规则创建并返回一个新的事件循环对象
- loop.run_forever() - 永远运行直到调用stop()

### 实例

下面介绍如何使用Asyncio库提供的循环事件语句来构建异步模式工作的应用程序。在代码中，我们定义了三个异步任务，每个任务按序号调用后续任务，具体如下：

```python
""" Asyncio loop """

import asyncio
import datetime
import time

def function_one(end_time, loop):
    print("function 1 called")
    print(end_time)
    if (loop.time() + 1.0) < end_time:
        loop.call_later(1, function_two, end_time, loop)
    else:
        loop.stop()

def function_two(end_time, loop):
    print("function 2 called")
    print(end_time)
    if (loop.time() + 1.0) < end_time:
        loop.call_later(1, function_three, end_time, loop)
    else:
        loop.stop()

def function_three(end_time, loop):
    print("function 3 called")
    print(end_time)
    if (loop.time() + 1.0) < end_time:
        loop.call_later(1, function_one, end_time, loop)
    else:
        loop.stop()

loop = asyncio.get_event_loop()

# Schedule the first call to display_date()
end_loop_time = loop.time() + 9.0
loop.call_soon(function_one, end_loop_time, loop)

# Blocking call interrupted by loop.stop()
loop.run_forever()
loop.close()
```

我们通过调用`call_soon`安排函数`function_one`的第一次调用，并通过end_time定义了函数的执行时间上限，然后通过`call_later`方法调用`function_two`，再之后是第三个函数，函数的任务相当简单，打印出函数名称和结束的上限时间。输出结果如下：

```python
$ python asyncio_loop.py
function 1 called
277782.369758452
function 2 called
277782.369758452
function 3 called
277782.369758452
function 1 called
277782.369758452
function 2 called
277782.369758452
function 3 called
277782.369758452
function 1 called
277782.369758452
function 2 called
277782.369758452
function 3 called
277782.369758452

```

