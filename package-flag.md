# flag

> https://www.godoc.org/flag

Package flag implements command-line flag parsing.  
Usage:  
  Define flags using flag.String(), Bool(), Int(), etc.  

## functions
### basic functions
```
func Int(name string, value int, usage string) *int
func IntVar(p *int, name string, value int, usage string)
func String(name string, value string, usage string) *string
func StringVar(p *string, name string, value string, usage string)
......
```
两种类型的函数，一种是返回指针，Var是直接传入参数；函数的参数说明：  
- 参数变量指针，获取参数传进来的值
- 参数名
- 参数默认值
- 参数说明

使用的方法如下：
```
var ip = flag.Int("flagname", 1234, "help message for flagname")
or

var flagvar int
func init() {
	flag.IntVar(&flagvar, "flagname", 1234, "help message for flagname")
}
```

在所有flag定义好后，就可以调用如下解析函数，该函数会把命令行参数的值，保存到定义的flags中：  
```
flag.Parse()
```

### special functions

## 命令行参数语法
```
-flag
-flag=x
-flag x  // non-boolean flags only
or
--flag
--flag=x
--flag x  // non-boolean flags only
```
一个或两个前置中划线`-`都是可以的，它们是等同的；最后一种形式不能用在bool类型参数，必须使用-flag=false的形式来关闭一个Bool参数。  

Flag parsing stops just before the first non-flag argument ("-" is a non-flag argument) or after the terminator "--".  
下面的命令只能设置name参数，后面的age被`-和--`中止了：  
```
go run flag.go -name test -- -age 88
go run flag.go -name test - -age 88
```

## Duration flags

Duration flags accept any input valid for time.ParseDuration.  

## flagset

The default set of command-line flags is controlled by top-level functions. The FlagSet type allows one to define independent sets of flags, such as to implement **subcommands** in a command-line interface. The methods of FlagSet are analogous to the top-level functions for the command-line flag set.  


