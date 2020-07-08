## netstat显示网络状态

### 概念

netstat 命令用于显示各种网络相关信息，如网络连接，路由表，接口状态 (Interface Statistics)，masquerade 连接，多播成员 (Multicast Memberships) 等等。

从整体上看，netstat的输出结果可以分为两个部分：一个是Active Internet connections，称为有源TCP连接，其中”Recv-Q”和”Send-Q”指%0A的是接收队列和发送队列。这些数字一般都应该是0。如果不是则表示软件包正在队列中堆积。这种情况只能在非常少的情况见到；另一个是Active UNIX domain sockets，称为有源Unix域套接口(和网络套接字一样，但是只能用于本机通信，性能可以提高一倍)。

### 语法

> netstat [参数]

#### 常见参数

|  参数  |   解释   |
| :--- | :--- |
| -a (all) | 显示所有选项，默认不显示LISTEN相关 |
| -t (tcp) | 仅显示tcp相关选项 |
| -u (udp) | 仅显示udp相关选项 |
| -l | 仅列出有在 Listen (监听) 的服务状态 |
| -p (programs) | 显示建立相关链接的程序名 |
| -r | 显示路由信息，路由表 |
| -e | 显示扩展信息，例如uid等 |
| -s | 按各个协议进行统计 |
| -c | 每隔一个固定时间，执行该netstat命令 |
| -n (numeric) | 直接使用IP地址，不通过域名服务器，不解析名称 |

**注意：**LISTEN和LISTENING的状态只有用-a或者-l才能看到。

### 常用命令

1. 查看tcp端口状态
```shell
[root@aliyun]# netstat -lnp | grep :80
tcp        0      0 0.0.0.0:80              0.0.0.0:*               LISTEN      2482/nginx: worker  
```

2. 查看路由表

```shell
root@shaun-pc:~# netstat -i
Kernel Interface table
Iface      MTU    RX-OK RX-ERR RX-DRP RX-OVR    TX-OK TX-ERR TX-DRP TX-OVR Flg
docker0   1500        0      0      0 0            12      0      0      0 BMU
enp1s0    1500   357382      0      0 0        123384      0      0      0 BMRU
lo       65536    22736      0      0 0         22736      0      0      0 LRU
wlp2s0    1500     6165      0      0 0           228      0      0      0 BMU
```




