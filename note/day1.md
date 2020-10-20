## 解释型语言
- php python 等需要解释器将代码转换层二进制，处理器才能运行处理，java 虚拟机同样承担了解释器的作用。

## 编译型语言
- 一次编译多次运行，省略中间解释环节，运行速度更快。
> - go 语言编译运行时默认跑满整个多核 cpu，这也是 go 语言天生支持并发的原因，性能较其他语言高，而不像 redis 单核历程，想跑满多核需要部署多个 redis 服务。
> - go 语言自带垃圾回收，开发不需要考虑垃圾回收，开发效率也比较高。

## 环境安装

### 下载go

- 安装地址：[https://golang.google.cn/dl/](https://golang.google.cn/dl/)，安装目录选择不复杂的，好记的。

- 安装完成后，输入`go version`查看 go 版本号。

### 配置GOPATH路径

1. 在自己的电脑上新建一个目录`D:\go`（存放我编写的Go语言代码）
2. 在环境变量里，新建一项：`GOPATH:D:\go`
3. 在`D:\go`下新建三个文件夹，分别是：`bin`、`src`、`pkg`
4. 把`D:\go\bin`这个目录添加到`PATH`这个环境变量的后面
   1. Win7是英文的`;`分隔
   2. Win10是单独一行
5. 你电脑上`GOPATH`应该是有默认值的，通常是`%USERPROFILE%/go`， 你把这一项删掉，自己按照上面的步骤新建一个就可以了。 

### 下载并安装vscode

- 下载地址：[https://code.visualstudio.com/Download](https://code.visualstudio.com/Download)

- 安装 Go 扩展，用于支持go语言开发。

- 安装 Go 扩展包：

  1. 打开终端执行：go env -w GOPROXY=https://goproxy.cn,direct
  2. 配置 vscode：`"go.gopath": "C:\\gocode"`

  2. 按下`Ctrl+Shift+P`，在输入框中输入`>go:install`，下面会自动搜索相关命令，选择`Go:Install/Update Tools`这个命令，选中安装扩展点击确定，就会进行安装。

## 编译程序

### go build

1. 在项目目录下执行`go build`
2. 在其他路径下执行`go build`， 需要在后面加上项目的路径（项目路径从GOPATH/src后开始写起，编译之后的可执行文件就保存在当前目录下）
3. `go build -o hello.exe`

###  go run

- 像执行脚本文件一样执行Go代码

### go install

- `go install`分为两步：
	1. 先编译得到一个可执行文件
	2. 将可执行文件拷贝到`GOPATH/bin`

### 交叉编译

- Go支持跨平台编译

例如：在windows平台编译一个能在linux平台执行的可执行文件

```bash
SET CGO_ENABLED=0  // 禁用CGO
SET GOOS=linux  // 目标平台是linux
SET GOARCH=amd64  // 目标处理器架构是amd64
```

执行`go build`

Mac平台交叉编译：

```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build
```