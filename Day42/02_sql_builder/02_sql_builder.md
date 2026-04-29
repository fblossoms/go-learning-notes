# SQL语句来源
- 1、手写
- 2、SQLBuilder 构造生成SQL语句，就是个字符串，拿着SQL语句，用Query（SQL语句）
  - SQLBuilder是一个用于生成SQL语句的库
- 3、ORM库内部可以构造SQL语句，不过通过面向对象的方式调用函数，函数内部生成语句，不用手写
  - CRUD 面向对象的方法，不用写SQL语句