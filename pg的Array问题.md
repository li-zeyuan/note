# pg的Array问题

### 表结构

```sql
CREATE TABLE Employees (
   id int PRIMARY KEY,
   name VARCHAR (100),
   contact TEXT []
);
```

### insert

- 单个、多个插入：使用array[ ]、\`{ }`、::text[]都能插入成功

```sql
-- 成功
INSERT INTO Employees 
VALUES
   (
      1,
      'Alice John',
       -- 使用array，
      ARRAY [ '(408)-743-9045',
      '(408)-567-7834' ]
   );
```

```sql
-- 成功
INSERT INTO Employees 
VALUES
   (
    2,
      'Kate Joel',
      '{"(408)-783-5731"}'
   ),
   ( 
      3,
      'James Bush',
       -- 插入时可以不用::text[]强转
      '{"(408)-745-8965","(408)-567-78234"}'
   );
```

### update

- 更新单个：使用array[ ]、\`{ }`、::text[]都能更新成功

- 更新多个：

  - 使用array[ ]、::text[ ]、能更新成功

  - \`{ }`不能失败，解决方法：model struct加pg tag，如：

  - ```
    LessonId   []int64     `sql:"lesson_id,type:bigint[]"`
    ```

```sql
-- 成功
UPDATE Employees
-- 更新当个，可以不用强转；或array
SET contact = '{"707","38"}'
WHERE id = 1
```

```sql
-- 不成功
UPDATE oto.Employees
SET contact = _data.contact
FROM (VALUES(
    2,
    '{"(408)-783-5731"}'
   ),
   (
     3,
     -- 批量插入时需要::text[]强转或array
     '{"(408)-745-8965","(408)-567-78234"}'
   )) AS _data(id, contact)
WHERE Employees.id = _data.id
```

