# channel

一个channel只能传递一种类型的值，这个类型需要在声明channel时指定。对channel的操作是阻塞的！即如果读取一个channel，它里面没有值，则读取的时候就会阻塞；
对于写也一样，往channel写时不会立即返回，直到被读取时，才会返回！  

对channel的熟练使用，才能真正理解和掌握Go语言并发编程！  

## channel定义
### 双向channel
```
var chanName chan ElementType
```
### 只读channel
```
var chanName <-chan ElementType
```
### 只写channel
```
var chanName chan<- ElementType
```
### 初始化-make
不带缓存，第一次读写时都会阻塞:  
```
ch := make(chan int)
```
带缓存，连续第channel_num次读写时都会阻塞:  
```
ch := make(chan int, channel_num)
```
### 关闭channel
```
close(chanName)
```

## 读写操作
### 读
```
<-ch
```
### 写
```
ch <- value
```

## 超时机制-select

Go语言没有提供直接的超时处理机制，但我们可以利用select机制。  

```
    // 首先，我们实现并执行一个匿名的超时等待函数
    timeout := make(chan bool, 1)
    go func() {
        //time.Sleep(1e9) // 等待1秒钟
        time.Sleep(3 * time.Second)
        timeout <- true
    }()

    ct := make(chan int)
    go func() {
        //time.Sleep(1e9) // 等待1秒钟
        time.Sleep(1 * time.Second)
        ct <- 666
    }()
    // 然后我们把timeout这个channel利用起来
    select {
        case vv := <-ct:
            // 从ct中读取到数据
            fmt.Println("Get ct channel value, ", vv)
        case <-timeout:
            // 一直没有从ch中读取到数据，但从timeout中读取到了数据
            fmt.Println("Timeout")
    }
```
