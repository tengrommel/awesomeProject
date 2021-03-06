# json

# 单元测试

# Go协程和主线程

# runtime

# channel

- 为什么需要channel
    - 主线程在等待所有的goroutine全部完成的时间很难确定
    - 如果主线程休眠时间长了，会加长等待时间，如果等待时间短了，可能还有goroutine处于工作状态，这是也会随着主线程推出而销毁
    - 通过全局变量加锁同步来实现通讯，也并不利于多个协程对全局变量的读写操作

- channel的介绍
    - channel本质就是一个数据结构-队列
    - 数据是先进先出
    - 线程安全，多goroutine访问时，不需要加锁，就是说channel本身就是线程安全的
    - channel时有类型的，一个string的channel只能存放string数据类型

- channel的注意事项
    - channel是引用类型
    - channel必须初始化才能写入数据，即make后才能使用

- 管道的关闭
> chan关闭了可以读但不能写

- channel的遍历
    - 在遍历时，如果channel没有关闭，则会出现deadlock的错误
    - 在遍历时，如果channel已经关闭，则会正常遍历数据，遍历完后，就会退出遍历

- select
    - 使用select可以解决从管道取数据冲

- 反射
    - 反射的使用场景
        - 使用函数的反射编写适配器
    - 基本介绍
        - 反射可以在运行时动态获取变量的各种信息，比如变量的类型(type)，类别(kind)
        - 如果是结构体变量，还可以获取到结构体本身的信息（包括结构体的字段、方法）
        - 通过反射，可以修改变量的值，可以调用关联的方法

- const
    - 常量在定义的时候必须初始化

- redis
    -Redis连接池

- What is message passing?
> Strategies
- interface
- Channels

- Composition
   - Inheritance
   > Behavior reuse strategy where a type is based upon another type,
   allowing it to inherit functionality from the base type.
   - Composition
   > Behavior reuse strategy where a type contains objects that have desired functionality.
   The type delegates calls to those objects to use their behaviors.

- Polymorphism
> The ability to transparently substitute a family of types that implement
a common set of behaviors.

