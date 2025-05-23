# 统一日志组件 (Unified Logging Component)

## 简介

统一日志组件提供了应用程序的日志记录功能，支持日志轮转、级别控制和格式化输出。基于 logrus 库实现，提供结构化 JSON 格式的日志输出，便于后续日志分析和处理。

## 核心功能

- **结构化日志**: 使用 JSON 格式输出日志，便于解析和分析
- **日志级别控制**: 支持多种日志级别 (Panic, Fatal, Error, Warn, Info, Debug, Trace)
- **日志轮转**: 按照时间自动切割日志文件，防止单个日志文件过大
- **文件保留策略**: 自动清理超过保留期限的历史日志文件
- **优雅降级**: 日志系统故障时会自动降级到标准输出

## 配置项

日志组件从应用配置中读取以下参数：

- **LogFilePath**: 日志文件存储路径
- **LogFileName**: 日志文件名称
- **LogLevel**: 日志记录级别
- **LogTimestampFmt**: 时间戳格式
- **LogMaxAge**: 日志文件最大保留时间（小时）
- **LogRotationTime**: 日志轮转时间间隔（小时）

## 文件权限说明

**0755**: `Unix/Linux` 系统中常用的文件权限表示法。使用八进制（octal）数字系统来表示文件或目录的权限。每个数字表示一组权限，分别对应用户、用户组和其他人

- 第一个数字（0）：表示文件类型。对于常规文件，通常为 0
- 第二个数字（7）：表示文件所有者（用户）的权限 (这里 7 表示文件所有者拥有读（4）、写（2）和执行（1）的权限，合计 4 + 2 + 1 = 7)
- 第三个数字（5）：表示与文件所有者同组的用户组的权限 (这里 5 表示用户组和其他用户拥有读（4）和执行（1）的权限，合计 4 + 1 = 5)
- 第四个数字（5）：表示其他用户的权限
- 因此 0755 表示：
  - 文件所有者可以读、写、执行。
  - 用户组成员可以读、执行。
  - 其他用户可以读、执行。

## 使用方式

通过全局变量 `global.SysLog` 在应用的任何位置使用日志功能：

```go
// 记录信息日志
global.SysLog.Info("应用启动成功")

// 记录带字段的错误日志
global.SysLog.WithFields(logrus.Fields{
    "user": "admin",
    "action": "login",
}).Error("登录失败: 密码错误")
```
