# interface struct

本文主要研究一下interface和struct之间的关系：  
- struct和interface代码定义
```
type struct_name struct {
	......
}
```
- interface的定义  
**interface是一组方法的集合！**  
- struct的继承
```
type sa struct {
	......
}
type sb struct {
	sa
	......
}
```
- interface的赋值
  - 将对象实例赋值给interface
    对象实例（包括struct）实现了interface的**所有**方法，则可以将对象赋值给interface对象；  
  - 将一个interface赋值给另一个interface
    一个interface A实现了另一个interface B的**所有**方法，则可以将interface A赋值给interface B；  
- struct继承初始化
如下代码，sb继承了sa，在初始化时，可以按如下进行：  
```
	sbo := sb {
		sa: sa {
			name: "jonny",
			age: 32,
		},
		job: "IT",
	}
```

##

```
package main

import (
	"fmt"
)

type test interface {
	bar()
	foo()
}

type sa struct {
	name string
	age int
}

func (s sa) bar() {
	fmt.Println(s.name)
}

func (s sa) foo() {
	fmt.Println(s.age)
}

type sb struct {
	sa
	job string
}

// sb结构体并没有实现foo函数，即没有实现test接口，
// 但由于它继承的sa已经实现，所以sb结构体也是实现了test接口
// 也可以直接赋值给test接口的对象
func (s sb) jaz() {
	fmt.Println(s.job)
}

func (s sb) bar() {
	s.sa.name = "freedom"
	fmt.Println(s.name)
}

func main () {
	sao := sa {
		name: "jonny",
		age: 32,
	}

	//结构体实现了接口，可以把结构体赋值给接口
	var to test = sao
	to.bar()
	to.foo()

	sbo := sb {
		sa: sa {
			name: "jonny",
			age: 32,
		},
		job: "IT",
	}

	var tt test = sbo
	tt.bar()
	tt.foo()

	sbo.jaz()
}


```
