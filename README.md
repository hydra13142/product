product
=======

four iterators of golang

这个库，包含四个类型：Isolate、Ordered、Arrange、Combine

它们都内部维护了一个[]int切片，并且都实现了接口Iterator：
    
    Pick(int)int 来获取[]int切片成员的值
    Next() 来生成下一个符合要求的切片
    Over() 告知使用者，是否已枚举结束
    
可以通过如下函数获取对应类型的对象指针：

    func NewIsolate(l,t int) *Isolate
    func NewOrdered(l,t int) *Ordered
    func NewArrange(l,t int) *Arrange
    func NewCombine(l,t int) *Combine
    
    其中，l表示内部[]int切片的长度，而t表示取值范围为[0,t)

此外，它们还都实现了如下两个函数:
    
    Init(l,t int) 对象进行重新初始化
    Length() int 返回内部[]int切片的长度
    
四种类型，对内部切片的要求是不同的：
    
    Isolate 切片的每一个成员都是独立的，从[0,t)选取，如1312
    Ordered 切片的每一个成员都是独立从[0,t)选取的，但保证每个成员都不小于它前面的成员，如0112
    Arrange 切片的所有成员都是从[0,t)选取的，且每个成员都是独一无二的，如3214（亦即：排列）
    Combine 切片的所有成员都是从[0,t)选取的，各成员独有，且为顺序递增的，如1234（亦即：组合）

New...函数或者Init后，内部序列必然是满足要求的最小序列：

    Isolate、Ordered 为000....
    Arrange、Combine 为123....
    每次调用Next()函数都向前查找返回第一个符合要求的序列
    
本包主要用于暴力枚举，多次重复选择检验等，没有随机元素；属于无脑强上流。
