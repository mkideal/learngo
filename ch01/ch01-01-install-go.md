# 安装

## 下载

Go 是跨平台的语言，各个系统平台下的安装也非常简单。首先从Go官方下载安装包，不过墙内的朋友需要翻墙才能进入Go官方网站。但是也有别的镜像地址可用

* [Go 官方下载地址](https://golang.org/dl/)
* [Golang 中国镜像](http://www.golangtc.com/download)

截至 `2016-08-30 00:33`，Go的最新版本为1.7，一般下载列表里都有已编译好的二进制包，也有原代码包。

* Windows 下的可以下载 `*.msi` 安装包
* Mac 下的可以下载 `*.pkg` 安装包
* Linux 下的可以下载二进制的 `*.tar.gz` 包或源代码 `*.tar.gz` 包（二进制的包要比源代码的包要大很多）

## 安装下载包

安装过程很简单，Windows和Mac下基本用鼠标点击安装包，然后一直下一步或确认到最后就Ok了。

Linux下解压二进制`*.tar.gz`包，假设解压到 $HOME/go 目录，那么需要设置 `GOROOT` 和 `PATH` 环境变量

```sh
export GOROOT=$HOME/go
export PATH=$PATH:$GOROOT/bin
```

## 建立工作目录，设置环境变量

无论在哪个平台下安装，最后都需要建立一个写Go的目录，这个就是Golang工作目录。假设这个工作目录叫做 `workspace`，那么在`workspace`下需要建立一个名为 `src`的子目录，顾名思义 `src`就是代码目录。

再添加环境变量 `GOPATH`，其值为 workspace 的完整路径，假设就叫 `/path/to/workspace`

```sh
export GOPATH=/path/to/workspace
```

> Windows 的环境变量设置方式：添加一个用户自定义的环境变量，取名叫做 `GOPATH`, 值为 workspace 目录完整路径

## Hello,world

我们再在src目录下建一个项目目录，这里我们建一个 `helloworld` 目录，现在我们的目录结构如下:

	workspace
	└── src
		└── helloworld

在 helloworld 下我们建立一个名为 main.go 的文件，文件内容如下：

```go
// index:ch01-01-01
package main

import "fmt"

func main() {
	fmt.Print("hello,world")
}
// output:
// hello,world
```

> Go 语言的原文件使用 .go 作为后缀

然后在 helloworld 目录下执行（Windows下使用cmd命令行窗口，Mac或Linux下使用终端shell）

```sh
go run main.go
```

如果可以看到 `hello,world` 的输出，那么说明你已经安装成功了。

> 由于 Windows 下的命令行和 Unix 的终端有差异，为避免每次需要敲命令时都繁琐的特别区分，以后统一使用 `命令行` 一词。另外，Go 作为主要应用于服务器的语言，强烈建议大家在 Linux 或 Mac 下进行 Go 开发。目前不能摆脱 Windows 的可以考虑安装一个 Linux 发行版的虚拟机，在学习 Go 之余可以花点时间熟悉 Linux 系统，这里推荐使用 `Linux Mint` 这个发行版本

[上一级](./ch01.md)
[下一节](./ch01-02-development-environment.md)
