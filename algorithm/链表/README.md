# 链表
    链表由一个个数据节点组成，它是一个递归结构，要么它是空的，要么它存在一个指向另外一个数据节点的引用
    链表可以通过数组来实现，但是数组是连续的，增加和删除节点时容易造成冗余，效果不佳，因此golang的链表使用结构体来实现
    

container/list 实现了一个双向链表
container/ring 实现了环形链表的操作