# srun4-config
读取srun配置

Config 包

Config 包提供了从 srun.conf 和 system.conf 文件中读取和管理配置的功能。

安装

使用以下命令安装包：
```shell
go get -u github.com/Eros-Vertigo/srun4-config
```
使用方法

在你的 Go 代码中导入该包：
```go
import (
"fmt"

    "github.com/Eros-Vertigo/srun4-config"
)
```
初始化

在使用该包之前，你需要设置配置文件目录路径。你可以通过调用 config.SetConfigPath() 函数来实现：

```config.SetConfigPath("/path/to/config/dir")```

默认情况下，该包假设为生产环境，并从预定义的路径（/srun3/etc/srun.conf 和 /srun3/etc/system.conf）读取配置文件。但是，如果你使用 SetConfigPath() 设置了不同的配置文件路径，该包将使用指定的路径。

读取配置

要获取配置，使用 config.GetConfig() 函数：
```go
config, err := config.GetConfig()
if err != nil {
fmt.Printf("获取配置失败：%v\n", err)
return
}
```
GetConfig() 函数将返回一个 Config 结构体，其中包含加载的配置。

访问配置

一旦获得 Config 结构体，可以使用点符号访问各个配置值。例如：
```go
fmt.Println("Srun 用户名：", config.SrunConfig.Username)
fmt.Println("System 在线服务器：", config.SystemConfig.OnlineServer)
```

示例

以下是使用 Config 包的示例代码：
```go
package main

import (
"fmt"

    "github.com/Eros-Vertigo/srun4-config"
)

func main() {
// 设置配置文件目录路径
config.SetConfigPath("/path/to/config/dir")

    // 获取配置
    config, err := config.GetConfig()
    if err != nil {
        fmt.Printf("获取配置失败：%v\n", err)
        return
    }

    // 访问各个配置
    fmt.Println("Srun 用户名：", config.SrunConfig.Username)
    fmt.Println("System 在线服务器：", config.SystemConfig.OnlineServer)
}
```