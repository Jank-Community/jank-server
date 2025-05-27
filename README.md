# Jank

<p style="text-align: center;">
  <a><img src="https://s2.loli.net/2025/03/14/BnchjpPLeIaoO75.png" alt="Jank"></a>
</p>

<p style="text-align: center;">
  <em>Jank，一个轻量级的博客系统，基于 Go 语言和 Echo 框架开发，强调极简、低耦合和高扩展</em>
</p>

<p style="text-align: center;">
  <a href="https://img.shields.io/github/stars/Done-0/Jank?style=social" target="_blank">
    <img src="https://img.shields.io/github/stars/Done-0/Jank?style=social" alt="Stars">
  </a> &nbsp;
  <a href="https://img.shields.io/github/forks/Done-0/Jank?style=social" target="_blank">
    <img src="https://img.shields.io/github/forks/Done-0/Jank?style=social" alt="Forks">
  </a> &nbsp;
  <a href="https://img.shields.io/github/contributors/Done-0/Jank" target="_blank">
    <img src="https://img.shields.io/github/contributors/Done-0/Jank" alt="Contributors">
  </a> &nbsp;
  <a href="https://img.shields.io/github/issues/Done-0/Jank" target="_blank">
    <img src="https://img.shields.io/github/issues/Done-0/Jank" alt="Issues">
  </a> &nbsp;
  <a href="https://img.shields.io/github/issues-pr/Done-0/Jank" target="_blank">
    <img src="https://img.shields.io/github/issues-pr/Done-0/Jank" alt="Pull Requests">
  </a> &nbsp;
  <a href="https://img.shields.io/github/license/Done-0/Jank" target="_blank">
    <img src="https://img.shields.io/github/license/Done-0/Jank" alt="License">
  </a>
</p>

---

Jank 是一个轻量级的博客系统，基于 Go 语言和 Echo 框架开发，设计理念强调简约、低耦合和高扩展，旨在为用户提供优雅的博客体验。

## 快速链接

👉 [演示站](https://www.jank.org.cn) | [开发社区](https://github.com/Jank-Community)

## 界面预览

![首页](https://s2.loli.net/2025/04/07/l1tGYV4WkmoiIHv.png)
![文章列表](https://s2.loli.net/2025/04/07/xR62vhWKsmgw3Ht.png)
![文章详情1](https://s2.loli.net/2025/04/07/DbcJzryKmBNR7vQ.png)
![文章详情2](https://s2.loli.net/2025/04/07/iNpXyMdkjaDbn92.png)

## 技术栈

- **Go 语言**：热门后端开发语言，适合构建高并发应用。
- **Echo 框架**：高性能的 Web 框架，支持快速开发和灵活的路由管理。
- **数据库**：开源的关系型数据库，支持 Postgres、MySQL 和 SQLite。
- **Redis**：热门缓存解决方案，提供快速数据存取和持久化选项。
- **JWT**：安全的用户身份验证机制，确保数据传输的完整性和安全性。
- **Docker**：容器化部署工具，简化应用的打包和分发流程。
- **前端**：react + ts + vite + shadcn/ui + tailwindcss。

## 功能模块

- **账户模块**：实现 JWT 身份验证，支持用户登录、注册、注销、密码修改和个人信息更新。
- **权限模块**：实现 RBAC（Role-Based Access Control）角色权限管理，支持用户-角色-权限的增删改查。
  - 基本功能已实现，考虑到用户使用的不友好性和复杂性，因此暂不推出此功能。
- **文章模块**：提供文章的创建、查看、更新和删除功能。
- **分类模块**：支持类目树及子类目树递归查询，单一类目查询，以及类目的创建、更新和删除。
- **评论模块**：提供评论的创建、查看、删除和回复功能，支持评论树结构的展示。
- **插件系统**：正在火热开发中，即将推出...
- **其他功能**：
  - 提供 OpenAPI 接口文档
  - 集成 Air 实现热重载
  - 提供 Logrus 实现日志记录
  - 支持 CORS 跨域请求
  - 提供 CSRF 和 XSS 防护
  - 支持 Markdown 的服务端渲染
  - 集成图形验证码功能
  - 支持 QQ/Gmail/Outlook 等主流邮箱服务端发送能力
  - 支持 oss 对象存储（MinIO）
  - **其他模块正在开发中**，欢迎提供宝贵意见和建议！

## 开发指南

### 本地开发

1. **安装依赖**

```bash
# 安装 swagger 工具
go install github.com/swaggo/swag/cmd/swag@latest
# 安装 air，需要 go 1.22 或更高版本
go install github.com/air-verse/air@latest
# 安装依赖包
go mod tidy
```

2. **配置数据库和邮箱**

```yaml
APP:
  APP_NAME: "JANK_BLOG"
  APP_HOST: "127.0.0.1" # 如果使用 docker，则改为"0.0.0.0"
  APP_PORT: "9010"
  EMAIL:
    EMAIL_TYPE: "qq" # 支持的邮箱类型: qq, gmail, outlook
    FROM_EMAIL: "<FROM_EMAIL>" # 发件人邮箱
    EMAIL_SMTP: "<EMAIL_SMTP>" # SMTP 授权码
  SWAGGER:
    SWAGGER_HOST: "127.0.0.1:9010"
    SWAGGER_ENABLED: true

DATABASE:
  DB_DIALECT: "postgres" # 数据库类型: postgres, mysql, sqlite
  DB_NAME: "jank_db"
  DB_HOST: "127.0.0.1" # 如果使用 docker，则改为"postgres_db"
  DB_PORT: "5432"
  DB_USER: "fender"
  DB_PSW: "Lh20230623"
  DB_PATH: "./database" # SQLite 数据库文件路径
```

3. **启动服务**

```bash
# 方式一：直接运行
go run main.go

# 方式二：使用 Air 热重载（推荐）
air -c ./configs/.air.toml
```

### Docker 部署

1. **修改配置**
   修改 `configs/config.yaml` 和 `docker-compose.yaml` 中的相关配置

2. **启动容器**

```bash
docker-compose up -d
```

## 开发路线图

![开发路线图](https://s2.loli.net/2025/03/09/qJrtOeFvD95PV4Y.png)

## 社区支持

### 官方社区

<img src="https://s2.loli.net/2025/05/27/bWwu4sN5mHfTSPa.jpg" alt="官方社区" width="300" />

> 注：因社群成员较多，请自觉遵守规范。严禁讨论涉黄、赌、毒及政治敏感内容，禁止发布任何形式的不良广告。

### 贡献者

<a href="https://github.com/Done-0/Jank/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=Done-0/Jank" alt="贡献者名单" />
</a>

### 特别鸣谢

<p>
  <a href="https://github.com/vxincode">
    <img src="https://github.com/vxincode.png" width="70" height="70" style="border-radius: 50%;" />
  </a>
  <a href="https://github.com/WowDoers">
    <img src="https://github.com/WowDoers.png" width="70" height="70" style="border-radius: 50%;" />
  </a>
</p>

## 联系方式

- **QQ**: 927171598
- **微信**: l927171598
- **邮箱**: fenderisfine@outlook.com

> 合作、推广和赞助可联系作者

## 许可证

本项目遵循 [MIT 协议](https://opensource.org/licenses/MIT)

## 项目趋势

<img src="https://api.star-history.com/svg?repos=Done-0/Jank&type=timeline" width="100%" height="65%" alt="GitHub Stats">
