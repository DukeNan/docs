## CELERY 的多任务 多队列

celery是一个分布式的任务调度模块，那么怎么实现它的分布式功能呢，celery可以支持多台不同的计算机执行不同的任务或者相同的任务。

如果要说celery的分布式应用的话，就要提到celery的消息路由机制，提到AMQP协议。

简单理解：

可以有多个"消息队列"（message Queue），不同的消息可以指定发送给不同的Message Queue，

而这是通过Exchange来实现的，发送消息到"消息队列"中时，可以指定routiing_key，Exchange通过routing_key来吧消息路由（routes）到不同的"消息队列"中去。

![celery](../_media/images/python/celery_01.png)

exchange 对应 一个消息队列(queue)，即：通过"消息路由"的机制使exchange对应queue，每个queue对应每个worker。

下面我们来看一个列子：

```python
vi tasks.py
 
#!/usr/bin/env python
#-*- coding:utf-8 -*-
from celery import Celery
 
app = Celery()
app.config_from_object("celeryconfig")  # 指定配置文件
 
@app.task
def taskA(x,y):
return x + y
 
@app.task
def taskB(x,y,z):
return x + y + z
 
@app.task
def add(x,y):
return x + y
```

编写配置文件，配置文件一般单独写在一个文件中。

```python
vi celeryconfig.py
 
#!/usr/bin/env python
#-*- coding:utf-8 -*-
 
from kombu import Exchange,Queue
 
BROKER_URL = "redis://:your_password@127.0.0.1:6379/1" 
CELERY_RESULT_BACKEND = "redis://:your_password@127.0.0.1:6379/1"
 
CELERY_QUEUES = (
Queue("default",Exchange("default"),routing_key="default"),
Queue("for_task_A",Exchange("for_task_A"),routing_key="for_task_A"),
Queue("for_task_B",Exchange("for_task_B"),routing_key="for_task_B") 
)
# 路由
CELERY_ROUTES = {
'tasks.taskA':{"queue":"for_task_A","routing_key":"for_task_A"},
'tasks.taskB':{"queue":"for_task_B","routing_key":"for_task_B"}
}
CELERY_DEFAULT_QUEUE = 'default'   # 设置默认的路由
CELERY_DEFAULT_EXCHANGE = 'default'
CELERY_DEFAULT_ROUTING_KEY = 'default'
 
CELERY_TASK_RESULT_EXPIRES = 10  # 设置存储的过期时间　防止占用内存过多
```

远程客户端上编写测试脚本

```python
vi test.py
 
from tasks import *
re1 = taskA.delay(100, 200)
print(re1.result)
re2 = taskB.delay(1, 2, 3)
print(re2.result)
re3 = add.delay(1, 2)
print(re3.status)
```

启动两个worker来分别指定taskA、taskB，开两个窗口分别执行下面语句。

```bash
celery -A tasks worker -l info -n workerA.%h -Q for_task_A
 
celery -A tasks worker -l info -n workerB.%h -Q for_task_B
```

远程客户端上执行脚本可以看到如下输出：

```bash
python test.py 
300
6
PENDING
```

在taskA所在窗口可以看到如下输出

```bash
.......
.......
.......
task_A
 
[tasks]
  . tasks.add
  . tasks.taskA
  . tasks.taskB
 
[2018-05-27 19:23:49,235: INFO/MainProcess] Connected to redis://47.106.106.220:5000/1
[2018-05-27 19:23:49,253: INFO/MainProcess] mingle: searching for neighbors
[2018-05-27 19:23:50,293: INFO/MainProcess] mingle: all alone
[2018-05-27 19:23:50,339: INFO/MainProcess] celery@workerA.izwz920j4zsv1q15yhii1qz ready.
[2018-05-27 19:23:56,051: INFO/MainProcess] sync with celery@workerB.izwz920j4zsv1q15yhii1qz
[2018-05-27 19:24:28,855: INFO/MainProcess] Received task: tasks.taskA[8860e78a-b82b-4715-980c-ae125dcab2f9]  
[2018-05-27 19:24:28,872: INFO/ForkPoolWorker-1] Task tasks.taskA[8860e78a-b82b-4715-980c-ae125dcab2f9] succeeded in 0.0162177120219s: 300
```

在taskB所在窗口可以看到如下输出：

```bash
.......
.......
.......
task_B
[tasks]
  . tasks.add
  . tasks.taskA
  . tasks.taskB
 
[2018-05-27 19:23:56,012: INFO/MainProcess] Connected to redis://47.106.106.220:5000/1
[2018-05-27 19:23:56,022: INFO/MainProcess] mingle: searching for neighbors
[2018-05-27 19:23:57,064: INFO/MainProcess] mingle: sync with 1 nodes
[2018-05-27 19:23:57,064: INFO/MainProcess] mingle: sync complete
[2018-05-27 19:23:57,112: INFO/MainProcess] celery@workerB.izwz920j4zsv1q15yhii1qz ready.
[2018-05-27 19:24:33,885: INFO/MainProcess] Received task: tasks.taskB[5646d0b7-3dd5-4b7f-8994-252c5ef03973]  
[2018-05-27 19:24:33,910: INFO/ForkPoolWorker-1] Task tasks.taskB[5646d0b7-3dd5-4b7f-8994-252c5ef03973] succeeded in 0.0235358460341s: 6
```

我们看到状态是PENDING，表示没有执行，这个是因为没有celeryconfig.py文件中指定改route到哪一个Queue中，所以会被发动到默认的名字celery的Queue中，但是我们还没有启动worker执行celery中的任务。下面，我们来启动一个worker来执行celery队列中的任务。

```bash
celery -A tasks worker -l info -n worker.%h -Q celery
```

再次在远程客户端执行test.py，可以看到结果执行成功，并且刚新启动的worker窗口有如下输出：

```bash
.......
.......
.......
[tasks]
  . tasks.add
  . tasks.taskA
  . tasks.taskB
 
[2018-05-27 19:25:44,596: INFO/MainProcess] Connected to redis://47.106.106.220:5000/1
[2018-05-27 19:25:44,611: INFO/MainProcess] mingle: searching for neighbors
[2018-05-27 19:25:45,660: INFO/MainProcess] mingle: sync with 2 nodes
[2018-05-27 19:25:45,660: INFO/MainProcess] mingle: sync complete
[2018-05-27 19:25:45,711: INFO/MainProcess] celery@worker.izwz920j4zsv1q15yhii1qz ready.
[2018-05-27 19:25:45,868: INFO/MainProcess] Received task: tasks.add[f9c5ca2b-623e-4c0a-9c45-a99fb0b79ed5]  
[2018-05-27 19:25:45,880: INFO/ForkPoolWorker-1] Task tasks.add[f9c5ca2b-623e-4c0a-9c45-a99fb0b79ed5] succeeded in 0.0107084610499s: 3
```

