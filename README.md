Golang 宝典
===========

`golang`，简称 [go](https://zh.wikipedia.org/wiki/Go)，是 google 2009 年推出的强类型编译型系统语言，具有极强的语言级高并发能力。
由于拥有极强的高并发能力，go 语言尤其适合于做服务端开发。go 语言设计简单，非常好学，它具备垃圾回收功能，这和 c 有天壤之别。
go 在语法上和 c 相似，但仅仅只是相似，事实上差异还是蛮大的，比如说

* 语句结尾不用分号
* if,for 等语句不用圆括号
* 使用 var 申明变量或干脆使用 `:=` 自动推导类型
* 类型放在变量后面
* 函数使用 func 申明
* ...

具体的语法将在后续章节中逐步介绍，这里不啰嗦了
先来看一看 go 语言版的 `hello world`

```go
// runnable
package main

import "fmt"

func main() {
	fmt.Println("hello, world")
}
// output:
// hello, world
```

* [第一章：准备工作](./ch01/ch01.md)
	* [安装 go](./ch01/ch01-01-install-go.md)
	* [开发环境的搭建](./ch01/ch01-02-development-environment.md)
