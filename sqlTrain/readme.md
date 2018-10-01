# sql

## 分类存储
- 为了防止冗余
- 建立外键

## DQL

语法：
>select 查询列表 from 表名；

类似System.out.println(打印东西)

特点：
- 1、查询列表可以是：表中的字段、常量值、表达式、函数
- 2、查询的结果是一个虚拟的表格

1、查询表中的单个字段
>select last_name from employees;

2、查询表中的多个字段
>select last_name, salary, email from employees;

3、查询表中的所有字段
>select * from employees;

4、查询常量值 表达式 函数
- select 100;
- select 'john'
- select 100%98
- select VERSION()

5、起别名
> 便于理解 如果要查询的字段有重名的情况，使用别名可以区分来
- SELECT 100%98 AS 结果;
- SELECT last_name AS 姓, first_name AS 名 FROM employees;
- SELECT last_name 姓, first_name 名 FROM employees;

6、去重
- select DISTINCT department_id from employees;

7、+号的作用
SELECT last_name+first_name AS 姓名 FROM employees;

8、CONCAT('a', 'b', 'c');

9、分类

   - 按条件表达式筛选

    > < = != <> >= <=
   - 按逻辑表达式筛选
   > && || ! and or not
   - 模糊查询
   >like  between  and  in  is null

10、like 一般和通配符搭配使用

- % 包含零个字符
- _ 单个字符

11、between and 在什么之间

- 使用between and 可以提高语句的简洁度
- 包含临界值

12、in 查询员工的工种编号
>判断某字段的值是否属于in列表中的某一项

13、<> 或 = 不能判断null值

14、安全等于<=>
>可以判断null值

15、排序查询
> select * from employees;

- desc降序
- asc生序

order by 支持别名

select LENGTH(last_name) 字节长度, last_name, salary FROM employees ORDER BY LENGTH(last_name) DESC;

16、常见函数

功能：类似于java的方法，将一组逻辑语句封装在方法体中，对外暴露方法名

- 1、隐藏了实现细节
- 2、提高代码的重要性

>调用： select 函数名（实参列表）

- 1 叫什么（函数名）
- 2 干什么（函数功能）

分类：

- 1、单行函数
- 2、分组函数

## 一、字符函数

- 1、length 获取参数值的字节个数
- 2、concat 拼接字符串
- 3、upper、lower
- 4、substr、substring

## 二、数学函数

- round四舍五入
- ceil向上取整

## 三、日期函数
- now 返回当前系统日期+时间
- curdate 返回当前日期
- curtime 返回当前时间
- str_to_date
- format

## 四、流程控制

- if函数： if else的效果
- case函数：
    - switch case 的效果
    - case 要判断的字段或表达式

   when 常量1 then要显示的值1或语句1；
## 分组函数
分类
> sum 求和、avg 平均值、max最大值、min最大值、count计算个数

- 1、简单的使用
>SELECT SUM(salary) from employees;

- 2、参数支持哪些类型

- 3、是否忽略NULL

- 4、distinct去重

- 5、count

    - count(*)统计总行数
    - count(1)