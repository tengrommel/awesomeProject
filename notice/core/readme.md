# 队列

- 队列是一个有序列表，可以用数组或是链表来实现。
- 遵循先入先出的原则。即：先存入队列的数据，要先取出。后存入的要后取出。

数组模拟队列

- 队列本身是有序列表，若使用数组的结构来存储队列的数据，则队列数组的声明如下其中maxSize是该队列的最大容量。
- 因为队列的输出、输入是分别从前后端来处理，因此需要两个变量front及rear分别记录队列前后端的下标，front会随着数据输出而改变，而rear则是随着数据而变化。

先完成一个非环形的队列（数组来实现）