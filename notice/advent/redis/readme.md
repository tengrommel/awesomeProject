# 实现二级缓存
- 程序运行起来之后，提示“请输入命令：”，输入getall，查询并显示所有人员信息
- 第一次时查询mysql并将结果缓存在redis，设置60秒的过期时间
- 以后的每次查询，如果redis有数据就从redis加载，没有则重复上一步的操作

# 并发概念

异步async并行：多个任务并发执行

同步sync串行：多个任务依次执行

阻塞block某个并发任务由于拿不到资源无法运行


# 异步回调
> A线程唤起B线程，令其干活同时给B一个回调函数命令B在干完活以后，执行这个回调函数