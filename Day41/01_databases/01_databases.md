# 数据库
- 原理
  - 一种数据结构，但是比常见的数据结构更麻烦，因为要持久化，从磁盘或硬盘中存储和拿取
- 作用
  - 持久化数据，掉电不丢失

- 需要解决的问题
  - IO读取慢
    - 分摊任务
  - 高并发
- 概念
  - 认证：确认用户身份
  - 权限：确认不同身份的访问范围

- SQL
  - 数据查询与分析的第一语言
  - 一种对关系型数据库进行查询、更新、管理的可编程语言
  - 语言规范
    - SQL语句大小写不敏感
      - 一般建议，SQL的关键字、函数等大写
    - SQL语句末尾应该使用分号结束
    - 注释
      - 多行注释 /*注释内容*/
      - 单行注释 -- 注释内容
      - MySQL注释可以使用 #
  - 分类
    - DDL数据定义语言
      - 负责数据库定义表或库等、数据库对象定义，由CREATE、ALTER、DROP三种语句组成
    - DML数据操作语言
      - 负责对数据库对象操作，CRUD增删改查
    - DCL数据控制语言
      - 负责数据库权限访问控制，由GRANT（授权）和REVOKE（撤销）两个指令组成
    - TCL事务控制语言
      - 负责处理ACID事务，解决并发问题，支持commit、rollback指令
          - A 原子性，不可拆分，事务没做成功，缺一步就要回滚，要么不做，要么一次做好
          - C 一致性，多个事务操作应该跟使用串行、并行思路的结构是相同的
          - I 隔离性，解决并发，多人操作同一数据不影响
          - D 持久化，数据落在磁盘上

    - 语法
      - 授权、撤销
          - GRANT授权、REVOKE撤销

```txt
                        GRANT ALL ON employees.* TO 'wayne'@'%' IDENTIFIED by 'your_password'; -- MySQL8+ 不支持该语法，见创建用户
                        REVOKE ALL ON *.* FROM wayne;
```

          - 补充说明
              - *               为通配符，指代任意库或者任意表
              - *.*             所有库的所有表
              - employees.*     表示 employees 库下所有的表
              - %               为通配符，它是 SQL 语句的通配符，匹配任意长度字符串，即所有字符

      - 用户管理
          - 创建用户

```txt
                        CREATE USER 'wayne'@'192.168.%' IDENTIFIED BY 'your_password';
                        GRANT ALL ON employees.* TO 'wayne'@'192.168.%' -- 授权
```

              - 补充说明
                - 两段'wayne'@'192.168.%'必须完全一致，表示用户在该网段下可以访问 employees 数据库下的所有表

          - 删除用户（慎用）

```txt
                        DROP USER 'wayne'@'192.168.%';
                        DROP USER 'wayne'; -- DROP USER 'wayne'@'%';

```

      - 表库操作
          - 创建数据库
              - 库是数据的集合，所有数据按照数据模型组织在数据库中

```txt
                            CREATE DATABASE IF NOT EXISTS test CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
                            CREATE DATABASE IF NOT EXISTS test CHARACTER SET utf8;
```

                - 补充说明
                    - 字符集标准若未指定就默认继承
                    - CHARACTER SET 指定字符集。 utf8mb4 是 utf8 的扩展，支持 4 字节 utf8mb4（支持emoji标签），需要 MySQL5.5.3+。
                    - COLLATE 指定字符集的校对规则，用来做字符串的比较的。例如 a、A 谁大？
                    - IF NOT EXISTS：避免数据库已存在时报错，是安全创建的常用写法
                    - utf8mb4：MySQL 推荐的标准字符集，完整支持 Unicode（包括 emoji 等 4 字节字符），向下兼容utf8
                    - utf8mb4_general_ci：ci代表case insensitive（大小写不敏感），是通用的校对规则
                    - utf8：在 MySQL 中等同于utf8mb3，仅支持 3 字节 Unicode，不推荐用于新业务

          - 删除数据库（慎用）

```txt
                        DROP DATABASE IF EXISTS test;

```

          - 创建表
              - 表分为行和列，MySQL 是行存数据库。数据是一行行存的，列必须固定多少列。
              - - 行 Row，也称为记录 Record 或元组
              - - 列 Column，也称为字段 Field 或属性
              - - 字段的取值范围叫做域 Domain。例如 gender 字段的取值就是 M 或者 F 两个值。

```txt
                        +--------+------------+------------+-----------+--------+------------+
                        | emp_no | birth_date | first_name | last_name | gender | hire_date  |
                        +--------+------------+------------+-----------+--------+------------+
```

              - | 10001  | 1953-09-02 | Georgi     | Facello   | 1      | 1986-06-26 |  → row、行、记录、元组

```txt
                        | 10002  | 1964-02-02 | Bezalel    | Simmel    | 2      | 1985-11-21 |
                        | 10003  | 1959-12-03 | Parto      | Bamford   | 1      | 1986-08-28 |
                        | 10004  | 1954-05-01 | Chirstian  | Koblick   | 1      | 1986-12-01 |
                        | 10005  | 1955-01-21 | Kyoichi    | Maliniak  | 1      | 1989-09-12 |
                        | 10006  | 1953-04-20 | Anneke     | Preusig   | 2      | 1989-06-02 |
                        | 10007  | 1957-05-23 | Tzvetan    | Zielinski | 2      | 1989-02-10 |
                        | 10008  | 1958-02-19 | Saniya     | Kalloufi  | 1      | 1994-09-15 |
                        | 10009  | 1952-04-19 | Sumant     | Peac      | 2      | 1985-02-18 |
                        | 10010  | 1963-06-01 | Duangkaew  | Piveteau  | 2      | 1989-08-24 |
                        +--------+------------+------------+-----------+--------+------------+
```

                - ↓
                - 列、字段、Field、Column

```txt
                            CREATE TABLE `employees` (
                              `emp_no` int(11) NOT NULL,
```

                  - `birth_date` date NOT NULL,

```txt
                              `first_name` varchar(14) NOT NULL,
                              `last_name` varchar(16) NOT NULL,
                              `gender` smallint(6) NOT NULL DEFAULT '1' COMMENT 'M=1, F=2',
```

                  - `hire_date` date NOT NULL,
                  - PRIMARY KEY (`emp_no`)
                - ) ENGINE=InnoDB DEFAULT CHARSET=utf8;

              - 补充说明
                - 元组理论上谁在前谁在后是没有关系的，但可能就作为存储顺序，存储顺序又会决定读写效率
                - “InnoDB”为存储引擎，可选其他
                - 反引号标注的名称，会被认为是非关键字，使用反引号避免冲突。
                - navicat可以右键表选择新建表手动输入创建表
                - navicat可以点击上方的“模型”手动建表（建模）

          - 删除表（慎用）

```txt
                        DROP TABLE `employees`;
                        DROP TABLE IF EXISTS `employees`;

```

          - DESC

```txt
                        {DESCRIBE | DESC} tbl_name [col_name | wild] 查看列信息
                            DESC employees; -- 列出所有字段
                            DESC employees '%name'; -- 列出字段名称后缀为name的所有字段

```

- 数据库重要概念
  - 关系
    - 在关系数据库中，关系就是二维表，由行和列组成
      - - 行、列、字段
      - - 维数：关系的维数指关系中属性的个数
      - - 基数：元组的个数，条目数，即不一样的行的行数。
    - 注意在关系中，属性的顺序并不重要。理论上，元组顺序也不重要，但是由于元组顺序与存储相关，会影响查询效率

  - 举例

```txt
            +----+--------+------+-----------+--------+
            | id | emp_no | name | login_name| gender |
            +----+--------+------+-----------+--------+
            | 1  | 10001  | tom  | alex      | M      |
            | 2  | 10002  | tom  | alex123   | F      |
```

  - 候选键
    - - id 设置为自增 int 型
    - - emp_no 工号，不重复，不能为null
    - - lodin_name 登录名，避免重复，不能为null。条件比对：WHERE lodin_name='alex' and password='xxx'，生产中，该需求需要设置索引，因为不可能全表遍历扫描
    - id和emp_no都是候选键，因为他们能够代表一行，如果lodin_name不重复，则也当作候选键
    - 多个字段组合才能确定一行，这几个字段合起来是一个候选键，目前不建议作为主键

  - NOT NULL 约束

  - 可以作为唯一键 unique key
    - - cellphone 不重复，可以为null
    - - email 不重复，可以为null
    - 唯一键约束：有值必须唯一，无值即为null
    - 构建唯一键索引（底层也是B+树）
    - 比主键效率低一些唯一确定一行

  - 主键primary key
    - id好处，无业务含义，简单高效
    - 主键约束：唯一，不重复，不能为null
    - 不能为null
    - 作用：数据库中最高效的定位唯一一行，快速定位数据
    - 为何高效？跟存储结构有关（B+树），构建了主键索引
    - innodb，要求表中必须明确指出主键
    - 优先利用主键编程，因为最高效
    - 一张表至多只能有一个主键

  - 索引index
    - 如果字典的目录，快速检索，可以定位一条或多条（少数几条），如何建索引也是需要考虑的问题
    - 所谓目录，占用空间，空间换时间，每增删改一行都要写硬盘，意味着IO高
    - 不用主键索引、唯一键索引，相当于遍历新华字典（全表扫描），效率慢
    - 条件快速比对，对相应的字段增加合适的索引，使用索引提高了检索性能，降低降低了写入性能
    - 有没有唯一的要求？
      - 没有。有可能一个值对应多行
    - 问题：我经常查男生、女生，请问在gender上要不要建立索引？
      - 不要。区分度不大，且每类有大量的数据，无法快速定位
      - 这个需求就有问题
      - 区分度不大的不要建索引，不经常加查询的不要建索引
    - 优劣：可能提高了检索效率，但是影响写入效率，这需要达到一个平衡
    - 索引的数据结构：
      - 哈希hash      y=hash(x)，适合条件都是等值的比较O(1)，进行范围查找时效率极差
      - B+树B+Tree    可以范围，也可以等值，但是等值查找比hash慢


  - 树Tree
    - 父子结构数据结构
    - 度数等于最多子节点数量
    - 每个节点只能至多一个父节点，除Root根节点没有父。（多个树就是图Graph）
      - 如果要求该树的最大度数为2，那么就是二叉树
      - 如果继续要求左右子节点有序，这就是二叉有序树
      - 如果继续要求左子树小于其根，这就是二叉排序树，有可能很不平衡，最大值有倾斜现象，导致探索深度太深从而引起过多的IO
    - 降低深度办法
      - 1、多开叉
      - 2、再平衡 二叉平衡树，有序、排序
    - 辅助学习网址
      - https://www.cs.usfca.edu/~galles/visualization/Algorithms.html

  - B+树
    - 一种特殊的 m 叉树，也是一种平衡树，兼顾文件系统存储和操作的考虑（为了减少IO），同时减少了树的深度。
    - B+树节点组织成一棵树。节点分为内部节点和叶子节点，最下一层为叶子节点，上层都称为内部节点
    - 内部节点不存储数据，叶子节点不存储指针，存储主键和B+树。
    - 叶子节点深度一致。叶子节点包含所有索引字段值。
    - 每个 leaf node 上保存数据，所有的 leaf node 组织成双向链表。假设读取 16 到 22 的数据，找到 18 后（大于 16 的第一个数据是 18），顺着链表往后遍历读取即可。
    - 每一个节点大小为 1 页，1 页默认为 16KB。
    - 树高为 3，就可以存储 2 千多万条记录，主键查询只需 3 次 IO 操作。

    - 问题：假设id 1~10000000 有一千万条数据，可以使用WHERE id > 5000吗？
        - 相当于遍历，数据库访问磁盘IO、网络IO效率慢

# SQL 必学必会
- 关系操作
  - 关系：关系在数据库中，就是二维表。关系数据库就是对表的操作
    - 选择：又称为限制，是从关系中选择出满足给定条件的元组
    - 投影：在关系上投影就是从选择出若干属性列组成新的关系
    - 连接：将不同的两个关系连接成一个关系
- 查询

```txt
        SELECT * FROM book -- 指定字段，* 代表所有字段（不要这样做，用什么字段就投影什么字段），FROM表示从book表进行全表扫描，慎用，效率低
```

  - 要限制返回结果的最大条目

```txt
            LIMIT n 拿几个，默认从索引0开始
            OFFSET n 从n开始，可用于简单分页

```

  - 条件查询WHERE 条件为真将被选择出来

```txt
            WHERE last_name like '%p%'
            WHERE last_name like '% P%' 这是一种模糊查询，但全表扫描，效率低，少量数据随便，数据越多越不要用
            WHERE last_name like '% P' 模糊查询，少量数据随便，数据越多越不要用
            WHERE last_name like 'P%' and xxx 模糊查询，左前缀偶尔还可以用用，但是最好和其他有索引的字段配合
```

  - 排序

```txt
            ORDER BY emp_no DESC -- 默认升序（ASC），降序（DESC）
```

      - emp_no是主键，主键构成B+，主键天然就是排过序的
      - 排序操作要在投影结束后再排序
    - 一般字段排序，由于不能借助索引，效率不高

  - 聚合 reduce合并结果
    - 使用聚合函数，统计
      - sum累加求和
      - avg算术平均
      - min最小值
      - max最大值
      - count统计个数

```txt
                     SELECT count(emp_no) from employees 按照主键统计个数，1行1个值
                     SELECT count(emp_no), min(emp_no), max(emp_no), avg(emp_no), sum(emp_no) from employees 1行，5列
```

    - 分组就是分摊，一摊统计一个出一条记录

  - 分组

```txt
            SELECT count(emp_no) from employees 一旦使用了聚合函数，就默认分组，没有显式的指定分组就当作一组，一组就是一摊出一条
```

    - 分组语句

```txt
                GROUP BY 分组字段们
```

      - SELECT后面写分组字段们或字段们或聚合函数是可以的，写非分组字段无意义

```txt
            SELECT gender, emp_no, count(emp_no) from employees GROUP BY gender, emp_no
```

      - 按照gender和emp_no值组合起来的值分摊

```txt
            GROUP BY 指定分组字段们
```

      - 如果是表中直接筛选使用WHERE（写在FROM之后）
      - 如果是分组统计后的结果才能筛选，聚合使用HAVING筛选过滤

  - 子查询
    - 查询语句可用嵌套，内部查询就是子查询
    - 子查询必须在一组小括号中
    - 子查询中不能使用ORDER BY
    - 子查询效率不高

  - JOIN（慎用）

```txt
            SELECT * FROM employees, salaries	-- 隐式连接
            SELECT * FROM employees CROSS JOIN salaries -- 笛卡尔乘积（交叉连接），禁止使用
```

    - 加条件使用
      - 对连接表加条件 ON 做等值条件，和JOIN配合
    - 内连接INNER JOIN
      - 省略为JOIN，等值连接，只选某些field相等的元组（行），使用ON限定关联的结果
      - 相当于求等值后的交集
      - 自然连接，特殊的等值连接，会去掉重复的列
      - 用的少
    - 外连接OUTER JOIN
      - 分为
          - 左外连接LEFT JOIN：左边所有数据不管能否配对，都要出来（数量 = 能配对 + 不能配对）
          - 右外连接RIGHT JOIN：右边所有数据不管能否配对，都要出来
          - 全外连接
    - 自连接
      - 自己和自己连接
      - 用于同一张表内部实现层级结构
    - JOIN使用总结
      - 对于小规模数据，随便JOIN
      - 对于大规模数据，应减少JOIN字数，也就是减少表直接连接次数，因为数据直接比较都是在内存中完成的，大量数据要从IO设备搬到内存
      - 必要时，对某些字段要在其他表中增加冗余字段来减少JOIN。JOIN字段使用索引会被优化

  - 执行顺序

```txt
            FROM -> ON -> JOIN -> WHERE -> GROUP BY -> HAVING -> SELECT -> DISTINCT -> ORDER BY -> LIMIT
```
