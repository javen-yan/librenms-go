# LibreNMS Go SDK

[![Go Version](https://img.shields.io/badge/Go-1.21+-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

LibreNMS Go SDK 是一个用于与 LibreNMS API 进行交互的 Go 语言客户端库。该库提供了完整的 CRUD 操作支持，适用于构建 LibreNMS 管理工具、自动化脚本或 Terraform Provider。

## 🚀 功能特性

- **完整的 API 支持**: 支持 LibreNMS API v0 版本的所有主要功能
- **资源管理**: 提供以下资源的完整 CRUD 操作：
  - 🚨 告警规则 (Alert Rules)
  - 🖥️ 设备管理 (Devices)
  - 👥 设备组 (Device Groups)
  - 📍 位置管理 (Locations)
  - 🔧 服务管理 (Services)
  - 🔌 端口管理 (Ports)
  - 🗂️ 库存管理 (Inventory)
  - 🛣️ 路由管理 (Routing)
  - 🔀 交换管理 (Switching)
- **类型安全**: 使用 Go 强类型系统，提供类型安全的 API 调用
- **错误处理**: 完善的错误处理和响应检查
- **日志支持**: 内置结构化日志记录
- **HTTP 客户端**: 基于标准库的 HTTP 客户端，支持自定义配置

## 📦 安装

```bash
go get github.com/javen-yan/librenms-go
```

## 🔧 快速开始

### 创建客户端

```go
package main

import (
    "log"
    "log/slog"
    
    "github.com/javen-yan/librenms-go"
)

func main() {
    // 创建新的 LibreNMS 客户端
    client, err := librenms.NewClient(
        "http://your-librenms-server:8000",  // LibreNMS 服务器地址
        "your-api-token",                    // API 令牌
        librenms.WithLogLevel(slog.LevelDebug), // 可选：设置日志级别
    )
    if err != nil {
        log.Fatalf("创建 LibreNMS 客户端失败: %v", err)
    }
    
    // 使用客户端...
}
```

### 基本使用示例

#### 获取系统信息

```go
// 获取系统信息
systemInfo, err := client.System.Get()
if err != nil {
    log.Printf("获取系统信息失败: %v", err)
} else {
    if len(systemInfo.System) > 0 {
        sys := systemInfo.System[0]
        fmt.Printf("LibreNMS 版本: %s\n", sys.LocalVer)
        fmt.Printf("数据库版本: %s\n", sys.DatabaseVer)
        fmt.Printf("PHP 版本: %s\n", sys.PHPVer)
    }
}
```

#### 设备管理

```go
// 列出所有设备
devices, err := client.Device.List(nil)
if err != nil {
    log.Printf("获取设备列表失败: %v", err)
} else {
    fmt.Printf("找到 %d 个设备\n", devices.Count)
    for _, device := range devices.Devices {
        fmt.Printf("- %s (%s) - %s\n", device.Hostname, device.Display, device.OS)
    }
}

// 获取特定设备
device, err := client.Device.Get("123")
if err != nil {
    log.Printf("获取设备详情失败: %v", err)
} else {
    fmt.Printf("设备 ID: %d\n", device.Devices[0].DeviceID)
    fmt.Printf("主机名: %s\n", device.Devices[0].Hostname)
}
```

#### 告警管理

```go
// 列出所有告警
alerts, err := client.Alert.List(nil)
if err != nil {
    log.Printf("获取告警列表失败: %v", err)
} else {
    fmt.Printf("找到 %d 个告警\n", alerts.Count)
    for _, alert := range alerts.Alerts {
        fmt.Printf("- %s (%s) - %s\n", alert.Name, alert.Hostname, alert.Severity)
    }
}
```

#### 位置管理

```go
// 列出所有位置
locations, err := client.Location.List()
if err != nil {
    log.Printf("获取位置列表失败: %v", err)
} else {
    fmt.Printf("找到 %d 个位置\n", locations.Count)
    for _, location := range locations.Locations {
        fmt.Printf("- %s\n", location.Name)
    }
}
```

## 📚 API 参考

### 客户端配置选项

```go
client, err := librenms.NewClient(
    "http://server:8000",
    "token",
    librenms.WithLogLevel(slog.LevelDebug),    // 设置日志级别
    librenms.WithHTTPClient(customHTTPClient), // 自定义 HTTP 客户端
    librenms.WithLogger(customLogger),         // 自定义日志记录器
)
```

### 支持的资源类型

| 资源 | 包名 | 主要方法 |
|------|------|----------|
| 设备 | `client.Device` | `List()`, `Get()`, `Add()`, `Update()`, `Delete()` |
| 设备组 | `client.DeviceGroup` | `List()`, `Get()`, `Add()`, `Update()`, `Delete()` |
| 位置 | `client.Location` | `List()`, `Get()`, `Add()`, `Update()`, `Delete()` |
| 服务 | `client.Service` | `List()`, `Get()`, `Add()`, `Update()`, `Delete()` |
| 告警 | `client.Alert` | `List()`, `Get()`, `Add()`, `Update()`, `Delete()` |
| 告警规则 | `client.AlertRule` | `List()`, `Get()`, `Add()`, `Update()`, `Delete()` |
| 端口 | `client.Port` | `List()`, `Get()`, `Add()`, `Update()`, `Delete()` |
| 系统 | `client.System` | `Get()` |

## 🧪 测试

运行测试套件：

```bash
go test ./...
```

运行特定包的测试：

```bash
go test ./device
go test ./alert
```

## 📖 完整示例

查看 `examples/` 目录中的完整示例代码，了解如何使用 SDK 的各种功能。

## 🔗 相关链接

- [LibreNMS 官方文档](https://docs.librenms.org/)
- [LibreNMS API 文档](https://docs.librenms.org/API/)

## 🤝 贡献

欢迎提交 Issue 和 Pull Request！

## 📄 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。

## ⚠️ 注意事项

- 确保您的 LibreNMS 服务器已启用 API 访问
- API 令牌需要适当的权限才能执行相应操作
- 建议在生产环境中使用 HTTPS 连接
- 请遵循 LibreNMS 的 API 使用限制和最佳实践

## 🆘 支持

如果您遇到问题或有疑问，请：

1. 查看 [LibreNMS 官方文档](https://docs.librenms.org/)
2. 检查 [Issues](https://github.com/javen-yan/librenms-go/issues) 页面
3. 创建新的 Issue 描述您的问题
