# sync

> https://godoc.org/sync 

sync包提供了基础的同步原语，比如互斥锁。除了Once和WaitGroup这两个type，大多数type是供低级库例程使用的；高级同步最好通过channel和communication实现！  

在这个包中定义的包含了类型的值不应该被拷贝。  

## type Once

Once是一个只会执行一次的对象！可以用在只需要运行一次初始化函数的场景。

### func (*Once) Do
```
func (o *Once) Do(f func())
```
Do calls the function f if and only if Do is being called for the first time for this instance of Once.  
if once.Do(f) is called multiple times, only the first call will invoke f, even if f has a different value in each invocation.  

Once类型的实例只有在第一次调用Do函数时，会去调用f函数。  

### example

```
func test_once() {

	var once sync.Once

	onceBody := func() {
		fmt.Println("Only once")
	}

	for i := 0; i < 10; i++ {
		once.Do(onceBody)
	}

}
```

## type WaitGroup

A WaitGroup waits for a collection of goroutines to finish. The main goroutine calls Add to set the number of goroutines to wait for. 
Then each of the goroutines runs and calls Done when finished. At the same time, Wait can be used to block 
until all goroutines have finished.  

A WaitGroup must not be copied after first use.  

### function

```
func (wg *WaitGroup) Add(delta int)
func (wg *WaitGroup) Done()
func (wg *WaitGroup) Wait()
```

### example

This example fetches several URLs concurrently, using a WaitGroup to block until all the fetches are complete.  

```
var wg sync.WaitGroup
var urls = []string{
    "http://www.golang.org/",
    "http://www.google.com/",
    "http://www.somestupidname.com/",
}
for _, url := range urls {
    // Increment the WaitGroup counter.
    wg.Add(1)
    // Launch a goroutine to fetch the URL.
    go func(url string) {
        // Decrement the counter when the goroutine completes.
        defer wg.Done()
        // Fetch the URL.
        http.Get(url)
    }(url)
}
// Wait for all HTTP fetches to complete.
wg.Wait()
```



