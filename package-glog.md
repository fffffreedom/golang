# glog

> https://www.godoc.org/github.com/golang/glog  

Package glog implements logging analogous to the Google-internal C++ INFO/ERROR/V setup.  

It provides functions Info, Warning, Error, Fatal, plus formatting variants such as Infof. 
It also provides V-style logging controlled by the -v and -vmodule=file=2 flags.  

Log output is buffered and written periodically using Flush. 
Programs should call Flush before exiting to guarantee all log output is written.  

By default, all log statements write to files in a temporary directory. 
This package provides several flags that modify this behavior. As a result, flag.Parse must be called before any logging is done.  

## 控制打印信息的命令行参数
```
-logtostderr=false
	Logs are written to standard error instead of to files.
-alsologtostderr=false
	Logs are written to standard error as well as to files.
-stderrthreshold=ERROR
	Log events at or above this severity are logged to standard
	error as well as to files.
-log_dir=""
	Log files will be written to this directory instead of the
	default temporary directory.

Other flags provide aids to debugging.

-log_backtrace_at=""
	When set to a file and line number holding a logging statement,
	such as
		-log_backtrace_at=gopherflakes.go:234
	a stack trace will be written to the Info log whenever execution
	hits that statement. (Unlike with -vmodule, the ".go" must be
	present.)
-v=0
	Enable V-leveled logging at the specified level.
-vmodule=""
	The syntax of the argument is a comma-separated list of pattern=N,
	where pattern is a literal file name (minus the ".go" suffix) or
	"glob" pattern and N is a V level. For instance,
		-vmodule=gopher*=3
	sets the V level to 3 in all Go files whose names begin "gopher".
```

## 例子
```
package main

import (
    "github.com/golang/glog"
    "flag"
)

func main() {
    //初始化命令行参数
    flag.Parse()

    //退出时调用，确保日志写入文件中
    defer glog.Flush()

    glog.Info("Info, glog")
    glog.Warning("Warning glog")
    glog.Error("Error glog")

    glog.Infof("Infof %d", 1)
    glog.Warningf("Warningf %d", 2)
    glog.Errorf("Errorf %d", 3)


    if glog.V(2) {
        glog.Info("glog.V()")
    }
    glog.V(2).Infoln("glog.V().Infoln")
}
```
在IntelliJ IDEA的Terminal中运行：  
```
go run glog.go -log_dir=./ -logtostderr=TrueOrFalse -v level
```
其中，glog.{Info|Warning|Error}会把信息打印到文件中，glog.V(N)指定了打印信息的等级，是否能打印出来，由命令行中的-v参数决定。  

`-v level`的level值如果小于N，则信息可以打印到终端或者文件中！大于等于N则不能被打印出来！  
`-log_dir`指定日志文件保存的路径。  
`-logtostderr`指定打印信息输出到标准错误输出还是文件。  

## 日志格式
```
Log line format: [IWEF]mmdd hh:mm:ss.uuuuuu threadid file:line] msg  
```
