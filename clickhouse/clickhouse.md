# Clickhouse
### SQL
### 表引擎
##### MergeTree
- 特点：
  - 支持分区和索引
- order by（required）：分区内按照该字段排序
- primary key（optional）：创建一级索引（稀疏索引），没有唯一约束，**必须是order by字段的前缀**
- partition by（optional）：分区规则，一般是时间
- settings（optional）：一些额外控制参数，如index_granularity索引粒度，默认8192
- TTL：支持列ttl，表级ttl
- 参考：
  - https://clickhouse.com/docs/en/engines/table-engines/mergetree-family/mergetree
- 目录结构
  ```
  root@/data/clickhouse# tree -L 1
  .
  |-- data                    // 数据、表元数据
  |-- format_schemas
  |-- log                     // ck-server日志文件
  `-- tmp
  
  root@/data/clickhouse/data# tree -L 1
  .
  |-- data                    // 数据、索引文件
  |-- dictionaries_lib
  |-- flags
  |-- metadata                // 表元数据
  |-- metadata_dropped
  |-- preprocessed_configs
  |-- status
  `-- store                   // data下数据文件是以软连接到store目录
  
  // 表partition
  root@/data/clickhouse/data/data/{database}/{table}# tree -L 1
  .
  |-- 20220917_20_20_0        // patition
  |-- 20220917_21_21_0
  |-- detached                // 记录损坏partition
  `-- format_version.txt      // version
  
  // partition目录
  root@/data/clickhouse/data/data/{database}/{table}/{patition}# tree -L 1
  .
  |-- checksums.txt                   // 校验文件
  |-- columns.txt                     // 列信息（字段名、类型）
  |-- count.txt                       // 总数
  |-- data.bin                        // 数据
  |-- data.mrk3                       // 数据标记文件，索引文件会用到该标记
  |-- default_compression_codec.txt   // 压缩
  |-- minmax_timestamp.idx            // 分区minimal索引
  |-- partition.dat                   // 分区信息
  `-- primary.idx                     // 主键索引
  ```
- partition命名规则
  ```
  20220917_1_1_0
  [分区名]-[最小分区块编号]-[最大分区块编号]-[合并数次]
  
  分区名：跟partition by参数有关，有整数字符串、日期、哈希值
  分区块编号：新生成的分区自增
  合并数次：合并一次加1
  ```
### 索引
- 参考：https://clickhouse.com/docs/en/engines/table-engines/mergetree-family/mergetree
##### 主键索引
##### 分区索引
##### 跳数索引
### explain
### 副本表
### 分片
### 常见问题
### 