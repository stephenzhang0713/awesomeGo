## 什么是Map
### key，value存储
最通俗的话说Map是一种通过key来获取value的一个数据结构，其底层存储方式为数组，在存储时key不能重复，当key重复时，value进行覆盖，我们通过key进行hash运算（可以简单理解为把key转化为一个整形数字）然后对数组的长度取余，得到key存储在数组的哪个下标位置，最后将key和value组装为一个结构体，放入数组下标处，看下图：
```
length = len(array) = 4
hashkey1 = hash(xiaoming) = 4
index1  = hashkey1% length= 0
hashkey2 = hash(xiaoli) = 6
index2  = hashkey2% length= 2
```
