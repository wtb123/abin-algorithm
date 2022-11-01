## 一、Redis 基本数据类型
String(字符串)、Hash(哈希)、List(链表)、Set(无序集合)、ZSet(有序集合)

Redis 中的每一个对象都由一个 redisObject 结构表示，该结构中和保存数据有关的三个属性分别是:
type 属性、encoding 属性、ptr属性

```c
typedef struct redisObject {
    // 类型
    unsigned type:4;
    // 编码
    unsigned encoding:4;
    // 指向底层实现数据结构的指针
    void *ptr;
    // ...
} robj;
```

