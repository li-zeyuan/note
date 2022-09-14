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