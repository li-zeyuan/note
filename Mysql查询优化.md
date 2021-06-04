# Mysql查询优化

### 背景

- 应对单表百万级的数据量，接口频繁超时
  - 原因：1、mysql慢查询
- 单条sql语句执行时间长，高达30s，mysql机器CPU瞬间打满
  - 原因：1、mysql慢查询；2、并发查表

### 索引

### count优化

### district 和 group by

### 总结

- 并发查表慎用，极容易打满mysql机器CPU
- sql语句需要经常explain
- 

### 参考

- mysql文档：https://dev.mysql.com/doc/refman/5.7/en/

- group by 优化：https://dev.mysql.com/doc/refman/5.7/en/group-by-optimization.html
- 索引优化：