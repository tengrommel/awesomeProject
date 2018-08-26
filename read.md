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
