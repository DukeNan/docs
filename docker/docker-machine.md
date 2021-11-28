### vagrant

```bash 
vagrant init centos/7
vagrant up
vagrant status
# 进入虚拟机
vagrant ssh default
# 停止虚拟机
vagrant halt
# 停止之后删除
vagrant destroy
```

### docker-machine

```bash
# 创建一个虚拟机 默认driver vi
docker-machine create demo
docker-machine ls
docker-machine stop demo
docker-machine rm demo

```

### docker swarm

```bash
# manager
docker swarm init --advertise-addr=192.168.99.101
docker swarm join --token SWMTKN-1-39xkb5sfwiilpuco7f7uzqn60ma4ybk0ms6j4qlsegqi5go2xu-3ifw8kw22japnin2nra4wg3sc 192.168.99.101:2377
# 查看节点
docker node ls
docker service ls
docker service ps web1
# 查看端口转发
sudo iptables -t nat -nvL
docker service scale helloworld=3

docker stack deploy -c docker-compose-v3.yml flask-redis
# 查看参数
docker stack ls


```

