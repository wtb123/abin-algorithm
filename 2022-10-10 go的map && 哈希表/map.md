### 一、go Map 的基本使用
1. 初始化 map 时推荐使用内置函数 make() 并指定预估的容量。
2. 修改键值对时，需要先查询指定的键是否存在，否则 map 将创建新的键值对。
3. 查询键值对时，最好检查键是否存在，避免操作零值。
4. 避免并发读写 map，如果需要并发读写，则可以使用额外的锁（互斥锁、读写锁），也可以考虑使用标准库 sync 包中的 sync.Map

### 二、go Map 的数据结构
1. go 中 map 的底层实现是一个散列表，go 使用 runtime.hmap 这个 struct 来表示 map
```go
// 
type hmap struct {

    count     int // 当前保存的元素个数
    flags     uint8  // 用于标记当前 map 的状态
    B         uint8  // bucket 数组大小 2^B
    noverflow uint16 // overflow 桶的大致数量
    hash0     uint32 // hash seed

    buckets    unsafe.Pointer // bucket 数组，数组长度为 2^B
    oldbuckets unsafe.Pointer // 老旧 bucket，用于扩容
    nevacuate  uintptr        // map 增长的时候使用
    extra *mapextra // optional fields
}

// Bucket：bmap 实际上是一个可以存放8个元素的unit8数组，存储的是高8位的hash值
type bmap struct {
	tophash [bucketCnt]unit8 // 存储高哈希值的8位
}

// bucket 的数据结构中的 data 和 overflow 并没有显式地在结构体中声明
type bmap struct {
	tophash   [bucketCnt]unit8 // 存储高哈希值的8位
	keys      [8]keytype
	values    [8]valuetype
	pad       uintptr
	overflow  unitptr
}

type mapextra struct {
    overflow    *[]*bmap
    oldoverflow *[]*bmap
    // nextOverflow holds a pointer to a free overflow bucket.
    nextOverflow *bmap
}

```
<img src="https://img-blog.csdnimg.cn/14fdb4a8324a4958ae21b2f38f875a08.jpg?x-oss-process=image/watermark,type_d3F5LXplbmhlaQ,shadow_50,text_Q1NETiBA546L6Ie05YiX,size_20,color_FFFFFF,t_70,g_se,x_16">
### 三、go Map 的增删查改

无论是元素的添加还是查询操作，都需要先根据键的 Hash 值确定一个 bucket，并查询该 bucket 是否存在指定的键。
* 对于查询操作而言，查到指定的键后获取值并返回，否则返回类型的空值。
* 对于添加操作而言，查到指定的键意味着当前的添加操作实际上是更新操作，否则在 bucket 中查找一个空余位置插入。

1. 查找过程
查找过程简述如下：
* 根据 key 值计算 Hash 值；
* 取 Hash 值低位与 hmap.B 取模来确定 bucket 的位置；
* 取 Hash 值高位，在 tophash 数组中查询；
* 如果 tophash[i] 中存储的哈希值与当前 key 的哈希值相等，则后去 tophash[i] 的key 值进行比较；
* 当前 bucket 中没有找到，则依次从溢出的 bucket 中查找。
如果当前的 map 处于搬迁过程中，那么查找时优先从 oldbuckets 数组中查找，如果 oldbuckets 找到的 bucket 已经搬迁则到新的 buckets 数组中查找。

2. 添加过程
新元素添加过程描述如下：
* 根据 key 算出 Hash 值；
* 取 Hash 值低位与 hmap.B 取模来确定 bucket 的位置；
* 查找该 key 是否存在，如果存在则直接更新；
* 如果该 key 不存在，则从该 bucket 中寻找空余位置并插入。
如果当前的 map 处于搬迁过程中，那么新元素会直接添加到新的 buckets数组中，单查询过程仍从 oldbuckets 数组中开始。

### 四、扩展：哈希表常见冲突解决
go 的 map 是使用拉链发来解决哈希冲突的，常见的哈希冲突解决方法有：
1. 开放地址法
为冲突的地址 H(key)，按照某种规则产生另一个地址的方法
* 线性探测法
* 平方探测法
* 随机探测法

2. 拉链法
将所有的散列地址相同的记录都存储在一个单链表中---称为同义词子表，散列表存储所有同义词的头指针。

3. 公共溢出法
基本思想
散列表包含基本表和溢出表两个部分，将发生冲突的记录存储在溢出表中。
查找方法
通过 H(key) 函数计算散列地址，先与基本表中记录进行比较，若相等，则查找成功；否则，到溢出表中顺序查找。

### 五、参考资料
1. [Go Map 底层原理](https://blog.csdn.net/star_of_science/article/details/121802354)
2. [小白学 go 基础篇3 -- map哈希表](https://blog.csdn.net/hinsss/article/details/119981795)
3. [Go 实现哈希表](https://blog.csdn.net/chengqiuming/article/details/117424064)

