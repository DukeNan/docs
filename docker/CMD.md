## Docker 中 RUN、CMD 与 ENTRYPOINT 的区别

### 前言
在说 CMD、RUN 和 ENTRYPOINT 的区别前，先来说说 Dockerfile，Dockerfile 是构建容器镜像的方式之一，其通过一系列的指令参数来完成镜像的构建，而这些参数正是包含了 CMD，、RUN、COPY、ADD 和 ENTRYPOINT 等一系列指令。因此在实际应用中我们更多都是通过 Dockerfile 来完成镜像的构建。接下来列举一些 Dockerfile 常用的指令。

### Dockerfile 常用指令

- FROM

指定基础（base）镜像，本地有镜像则直接使用，否则直接在线拉取（pull）。

- MAINTAINER

Author，对作者的简单描述，自定义。

- COPY

将文件或目录从 build context 复制到镜像，其支持两种格式：COPY src dest 和 COPY[“src”,“dest”]

注：原目标（src）只能是文件或目录。

- ADD

与 COPY 类似，复制文件到镜像，不同的是，ADD 的 src 是归档文件（tar、zip、tgz 等），这些归档文件会被自动解压到 dest （镜像目标路径），无需手动解压。

- ENV

设置环境变量，该变量可被后面的指令使用。

- EXPOSE

指定容器中的进程会监听的某个端口，指定后 Docker 可以将该端口暴露出来。

- VOLUME

将文件或目录声明为 volume，同样 Docker 可以将该目录或文件映射出来。

- WORKDIR

为后面的 RUN、CMD、ENTRYPOINT、ADD、COPY 指令设置镜像中的当前工作目录。

- RUN

在容器中运行指令的命令。

- CMD

启动容器时运行指定的命令，Dockerfile 中可以有多个 CMD 指令，但只有最后一个生效，如果 docker run 后面指定有参数，该参数将会替换 CMD 的参数。

- ENTRYPOINT

同样，在 Dockerfile 中可以有多个 ENTRYPOINT 指令，也是只有最后一个生效，但与 CMD 不同的是，CMD 或 docker run 之后的参数会被当作参数传给 ENTRYPOINT。

###  三者的区别

#### Shell 和 Exec 格式
通常，我们有两种方式来指定 RUN、CMD 和 ENTRYPOINT 要运行的命令，即 Shell 和 Exec 方式。CMD 和 ENTRYPOINT 推荐使用 Exec 格式，其可读性更强。

- Shell 格式

```bash
RUN yum install -y vim
CMD echo "hello zhurs"
ENTRYPOINT echo "hello zhurs"

# 运行容器时返回如下结果
hello zhurs
```

当指令执行时，Shell 格式会调用 `/bin/sh -c [command]`。

- Exec 格式

> [“executable”, “param1”, “param2”]

```bash
RUN ["yum", "install", "-y", "vim"]
CMD ["bin/echo", "zhurs"]
ENV wd world
ENTRYPOINT ["/bin/echo", "hello, $wd"]

# 运行容器时返回如下结果
hello $wd

# 可看到运行容器时并没有调用/bin/sh -c 没有被shell解析（环境变量wd并没有被替换）。

# 如果希望使用环境变量，可做如下操作
RUN ["yum", "install", "-y", "vim"]
CMD ["bin/echo", "zhurs"]
ENV wd world
ENTRYPOINT ["bin/sh", "-c", "/bin/echo", "hello, $wd"]

# 此时就会返回如下结果
hello world

```

#### RUN
RUN 指令通常用于安装应用和软件包，每条 RUN 指令都会生成新的镜像。

```bash
...
RUN apt update && apt install -y git
...

```

像在安装一些基础工具或应用的时候，apt update 和 apt install … 最好放在一个 RUN 指令下执行，因为这能够保证每次安装的是最新的包，如果 apt update 在单独的 RUN 下运行，则 apt install … 会使用 apt update 创建的镜像，而这一层镜像可能是很久以前缓存的镜像文件。

#### CMD

该指令用于用户启动容器时，容器来执行的命令，该命令会在容器启动且 docker run 后面没有指定其他命令时执行，所以小结三种情况：

- docker run 没指定其他命令：则启动容器时运行 CMD 后的命令；

- docker run 指定了其他命令：则启动容器时运行 CMD 后的命令会被忽略；

- Dockerfile 中有多条 CMD 指令时，仅最后一条生效。

CMD 的三种格式：

- shell 格式：CMD <二进制可执行命令> <指令1> <指令2> 如：CMD yum install -y vim

- exec 格式：CMD [“二进制可执行命令”, “指令1”, “指令2”] 如：RUN [“yum”, “install”, “-y”, “net-tools”]

- CMD [“a”,“b”] 格式：该格式是为 ENTRYPOINT 提供使用，此时 ENTRYPOINT 就必须使用 exec 格式，否则不生效。

#### ENTRYPOINT

  该指令可以让容器以应用程序或者服务的形式运行。与 CMD 不同的是，不管 docker run … 后是否运行有其他命令，ENTRYPOINT 指令后的命令一定会被执行。

ENTRYPOINT 的两种格式：

- shell 格式：同 CMD；
- exec 格式：同 CMD。

ENTRYPOINT 的 exec 格式可以可执行由 CMD 提供的额外参数，具体如下：

```bash
...
ENTRYPOINT ["/bin/echo", "hello"] CMD ["world"]
...
```


运行容器时：

- docker run -it

  > 运行的容器后无任何参数

  ```bash
  # 输出
  hello world
  ```

- docker run -it myworld

  > 运行的容器后跟了 myworld 参数

  ```bash
  # 输出
  hello myworld
  ```

  

**ENTRYPOINT 小结：**

无论运行的容器命令后是否有其他参数，ENTRYPOINT 一律执行，如果 ENTRYPOINT 后跟随有 CMD 指令参数，则该参数的内容将会作为 ENTRYPOINT 指令参数。如果 docker run ... 后有参数，ENTRYPOINT 则使用该参数，而不会使用 CMD 的参数，为什么呢？因为 docker run ... 后指定了参数，那么根据前面说到的 CMD 后的参数将会被忽略掉（或叫被替换）。

但是上面的结论是针对 ENTRYPOINT 的 exec 格式而言的，如果是 shell 格式，ENTRYPOINT 将会忽略掉任何 CMD 或 docker run … 提供的参数。当然 ENTRYPOINT 的 shell 格式也是会必然执行的。

#### 如何选择 CMD 和 ENTRYPOINT

> 根运行容器的属性来合理选择

- 如果运行的是一个 MySQL 容器，则优先使用 exec 格式的 ENTRYPOINT 指令，因为 CMD 不仅可以为 ENTRYPOINT 提供默认参数，同时在 `docker run ...`（带参数）的时候，该参数也会替换 CMD　默认参数。

- 如果只是简单的设置容器默认的启动命令，使用 CMD 即可，用户只需在 `docker run ...` 后添加参数即可替换默认值。
  

### 小结

- RUN：执行命令并创建新的镜像层，常用于安装软件包；
- CMD：设置容器启动后默认执行的命令及其参数，但 docker run 后跟参数时会替换（忽略） CMD；
- ENTRYPOINT：配置容器启动时运行的命令。