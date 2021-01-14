### 1.缓存的数据类型
- 数值
- 数据库记录（对象）
- 数据库查询语句
    - 语句采用hash算法 md5 
- 视图响应结果
- 一个页面
### 2.缓存数据有效期和淘汰策略
#### 2.1有效期作用
- 节省空间
- 做到数据弱一致性，有效期失效后，可以保证数据的一致性
#### 2.2过期策略
- 定时过期：每个记录单独追踪有效期

- 惰性过期：只在查询数据的时候才判断数据是否过期

- 定期过期：每隔100ms 检查有哪些数据过期  

  **redis选用的过期策略为 惰性过期+定期过期**

#### 2.3内存淘汰策略

- LRU(Least Recently Used)

  - 以操作过的时间前后来选择淘汰
- LFU(Least Frequently Used)
  - 以操作的频率大小来选择淘汰
  - 定期衰减 ，所有使用次数减半

### 3.缓存模式

#### 3.1缓存操作使用模式

- Cache Aside 缓存边缘
- Read-through 通读
- Write-through 通写
- Write-behind caching  缓存之后写入
  - 更新操作只更新缓存，不更新数据库，由cache异步批量同步数据库。

#### 3.2缓存数据更新方式

- 先更新数据库，在更新缓存
  - 问题：这种做法最大的问题就是两个并发的写操作导致脏数据

- 先删缓存，在更新数据库
  - 问题：两个并发的读和写操作导致脏数据
- 先更新数据库，再删除缓存
  - 问题：两个并发的读和写操作导致脏数据，数据库的写操作会比读操作慢得多，而且还要加锁，出现概率不大

#### 3.3项目使用

- 使用Read-throught + Cache aside

  - 构建一层抽象出来的缓存操作层，负责数据库查询和Redis缓存存取，在Flask的视图逻辑中直接操作缓存层工具。

- 更新采用先更新数据库，再删除缓存

  ```总结：读操作之后，才去设置缓存，
  总结：
  	读操作之后，才去设置缓存，
  	而写操作完mysql之后，一般不会修改缓存数据，而是删除缓存数据，等到下一次读操作的时候，才设置缓存
  ```

### 4.缓存问题

#### 4.1缓存穿透

- 问题
  - 利用频繁去访问缓存中没有的数据，使缓存失去意义
- 解决方法
  - 对于返回为NULL的依然缓存
  - 制定一些规则过滤一些不可能存在的数据，小数据用BitMap，大数据可以用布隆过滤器

#### 4.2缓存雪崩

- 问题
  - 指缓存不可用或者大量缓存由于超时时间相同在同一时间段失效，大量请求直接访问数据库，数据库压力过大导致系统雪崩
- 解决方法
  - 给缓存加上一定区间内的随机生效时间
  - 采用多级缓存，不同级别缓存设置的超时时间不同
  - 利用加锁或者队列方式避免过多请求同时对服务器进行读写操作(串行)

### 5.缓存设计

### 6.持久存储

#### 6.1 RDB快照持久化（默认开启）

- 触发机制

  - 定期触发

    ```python
    redis配置文件中：
    save 900 1    # 900秒中有1次及以上修改到redis中数据
    save 300 10   # 300秒中有10次及以上修改到redis中数据
    save 60 10000 # 60秒中有10000次及以上修改到redis中数据
    ```

  - 执行BGSAVE命令，手动触发RDB持久化

  - SHUTDOWN命令，关闭redis时触发

#### 6.2 AOF 追加文件持久化

- redis配置文件中：

  ```
  appendonly yes  # 是否开启AOF
  appendfilename "appendonly.aof"  # AOF文件
  ```

- 触发机制

  ```python
  # appendfsync always   # 执行每个命令都记录
  appendfsync everysec   # 每秒记录
  # appendfsync no   # 交给操作系统决定写到操作系统的时机
  ```

#### 6.3总结

- redis允许我们同时使用两种机制，通常情况下我们会设置AOF机制为everysec 每秒写入，则最坏仅会丢失一秒内的数据。

### 7.Redis高可用

#### 7.1Redis主从同步

- 客户端中通过命令 slaveof 主机地址 主机端口  来设置为从机

#### 7.2Redis哨兵机制

- 哨兵其实本身也是一个可执行程序，像我们启动redis客户端一样，只是它是用来看护redis实例进程的。

- 功能
  - Monitoring 监控
  - Notification 通知
  - Automatic failover 自动故障转移
  - Configuration provider 配置提供程序
- 启动方式
  - redis-sentinel sentinel.conf

### 8.Redis集群

- 创建集群命令
```python
redis-trib.rb create --replicas 1 192.168.192.200:7000 192.168.192.200:7001 192.168.192.200:7002 192.168.192.200:7003 192.168.192.200:7004 192.168.192.200:7005
```

- 终端中连接到redis集群

  `redis-cli -c -p 7000`

- 代码中连接到redis集群

  ```python
  from rediscluster import StrictRedisCluster
  # redis 集群
  REDIS_CLUSTER = [
      {'host': '127.0.0.1', 'port': '7000'},
      {'host': '127.0.0.1', 'port': '7001'},
      {'host': '127.0.0.1', 'port': '7002'},
  ]
  
  from rediscluster import StrictRedisCluster
  redis_cluster = StrictRedisCluster(startup_nodes=REDIS_CLUSTER)
  
  # 可以将redis_cluster就当作普通的redis客户端使用
  redis_master.delete(key)
  ```