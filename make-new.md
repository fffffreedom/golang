# make和new的区别

内建函数make用来为slice，map或chan类型分配内存和初始化一个对象(注意：只能用在这三种类型上)，跟new类似，第一个参数也是一个类型而不是一个值，
跟new不同的是，make返回类型的引用而不是指针，而返回值也依赖于具体传入的类型。  

务必记得make仅适用于map，slice 和 channel，并且返回的不是指针。应当用new获得特定的指针。  

```
package main

import (
	"fmt"
	"os"
)

// The make built-in function allocates and initializes an object of type
// slice, map, or chan (only). Like new, the first argument is a type, not a
// value. Unlike new, make's return type is the same as the type of its
// argument, not a pointer to it. The specification of the result depends on
// the type:
//	Slice: The size specifies the length. The capacity of the slice is
//	equal to its length. A second integer argument may be provided to
//	specify a different capacity; it must be no smaller than the
//	length, so make([]int, 0, 10) allocates a slice of length 0 and
//	capacity 10.
//	Map: An empty map is allocated with enough space to hold the
//	specified number of elements. The size may be omitted, in which case
//	a small starting size is allocated.
//	Channel: The channel's buffer is initialized with the specified
//	buffer capacity. If zero, or the size is omitted, the channel is
//	unbuffered.

// The new built-in function allocates memory. The first argument is a type,
// not a value, and the value returned is a pointer to a newly
// allocated zero value of that type.

func main() {
	a0 := make([]int, 0)
	fmt.Println(len(a0))

	var a1 []int = make([]int, 10)
	fmt.Println(a1)

	var a2 *[]int = new([]int)

	//前面只分配了指针，并没有分配空间，这里的访问，是会出错的
	//panic: runtime error: index out of range
	//所以后面还是要用make来分配空间
	//(*a2)[0] = 1

	*a2 = make([]int, 10)
	fmt.Println(a2)
	fmt.Println(*a2)
	/* 输出&[]，a2本身是一个地址 */

	//现在可以访问了
	(*a2)[0] = 1
	fmt.Println("(*a2)[0] = ", (*a2)[0])

	/* 创建一个初始元素个数为5的数组切片，元素初始值为0，并预留10个元素的存储空间 */
	ar := make([]int, 5, 10)
	fmt.Println(ar)

	s0 := make(chan os.Signal)
	fmt.Println("len of chan: ", len(s0))

	s1 := make(chan os.Signal, 5)
	fmt.Println("len of chan: ", len(s1))
}
```

