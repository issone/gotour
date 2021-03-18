package main

import (
	"fmt"
	"go_t/ch20/util"
)

func main() {
	fmt.Println("先导入fmt包，才能使用")
	util.Test1()
}

// 一个包中可以有多个 init 函数，但是它们的执行顺序并不确定，所以如果你定义了多个 init 函数的话，要确保它们是相互独立的，一定不要有顺序上的依赖。
// 一般用来做初始化操作，比如数据库连接和一些数据的检查，确保我们可以正确地使用这个包。
func init() { // init 函数没有返回值，也没有参数，它先于 main 函数执行
	fmt.Println("init in main.go ")
}

//  go mod init  xx   创建模块

// 在使用第三方模块之前，需要先设置下 Go 代理，也就是 GOPROXY，这样我们就可以获取到第三方模块了
// 推荐 goproxy.io 这个代理，非常好用，速度也很快。要使用这个代理，需要进行如下代码设置：
// go env -w GO111MODULE=on
// go env -w GOPROXY=https://goproxy.io,direct

// 除了第三方模块外，还有我们自己开发的模块，放在了公司的 GitLab上，这时候就要把公司 Git 代码库的域名排除在 Go PROXY 之外，为此 Go 语言提供了GOPRIVATE 这个环境变量帮助我们达到目的。通过如下命令即可设置 GOPRIVATE：
// # 设置不走 proxy 的私有仓库，多个用逗号相隔（可选）
// go env -w GOPRIVATE=*.corp.example.com

// go get -u github.com/gin-gonic/gin  安装第三方模块， 安装成功后，通过 import 命令导入使用

// 导入后，无法编译通过，因为还没有同步 Gin 这个模块的依赖，也就是没有把它添加到go.mod 文件中。通过如下命令可以添加缺失的模块：
// go mod tidy
// 运行这一命令，就可以把缺失的模块添加进来，同时它也可以移除不再需要的模块， 并自动同步go.mod文件
// 所以我们不用手动去修改 go.mod 文件，通过 Go 语言的工具链比如 go mod tidy 命令，就可以帮助我们自动地维护、自动地添加或者修改 go.mod 的内容。

// 在 Go 语言中，包是同一目录中，编译在一起的源文件的集合。包里面含有函数、类型、变量和常量，不同包之间的调用，必须要首字母大写才可以。
// 而模块又是相关的包的集合，它里面包含了很多为了实现该模块的包，并且还可以通过模块的方式，把已经完成的模块提供给其他项目（模块）使用，达到了代码复用、研发效率提高的目的。
