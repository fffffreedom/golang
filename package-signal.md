# signal

## basic

在实际项目中我们可能有下面的需求：  
- 修改了配置文件后，希望在不重启进程的情况下重新加载配置文件；  
- 当用 Ctrl + C 强制关闭应用后，做一些必要的处理；  

这时候就需要通过信号传递来进行处理了。golang中对信号的处理主要使用os/signal包中的两个方法：  
- notify方法用来监听收到的信号；  
- stop方法用来取消监听。  

## 监听信号

notify方法原型
```
func Notify(c chan<- os.Signal, sig ...os.Signal) 
```
第一个参数表示接收信号的管道；  
第二个及后面的参数表示设置要监听的信号，如果不设置表示监听所有的信号。  

## 例子
### 简单例子

```
package main

import ( 
    "fmt" 
    "os" 
    "os/signal" 
)

func main() { 
    c := make(chan os.Signal) 
    signal.Notify(c)
    //signal.Notify(c, syscall.SIGHUP, syscall.SIGUSR2)  //监听指定信号

    s := <-c //阻塞直至有信号传入 
    fmt.Println("get signal:", s) 
}
```
### 较完整例子
```
package main

import "fmt"
import "time"
import "os"
import "os/signal"
import "syscall"

type signalHandler func(s os.Signal, arg interface{})

//定义一个map，来保存信号及对应的处理函数
type signalSet struct {
    m map[os.Signal]signalHandler
}

//初始化signalSet，并关联map
func signalSetNew()(*signalSet){
    ss := new(signalSet)
    ss.m = make(map[os.Signal]signalHandler)
    return ss
}

//注册signal，把未注册的signal作为key保存到map中，并指定相应的handle
func (set *signalSet) register(s os.Signal, handler signalHandler) {
    if _, found := set.m[s]; !found {
        set.m[s] =  handler
    }
}

//信号处理函数，set.m[sig] 就是main中的handle函数
func (set *signalSet) handle(sig os.Signal, arg interface{})(err error) {
    if _, found := set.m[sig]; found {
        set.m[sig](sig, arg)
        return nil
    } else {
        return fmt.Errorf("No handler available for signal %v", sig)
    }

    panic("won't reach here")
}

func main() {
    go sysSignalHandleDemo()
    time.Sleep(time.Hour) // make the main goroutine wait!
}

func sysSignalHandleDemo() {
    //新建一个signalSet
    ss := signalSetNew()
    
    //信号处理函数
    handler := func(s os.Signal, arg interface{}) {
        fmt.Printf("handle signal: %v\n", s)
    }

    //信号注册
    ss.register(syscall.SIGINT, handler)
    ss.register(syscall.SIGUSR1, handler)
    ss.register(syscall.SIGUSR2, handler)

    for {
        c := make(chan os.Signal)
        var sigs []os.Signal
        
        //注册了的信号放在slice中
        for sig := range ss.m {
            sigs = append(sigs, sig)
        }
        signal.Notify(c)
        //signal.Notify(c, sigs...)
        sig := <-c

        err := ss.handle(sig, nil)
        if (err != nil) {
            fmt.Printf("unknown signal received: %v\n", sig)
            os.Exit(1)
        }
    }
}
```

## Reference

https://godoc.org/os/signal  
Golang的signal  
https://studygolang.com/articles/2333  
Go中的系统Signal处理  
http://tonybai.com/2012/09/21/signal-handling-in-go/  
