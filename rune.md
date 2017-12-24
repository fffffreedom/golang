# Unicode
## go语言编程
在Go语言中支持两个字符类型，一个是byte（实际上是uint8的别名），代表UTF-8字符串的单个字节的值；另一个是rune，代表单个Unicode字符。
关于rune相关的操作，可查阅Go标准库的unicode包。另外unicode/utf8包也提供了UTF8和Unicode之间的转换。
出于简化语言的考虑， Go语言的多数API都假设字符串为UTF-8编码。尽管Unicode字符在标准库中有支持，但实际上较少使用。
## The-Golang-Standard-Library-by-Example
http://books.studygolang.com/The-Golang-Standard-Library-by-Example/chapter02/02.5.html
## uft8库解读
https://www.cnblogs.com/golove/p/5889790.html
