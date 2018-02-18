# variable初始化

## 变量

### 变量声明
```
var variable variable-type
or
var (
  variable variable-type
  variable variable-type
)
```

### 变量初始化

```
// 正确的使用方式1
var variable variable-type = value

// 正确的使用方式2，编译器可以自动推导出v2的类型
var variable = value

// 正确的使用方式3，编译器可以自动推导出v3的类型
// :=用于明确表达同时进行变量声明和初始化的工作
variable := value
```
Go编译器可以从初始化表达式的右值推导出该变量应该声明为哪种类型，这让Go语言看起来有点像动态类型语言，尽管Go语言实际上是不折不扣的强类型语言（静态类型语言）。  

### 变量赋值

在Go语法中，变量初始化和变量赋值是两个不同的概念。

