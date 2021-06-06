# Mysql查询优化

### 背景

- 应对单表百万级的数据量，接口频繁超时
  - 原因：1、mysql慢查询
- 单条sql语句执行时间长，高达30s，mysql机器CPU瞬间打满
  - 原因：1、mysql慢查询；2、并发查表

### 索引

##### 联合索引数据结构

我们都知道联合索引遵循左前缀原则，这一特性其实是由其数据结构所决定的．index(col1, col2, col3)数据结构图：

![](https://github.com/li-zeyuan/access/blob/master/img/unified_index_data_structure.png)

联合索引数据结构特性：

- 也是一棵B+Tree，非叶子节点保存着col1
- 叶子节点按照索引的创建顺序保存
- 访问索引按照col1->col2->col3的顺序，即联合索引创建顺序

### count优化

- todo

### group by优化

- todo

##### Loose Index Scan（松散的索引扫描）

- 定义：

  - ```
    The most efficient way to process GROUP BY is when an index is used to directly retrieve the grouping columns. With this access method, MySQL uses the property of some index types that the keys are ordered (for example, BTREE). This property enables use of lookup groups in an index without having to consider all keys in the index that satisfy all WHERE conditions. This access method considers only a fraction of the keys in an index, so it is called a Loose Index Scan.
    ```

  - 官方的解释如上，大概意思是：group by使用索引直接检索索引列，不用考虑where子条件，仅考虑group by中这部分键

  - 是group by查询最优的方式 

  - **explain**显示`Using index for group-by`在 `Extra`列中

- 满足条件：idx(c1,c2,c3)`表上 的索引 `t1(c1,c2,c3,c4)

  - 查询在一个表
  - group by遵循最左前缀，如`SELECT c1, c2 FROM t1 GROUP BY c1, c2; `
  - 聚合函数只能是min( )和max( )，并且是指向同一列，如`SELECT MAX(c3), MIN(c3), c1, c2 FROM t1 WHERE c2 > *const* GROUP BY c1, c2; `
  - 若有其他where条件，必须是常数，MIN()或MAX() 函数的参数例外．如：`SELECT MAX(c3), MIN(c3), c1, c2 FROM t1 WHERE c2 > *const* GROUP BY c1, c2; `
  - 对于索引中的列，必须对整列值进行索引，而不仅仅是前缀。例如，与 `c1 VARCHAR(20), INDEX (c1(10))`．

- 例子

  - todo

##### Tight Index Scan（紧凑的索引扫描）

- 定义
  - 没有满足松散索引扫描的条件时，仍然可以避免为`GROUP BY`查询创建临时表。
- 满足条件：`idx(c1,c2,c3)`表上 的索引 `t1(c1,c2,c3,c4)`
  - group by条件有索引＂间隙＂．如：`SELECT c1, c2, c3 FROM t1 WHERE c2 = 'a' GROUP BY c1, c3;`
  - group by条件不是最左前缀．如：`SELECT c1, c2, c3 FROM t1 WHERE c1 = 'a' GROUP BY c2, c3;`
- 例子

### district优化

### 总结

- 并发查表慎用，极容易打满mysql机器CPU
- sql语句需要经常explain
- 联合索引是一个B+Tree，叶子节点指向有序索引字段
- group by最优是Loose Index Scan

### 参考

- mysql文档：https://dev.mysql.com/doc/refman/5.7/en/

- group by优化：https://dev.mysql.com/doc/refman/5.7/en/group-by-optimization.html
- 索引数据结构和算法逻辑：http://blog.codinglabs.org/articles/theory-of-mysql-index.html